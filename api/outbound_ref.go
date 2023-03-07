package api

import (
	"context"
	"strings"

	"github.com/jindogroup/kibl-client-go/models"
	"github.com/jindogroup/kibl-client-go/utils"
)

var _ OutboundReferenceService = &httpAPI{}

func (a *httpAPI) GetCombinedLineTypesRef(ctx context.Context) (Response[models.CombinedLineType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.CombinedLineType](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("combined-line-types"))
}

func (a *httpAPI) GetFeedSourcesRef(ctx context.Context, parms *models.OptionalFeedSourceParams) (Response[models.FeedSource], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.FeedSource](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("feed-sources"))
}

func (a *httpAPI) GetFeedSourceTypesRef(ctx context.Context) (Response[models.FeedSourceType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.FeedSourceType](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("feed-types"))
}

func (a *httpAPI) GetFixtureTypesRef(ctx context.Context) (Response[models.FixtureType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.FixtureType](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("fixture-types"))
}

func (a *httpAPI) GetInjuryLocationsRef(ctx context.Context) (Response[models.InjuryLocation], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.InjuryLocation](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("injury-locations"))
}

func (a *httpAPI) GetLeaguesRef(ctx context.Context, params *models.OptionalLeagueParams) (Response[models.League], error) {
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.League](true)(a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("leagues"))
}

func (a *httpAPI) GetLeaguesSettings(ctx context.Context, params *models.OptionalLeagueParams) (Response[models.LeagueSetting], error) {
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.LeagueSetting](true)(a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("leagues-settings"))
}

func (a *httpAPI) GetLocationsRef(ctx context.Context) (Response[models.Location], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Location](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("locations"))
}

func (a *httpAPI) GetMarketStatusesRef(ctx context.Context) (Response[models.MarketStatus], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.MarketStatus](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("market-statuses"))
}

func (a *httpAPI) GetMarketTypesRef(ctx context.Context) (Response[models.MarketType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.MarketType](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("market-types"))
}

func (a *httpAPI) GetParticipantTypesRef(ctx context.Context) (Response[models.ParticipantType], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.ParticipantType](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("participant-types"))
}

func (a *httpAPI) GetParticipantsRef(ctx context.Context, leagueIds []int64, params *models.OptionalParticipantParams) (res Response[models.Participant], err error) {
	_params := Params{}
	if len(leagueIds) > 0 {
		_params["league_id"] = strings.Join(utils.String.FromInt64Slice(leagueIds), ",")
	}
	if params != nil {
		_params.Merge(params.Params())
	}
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Participant](true)(a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("participants"))
}

func (a *httpAPI) GetParticipantsAliasesRef(ctx context.Context, leagueId int64, params *models.OptionalParticipantParams) (Response[models.ParticipantAlias], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.ParticipantAlias](true)(a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("participant-aliases"))
}

func (a *httpAPI) GetPlayersRef(ctx context.Context, leagueIds []int64, params *models.OptionalPlayerParams) (Response[models.Player], error) {
	_params := Params{}
	if len(leagueIds) > 0 {
		_params["league_id"] = strings.Join(utils.String.FromInt64Slice(leagueIds), ",")
	}
	if params != nil {
		_params.Merge(params.Params())
	}
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Player](true)(a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("players"))
}

func (a *httpAPI) GetRegionsRef(ctx context.Context) (Response[models.Region], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Region](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("region"))
}

func (a *httpAPI) GetSeasonsRef(ctx context.Context) (Response[models.Season], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Season](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("seasons"))
}

func (a *httpAPI) GetSegmentsRef(ctx context.Context) (Response[models.Segment], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Segment](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("segments"))
}

func (a *httpAPI) GetSidesRef(ctx context.Context, params *models.OptionalSideParams) (Response[models.Side], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.Side](true)(a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("sides"))
}

func (a *httpAPI) GetSportsRef(ctx context.Context, params *models.OptionalSportsParams) (Response[models.Sport], error) {
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Sport](true)(a.log, ctx, a.client, a.headers(), _params, a.getReferenceUrl("sports"))
}

func (a *httpAPI) GetStatesRef(ctx context.Context) (Response[models.State], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.State](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("states"))
}

func (a *httpAPI) GetSportsbooksRef(ctx context.Context) (Response[models.Sportsbook], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	return get[models.Sportsbook](true)(a.log, ctx, a.client, a.headers(), Params{}, a.getReferenceUrl("sportsbooks"))
}
