package api

import (
	"context"
	"time"

	"github.com/jindogroup/kibl-client-go/models"
)

var _ OutboundArchiveService = &httpAPI{}

func (a *httpAPI) GetArchivedPlayerNews(ctx context.Context, sinceLastUpdated time.Time) (Response[models.PlayerNews], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"since_last_updated": sinceLastUpdated.UTC().Format(KiblTimestampFormat)}
	return get[models.PlayerNews](a.log, ctx, a.client, a.headers(), params, a.getArchiveUrl("player-news"))
}
