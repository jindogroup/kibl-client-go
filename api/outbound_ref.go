package api

import (
	"context"
	"github.com/jindogroup/kibl-client-go/models"
)

var _ OutboundReferenceService = &httpAPI{}

func (a *httpAPI) GetCombinedLineTypesRef(ctx context.Context) (Response[models.CombinedLineType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.CombinedLineType](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("combined-line-types"))
}
func (a *httpAPI) GetFeedSourcesRef(ctx context.Context, parms *models.OptionalFeedSourceParams) (Response[models.FeedSource], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.FeedSource](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("feed-sources"))
}
func (a *httpAPI) GetFeedSourceTypesRef(ctx context.Context) (Response[models.FeedSourceType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.FeedSourceType](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("feed-types"))
}
func (a *httpAPI) GetFixtureTypesRef(ctx context.Context) (Response[models.FixtureType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.FixtureType](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("fixture-types"))
}
func (a *httpAPI) GetInjuryLocationsRef(ctx context.Context) (Response[models.InjuryLocation], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.InjuryLocation](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("injury-locations"))
}
func (a *httpAPI) GetLeaguesRef(ctx context.Context) (Response[models.League], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.League](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("leagues"))
}
func (a *httpAPI) GetLocationsRef(ctx context.Context) (Response[models.Location], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Location](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("locations"))
}
func (a *httpAPI) GetMarketStatusesRef(ctx context.Context) (Response[models.MarketStatus], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.MarketStatus](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("market-statuses"))
}
func (a *httpAPI) GetMarketTypesRef(ctx context.Context) (Response[models.MarketType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.MarketType](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("market-types"))
}
func (a *httpAPI) GetParticipantTypesRef(ctx context.Context) (Response[models.ParticipantType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.ParticipantType](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("participant-types"))
}
func (a *httpAPI) GetParticipantsRef(ctx context.Context, leagueId int64, params *models.OptionalParticipantParams) (Response[models.Participant], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.Participant](a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("participants"))
}
func (a *httpAPI) GetParticipantsAliasesRef(ctx context.Context, leagueId int64, params *models.OptionalParticipantParams) (Response[models.ParticipantAlias], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.ParticipantAlias](a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("participant-aliases"))
}
func (a *httpAPI) GetPlayersRef(ctx context.Context, leagueId int64, params *models.OptionalPlayerParams) (Response[models.Player], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.Player](a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("players"))
}
func (a *httpAPI) GetRegionRef(ctx context.Context) (Response[models.Region], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Region](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("regions"))
}
func (a *httpAPI) GetSeasonsRef(ctx context.Context) (Response[models.Season], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Season](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("seasons"))
}
func (a *httpAPI) GetSegmentsRef(ctx context.Context) (Response[models.Segment], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Segment](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("segments"))
}
func (a *httpAPI) GetSidesRef(ctx context.Context, params *models.OptionalSideParams) (Response[models.Side], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.Side](a.log, ctx, a.client, a.headers(), _params, a.getMappingUrl("sides"))
}
func (a *httpAPI) GetSportsRef(ctx context.Context) (Response[models.Sport], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Sport](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("sports"))
}
func (a *httpAPI) GetStatesRef(ctx context.Context) (Response[models.State], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.State](a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("states"))
}
