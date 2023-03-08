package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	_cache _xch
	_sConf map[string]_stCf
)

const (
	cacheExpiry time.Duration = 10 * time.Minute
	cachePurge  time.Duration = 15 * time.Minute
)

func hour(m int) time.Duration {
	return time.Duration(m) * time.Hour
}

func minute(m int) time.Duration {
	return time.Duration(m) * time.Minute
}

func init() {
	_sConf = map[string]_stCf{
		"sports/get/info":      {minute(1), minute(3)},
		"sports/get/reference": {minute(30), minute(43)},
		"sports/get/mapping":   {minute(1), minute(3)},
		"sports/get/archive":   {minute(2), minute(5)},
	}
	_cache = _xch{
		stores: make(map[string]*_st),
	}
}

type _stCf struct {
	expiry time.Duration
	purge  time.Duration
}

type _st struct {
	mx   sync.RWMutex
	data *cache.Cache
}

func (s *_st) key(params map[string]string) string {
	if params == nil {
		params = map[string]string{}
	}
	b, _ := json.Marshal(params)
	return strings.Replace(string(b), fmt.Sprintf("\"request_id\":\"%s\"", params["request_id"]), "", 1)
}

func (s *_st) get(params map[string]string) (out interface{}, err error, ok bool) {
	key := s.key(params)
	s.mx.RLock()
	defer s.mx.RUnlock()
	out, ok = s.data.Get(key)
	return
}

func (s *_st) set(params map[string]string, fetch func() (interface{}, error)) (out interface{}, err error) {
	var ok bool
	key := s.key(params)
	s.mx.RLock()
	defer s.mx.RUnlock()
	if out, ok = s.data.Get(key); ok {
		return
	} else if out, err = fetch(); err == nil {
		s.data.Set(key, out, cache.DefaultExpiration)
	}
	return
}

type _xch struct {
	mx     sync.RWMutex
	stores map[string]*_st
}

func (c *_xch) limits(path string) (expiry, purge time.Duration) {
	key := strings.Trim(path[:strings.LastIndex(path, "/")], "/")
	if cnf, ok := _sConf[key]; ok {
		return cnf.expiry, cnf.purge
	}
	return cacheExpiry, cachePurge
}

func (c *_xch) get(log *logrus.Entry, path string) (store *_st) {
	c.mx.RLock()
	var ok bool
	if store, ok = c.stores[path]; ok {
		c.mx.RUnlock()
		return store
	}
	c.mx.RUnlock()
	c.mx.Lock()
	defer c.mx.Unlock()
	if store, ok = c.stores[path]; ok {
		return store
	}
	expiry, purge := c.limits(path)
	log.WithFields(logrus.Fields{
		"store":           path,
		"purge-interval":  purge.Round(time.Minute),
		"expiry-interval": expiry.Round(time.Minute),
	}).Debug("creating store")
	store = &_st{
		data: cache.New(expiry, purge),
	}
	c.stores[path] = store
	return
}

func cacheGet[T any](log *logrus.Entry, ctx context.Context, client http.Client, headers Headers, params Params, _url string) (out Response[T], err error) {
	var ok bool
	var res interface{}
	var _u *url.URL

	if _url == "" {
		return out, errors.New("invalid blank url")
	} else if _u, err = url.Parse(_url); err != nil {
		return out, errors.Wrap(err, "failed to parse url")
	}

	store := _cache.get(log, _u.Path)
	start := time.Now()
	l := log.WithFields(logrus.Fields{
		"url":    fmt.Sprintf("%s://%s%s", _u.Scheme, _u.Host, _u.Path),
		"params": params, "cache": _u.Path,
	})
	if res, err, ok = store.get(params); err != nil {
		l.WithField("latency", time.Since(start).Milliseconds()).Error(err)
		return
	} else if ok {
		l.WithField("latency", time.Since(start).Milliseconds()).Info()
		return *res.(*Response[T]), nil
	}
	defer func() {
		lg := l.WithField("latency", time.Since(start).Milliseconds())
		if err != nil {
			lg.Error(err)
		} else {
			lg.Info()
		}
	}()
	if res, err = store.set(params, func() (interface{}, error) {
		rs, e := request[T](log, ctx, client, http.MethodGet, _url, headers, params)
		if err == nil {
			t := time.Now().UTC()
			rs.CachedAt = &t
		}
		return &rs, e
	}); err == nil {
		return *res.(*Response[T]), nil

	}
	return
}
