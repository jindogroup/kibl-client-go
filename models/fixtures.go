package models

import (
	"strings"

	"github.com/antihax/optional"
	"github.com/jindogroup/kibl-client-go/utils"
)

type Fixture struct {
	FixtureId           int64                  `json:"fixture_id"`
	LeagueId            int64                  `json:"league_id"`
	SeasonId            int64                  `json:"season_id"`
	LocationId          int64                  `json:"location_id"`
	FixtureTypeId       int64                  `json:"fixture_type_id"`
	StartTime           string                 `json:"start_time"`
	CutoffTime          string                 `json:"cutoff_time"`
	Name                string                 `json:"name"`
	InsertedOn          string                 `json:"inserted_on"`
	Participants        []FixturesParticipants `json:"participants"`
	Information         []FixtureInformation   `json:"informations,omitempty"`
	States              []State                `json:"states,omitempty"`
	Outcome             *Outcome               `json:"outcomes,omitempty"`
	FixtureExternalRefs []FixtureExternalRef   `json:"fixture_reference_external_ids,omitempty"`
}

type FixturesParticipants struct {
	FixtureParticipantId int64  `json:"fixture_participant_id"`
	FixtureId            int64  `json:"fixture_id"`
	ParticipantId        int64  `json:"participant_id"`
	SideId               int64  `json:"side_id"`
	Ordering             int64  `json:"ordering"`
	Rotation             int64  `json:"rotation"`
	InsertedOn           string `json:"inserted_on"`
}

type FixtureExternalRef struct {
	FixtureId    int64  `json:"fixture_id"`
	FeedSourceId int64  `json:"feed_source_id"`
	FeedId       int64  `json:"feed_id"`
	FeedName     string `json:"feed_name"`
}

type FixtureInformation struct {
	FixtureInformationId int64  `json:"fixture_information_id"`
	FixtureId            int64  `json:"fixture_id"`
	InformationId        int64  `json:"information_id"`
	Value                string `json:"value"`
	InsertedOn           string `json:"inserted_on"`
}

type OptionalFixtureInfoParams struct {
	FixtureIds []int64
}

func (p OptionalFixtureInfoParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if len(p.FixtureIds) > 0 {
		out["fixture_id"] = strings.Join(utils.String.FromInt64Slice(p.FixtureIds), ",")
	}
	return
}

type FixturesByRotation struct {
	Fixture
}

type FixturesByTeam struct {
	Fixture
}

type FixtureMapping struct {
	FeedFixtureMappingId int64  `json:"feed_fixture_mapping_id"`
	FixtureId            int64  `json:"fixture_id"`
	FeedSourceId         int64  `json:"feed_source_id"`
	FeedId               int64  `json:"feed_id"`
	FeedName             string `json:"feed_name"`
	InsertedOn           string `json:"inserted_on"`
}

type OptionalFixtureMappingParams struct {
	LeagueId  optional.Int64
	FixtureId optional.Int64
	StartTime optional.Time
}

func (p OptionalFixtureMappingParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if p.LeagueId.IsSet() {
		out["league_id"] = p.LeagueId.Value()
	}
	if p.FixtureId.IsSet() {
		out["fixture_id"] = p.FixtureId.Value()
	}
	if p.StartTime.IsSet() {
		out["start_time"] = p.StartTime.Value().UTC()
	}
	return
}

type OptionalFixturesInfoParams struct {
	FixtureIds         []int64
	StartTime          optional.Time
	EndTime            optional.Time
	IncludeExternalIds []string
	ExternalLeagueId   *LeagueMappingParam
}

func (p OptionalFixturesInfoParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if len(p.FixtureIds) > 0 {
		out["fixture_id"] = strings.Join(utils.String.FromInt64Slice(p.FixtureIds), ",")
	}
	if p.StartTime.IsSet() {
		out["start_time"] = p.StartTime.Value().UTC().Format(QueryTimeFormat)
	}
	if p.EndTime.IsSet() {
		out["end_time"] = p.EndTime.Value().UTC().Format(QueryTimeFormat)
	}
	if len(p.IncludeExternalIds) > 0 {
		out["include_external_ids"] = strings.Join(p.IncludeExternalIds, ",")
	}
	if p.ExternalLeagueId != nil {
		out["external_league_id"] = p.ExternalLeagueId.String()
	}
	return
}

type OptionalFixturesInfoByRotationParams struct {
	IncludeExternalIds []string
}

func (p OptionalFixturesInfoByRotationParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if len(p.IncludeExternalIds) > 0 {
		out["include_external_ids"] = strings.Join(p.IncludeExternalIds, ",")
	}
	return
}

type OptionalFixturesInfoByTeamsParams struct {
	IncludeExternalIds []string
}

func (p OptionalFixturesInfoByTeamsParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if len(p.IncludeExternalIds) > 0 {
		out["include_external_ids"] = strings.Join(p.IncludeExternalIds, ",")
	}
	return
}

type FixtureType struct {
	FixtureTypeId int64  `json:"fixture_type_id"`
	Name          string `json:"name"`
	Abrv          string `json:"abrv"`
	InsertedOn    string `json:"inserted_on"`
}

type FixtureTypeMapping struct {
	FeedFixtureTypeMappingId int64  `json:"feed_fixture_type_mapping_id"`
	FixtureTypeId            int64  `json:"fixture_type_id"`
	FeedSourceId             int64  `json:"feed_source_id"`
	FeedId                   int64  `json:"feed_id"`
	FeedName                 string `json:"feed_name"`
	FeedAbrv                 string `json:"feed_abrv"`
	InsertedOn               string `json:"inserted_on"`
}
