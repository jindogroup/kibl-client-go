package api

import (
	"context"
	"time"

	"github.com/jindogroup/kibl-client-go/models"
)

type Service interface {
	OutboundService
}

type OutboundService interface {
	OutboundArchiveService
	OutboundMappingService
	OutboundReferenceService
	OutboundSportsInformationService
}

type OutboundArchiveService interface {
	GetArchivedPlayerNews(ctx context.Context, sinceLastUpdated time.Time) (Response[models.PlayerNews], error)
}

type OutboundMappingService interface {
	GetFixtureTypesMapping(ctx context.Context, feedSourceId int64) (Response[models.FixtureTypeMapping], error)
	GetFixturesMapping(ctx context.Context, feedSourceId int64, params *models.OptionalFixtureMappingParams) (Response[models.FixtureMapping], error)
	GetLeaguesMapping(ctx context.Context, feedSourceId int64) (Response[models.LeagueMapping], error)
	GetLocationsMapping(ctx context.Context, feedSourceId int64) (Response[models.LocationMapping], error)
	GetMarketStatusMapping(ctx context.Context, feedSourceId int64) (Response[models.MarketStatusMapping], error)
	GetMarketTypesMapping(ctx context.Context, feedSourceId int64) (Response[models.MarketTypeMapping], error)
	GetParticipantsMapping(ctx context.Context, feedSourceId int64, optVars *models.OptionalParticipantMappingParams) (Response[models.ParticipantMapping], error)
	GetPlayersMapping(ctx context.Context, feedSourceId int64) (Response[models.PlayerMapping], error)
	GetSeasonsMapping(ctx context.Context, feedSourceId int64) (Response[models.SeasonMapping], error)
	GetSegmentsMapping(ctx context.Context, feedSourceId int64) (Response[models.SegmentMapping], error)
	GetSidesMapping(ctx context.Context, feedSourceId int64) (Response[models.SideMapping], error)
	GetSportsMapping(ctx context.Context, feedSourceId int64) (Response[models.SportMapping], error)
	GetStatesMapping(ctx context.Context, feedSourceId int64) (Response[models.StateMapping], error)
}

type OutboundReferenceService interface {
	GetCombinedLineTypesRef(ctx context.Context) (Response[models.CombinedLineType], error)
	GetFeedSourcesRef(ctx context.Context, parms *models.OptionalFeedSourceParams) (Response[models.FeedSource], error)
	GetFeedSourceTypesRef(ctx context.Context) (Response[models.FeedSourceType], error)
	GetFixtureTypesRef(ctx context.Context) (Response[models.FixtureType], error)
	GetInjuryLocationsRef(ctx context.Context) (Response[models.InjuryLocation], error)
	GetLeaguesRef(ctx context.Context) (Response[models.League], error)
	GetLocationsRef(ctx context.Context) (Response[models.Location], error)
	GetMarketStatusesRef(ctx context.Context) (Response[models.MarketStatus], error)
	GetMarketTypesRef(ctx context.Context) (Response[models.MarketType], error)
	GetParticipantTypesRef(ctx context.Context) (Response[models.ParticipantType], error)
	GetParticipantsRef(ctx context.Context, leagueId int64, params *models.OptionalParticipantParams) (Response[models.Participant], error)
	GetParticipantsAliasesRef(ctx context.Context, leagueId int64, params *models.OptionalParticipantParams) (Response[models.ParticipantAlias], error)
	GetPlayersRef(ctx context.Context, leagueId int64, params *models.OptionalPlayerParams) (Response[models.Player], error)
	GetRegionRef(ctx context.Context) (Response[models.Region], error)
	GetSeasonsRef(ctx context.Context) (Response[models.Season], error)
	GetSegmentsRef(ctx context.Context) (Response[models.Segment], error)
	GetSidesRef(ctx context.Context, params *models.OptionalSideParams) (Response[models.Side], error)
	GetSportsRef(ctx context.Context) (Response[models.Sport], error)
	GetStatesRef(ctx context.Context) (Response[models.State], error)
}

type OutboundSportsInformationService interface {
	GetFixturesInfo(ctx context.Context, leagueId string, params *models.OptionalFixturesInfoParams) (Response[models.Fixture], error)
	GetFixturesByRotationsInfo(ctx context.Context, rotations int64, startTime, endTime time.Time, params *models.OptionalFixturesInfoByRotationParams) (Response[models.FixturesByRotation], error)
	GetFixturesByTeamsInfo(ctx context.Context, leagueId int64, participantId1, participantId2 int64, startTime, endTime time.Time, params *models.OptionalFixturesInfoByTeamsParams) (Response[models.FixturesByTeam], error)
	GetInformationInfo(ctx context.Context, params *models.OptionalFixtureInfoParams) (Response[models.FixtureInformation], error)
	GetMarketsInfo(ctx context.Context, leagueId int64, params *models.OptionalMarketInfoParams) (Response[models.Market], error)
	GetOutcomeParticipantsInfo(ctx context.Context, params *models.OptionalOutcomeParticipantInfoParams) (Response[models.OutcomeParticipant], error)
	GetOutcomesInfo(ctx context.Context, params *models.OptionalOutcomeInfoParams) (Response[models.Outcome], error)
	GetParticipantsInfo(ctx context.Context, params *models.OptionalParticipantInfoParams) (Response[models.Participant], error)
	GetPlayerInjuriesInfo(ctx context.Context, params *models.OptionalPlayerInjuryInfoParams) (Response[models.PlayerInjury], error)
	GetPlayerNewsInfo(ctx context.Context, params *models.OptionalPlayerNewsInfoParams) (Response[models.PlayerNews], error)
}
