package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jindogroup/kibl-client-go/auth"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func HttpClient(log *logrus.Entry, conf *Conf) (client *httpAPI) {
	client = &httpAPI{
		log:                  log.WithField("api.client", "kibl.httpAPI"),
		client:               *http.DefaultClient,
		inboundHost:          conf.InboundHost,
		outboundHost:         conf.OutboundHost,
		retryTimes:           conf.RetryTimes,
		retryTimeout:         conf.RetryTimeout,
		requestTimeout:       conf.RequestTimeout,
		authenticateRequests: !conf.DisableAuth,
	}
	if !client.authenticateRequests {
		return
	}
	var err error
	client.session, err = auth.GetSession(client.log, conf.prefix)
	if err != nil {
		log.WithError(err).Fatal("failed to initialise kibl auth session")
	}
	return
}

type httpAPI struct {
	inboundHost          string
	outboundHost         string
	retryTimes           int
	retryTimeout         time.Duration
	requestTimeout       time.Duration
	authenticateRequests bool

	client  http.Client
	session auth.Session

	log *logrus.Entry
}

func (c *httpAPI) headers() (h Headers) {
	h = Headers{
		"Content-Type": "application/json",
	}
	if c.authenticateRequests {
		h["Authorization"] = "Bearer " + c.session.AccessToken()
	}
	return h
}

func (a httpAPI) inbound(path string) string {
	return a.inboundHost + "/sports/update/" + path
}

func (a httpAPI) outbound(path string) string {
	return a.outboundHost + "/sports/get/" + path
}

func (a httpAPI) archive(path string) string {
	return a.inboundHost + "/sports/archive/" + path
}

func (a httpAPI) getInfoUrl(path string) string {
	return a.outbound("info/" + path)
}

func (a httpAPI) getMappingUrl(path string) string {
	return a.outbound("mapping/" + path)
}

func (a httpAPI) getReferenceUrl(path string) string {
	return a.outbound("reference/" + path)
}

func (a httpAPI) getArchiveUrl(path string) string {
	return a.outbound("archive/" + path)
}

func get[T any](log *logrus.Entry, ctx context.Context, client http.Client, headers Headers, params Params, url string) (out Response[T], err error) {
	return request[T](log, ctx, client, http.MethodGet, url, headers, params)
}

func request[T any](log *logrus.Entry, ctx context.Context, client http.Client, method, url string, headers Headers, params Params) (out Response[T], err error) {
	start := time.Now()
	var requestId string = getRequestId(ctx)
	request, err := params.AddIfMissing("request_id", requestId).Request(method, url)
	if err != nil {
		return out, logError(log, err, "failed to create request object")
	}
	defer func() {
		l := log.WithFields(logrus.Fields{
			"latency": time.Since(start).Milliseconds(),
			"url":     fmt.Sprintf("%s/%s", request.URL.Host, request.URL.Path),
			"method":  method,
			"params":  params,
		})
		if err != nil {
			l.Error(err)
		} else {
			l.Info()
		}
	}()
	headers[KiblRequestID] = requestId
	headers.Set(request)
	response, err := client.Do(request.WithContext(ctx))
	if err != nil {
		return out, logError(log, err, "error fetching response from API")
	} else if response.StatusCode >= 400 {
		return out, logError(log, errors.Errorf("response returned %d", response.StatusCode), "")
	} else if response.ContentLength == 0 {
		return
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return out, logError(log, err, "error ready response body from API")
	}
	defer response.Body.Close()

	out = Response[T]{}
	if err = json.Unmarshal(data, &out); err != nil {
		return out, logError(log, err, "failed to unmarshall response to %T", out)
	}
	return
}

func logError(log *logrus.Entry, err error, message string, args ...interface{}) error {
	if err != nil {
		log.WithError(err).Errorf(message, args...)
		return err
	}
	return nil
}

func getRequestId(ctx context.Context) string {
	val := ctx.Value(KiblRequestID)
	if val != nil {
		return val.(string)
	}
	return uuid.New().String()
}
