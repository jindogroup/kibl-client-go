package api

import (
	"context"
	"strconv"

	"github.com/jindogroup/kibl-client-go/models"
)

func (a *httpAPI) GetFixtureTypesMapping(ctx context.Context, feedSourceId int64) (Response[models.FixtureTypeMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.FixtureTypeMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("fixture-type"))
}

func (a *httpAPI) GetFixturesMapping(ctx context.Context, feedSourceId int64, params *models.OptionalFixtureMappingParams) (Response[models.FixtureMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.FixtureMapping](a.log, ctx, a.client, a.headers(), _params, a.getMappingUrl("fixture"))
}

func (a *httpAPI) GetLeaguesMapping(ctx context.Context, feedSourceId int64) (Response[models.LeagueMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.LeagueMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("league"))
}

func (a *httpAPI) GetLocationsMapping(ctx context.Context, feedSourceId int64) (Response[models.LocationMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.LocationMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("location"))
}

func (a *httpAPI) GetMarketStatusMapping(ctx context.Context, feedSourceId int64) (Response[models.MarketStatusMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.MarketStatusMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("market-mtatus"))
}

func (a *httpAPI) GetMarketTypesMapping(ctx context.Context, feedSourceId int64) (Response[models.MarketTypeMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.MarketTypeMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("market-type"))
}

func (a *httpAPI) GetParticipantsMapping(ctx context.Context, feedSourceId int64, optVars *models.OptionalParticipantMappingParams) (Response[models.ParticipantMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.ParticipantMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("participants"))
}

func (a *httpAPI) GetPlayersMapping(ctx context.Context, feedSourceId int64) (Response[models.PlayerMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.PlayerMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("players"))
}

func (a *httpAPI) GetSeasonsMapping(ctx context.Context, feedSourceId int64) (Response[models.SeasonMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.SeasonMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("seasons"))
}

func (a *httpAPI) GetSegmentsMapping(ctx context.Context, feedSourceId int64) (Response[models.SegmentMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.SegmentMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("segments"))
}

func (a *httpAPI) GetSidesMapping(ctx context.Context, feedSourceId int64) (Response[models.SideMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.SideMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("sides"))
}

func (a *httpAPI) GetSportsMapping(ctx context.Context, feedSourceId int64) (Response[models.SportMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.SportMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("sports"))
}

func (a *httpAPI) GetStatesMapping(ctx context.Context, feedSourceId int64) (Response[models.StateMapping], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	params := Params{"feed_source_id": strconv.Itoa(int(feedSourceId))}
	return get[models.StateMapping](a.log, ctx, a.client, a.headers(), params, a.getMappingUrl("states"))
}
