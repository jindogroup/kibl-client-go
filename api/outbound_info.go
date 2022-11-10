package api

import (
	"context"
	"strconv"
	"time"

	"github.com/jindogroup/kibl-client-go/models"
)

func (a *httpAPI) GetFixturesInfo(ctx context.Context, leagueId int64, params *models.OptionalFixturesInfoParams) (Response[models.Fixture], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{"league_id": strconv.Itoa(int(leagueId))}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.Fixture](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("fixtures"))
}

func (a *httpAPI) GetFixturesByRotationsInfo(ctx context.Context, rotation int64, startTime, endTime time.Time, params *models.OptionalFixturesInfoByRotationParams) (Response[models.FixturesByRotation], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{
		"rotations":  strconv.Itoa(int(rotation)),
		"start_time": startTime.UTC().Format(KiblTimestampFormat),
		"end_time":   endTime.UTC().Format(KiblTimestampFormat),
	}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.FixturesByRotation](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("fixtures-by-rotations"))
}

func (a *httpAPI) GetFixturesByTeamsInfo(ctx context.Context, leagueId int64, participantId1, participantId2 int64, startTime, endTime time.Time, params *models.OptionalFixturesInfoByTeamsParams) (Response[models.FixturesByTeam], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{
		"league_id":     strconv.Itoa(int(leagueId)),
		"participant_1": strconv.Itoa(int(participantId1)),
		"participant_2": strconv.Itoa(int(participantId2)),
		"start_time":    startTime.UTC().Format(KiblTimestampFormat),
		"end_time":      endTime.UTC().Format(KiblTimestampFormat),
	}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.FixturesByTeam](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("fixtures-by-teams"))
}

func (a *httpAPI) GetInformationInfo(ctx context.Context, params *models.OptionalFixtureInfoParams) (Response[models.FixtureInformation], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.FixtureInformation](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("informations"))
}

func (a *httpAPI) GetMarketsInfo(ctx context.Context, leagueId int64, params *models.OptionalMarketInfoParams) (Response[models.Market], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{"league_id": strconv.Itoa(int(leagueId))}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.Market](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("markets"))
}

func (a *httpAPI) GetOutcomeParticipantsInfo(ctx context.Context, params *models.OptionalOutcomeParticipantInfoParams) (Response[models.OutcomeParticipant], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.OutcomeParticipant](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("outcome-participants"))
}

func (a *httpAPI) GetOutcomesInfo(ctx context.Context, params *models.OptionalOutcomeInfoParams) (Response[models.Outcome], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.Outcome](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("outcomes"))
}

func (a *httpAPI) GetParticipantsInfo(ctx context.Context, params *models.OptionalParticipantInfoParams) (Response[models.Participant], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.Participant](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("participants"))
}

func (a *httpAPI) GetPlayerInjuriesInfo(ctx context.Context, params *models.OptionalPlayerInjuryInfoParams) (Response[models.PlayerInjury], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.PlayerInjury](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("player-injuries"))
}

func (a *httpAPI) GetPlayerNewsInfo(ctx context.Context, params *models.OptionalPlayerNewsInfoParams) (Response[models.PlayerNews], error) {
	ctx, cancel := context.WithTimeout(ctx, a.requestTimeout)
	defer cancel()
	_params := Params{}
	if params != nil {
		_params.Merge(params.Params())
	}
	return get[models.PlayerNews](a.log, ctx, a.client, a.headers(), _params, a.getInfoUrl("player-news"))
}
