package models

import (
	"fmt"

	"github.com/antihax/optional"
)

type Participant struct {
	ParticipantId     int64  `json:"participant_id"`
	LeagueId          int64  `json:"league_id"`
	ParticipantTypeId int64  `json:"participant_type_id"`
	Name              string `json:"name"`
	Abrv              string `json:"abrv"`
	Mascot            string `json:"mascot"`
	InsertedOn        string `json:"inserted_on"`
}

type OptionalParticipantMappingParams struct {
	LeagueId optional.Int
}

type ParticipantType struct {
	ParticipantTypeId int64  `json:"participant_type_id"`
	Name              string `json:"name"`
	Abrv              string `json:"abrv"`
	InsertedOn        string `json:"inserted_on"`
}

type ParticipantAlias struct {
	ParticipantAliasId int64  `json:"participant_alias_id"`
	ParticipantId      int64  `json:"participant_id"`
	Name               string `json:"name"`
	InsertedOn         string `json:"inserted_on"`
}

type OptionalParticipantParams struct {
	ParticipantId optional.Int
}

func (p OptionalParticipantParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if p.ParticipantId.IsSet() {
		out["participant_id"] = p.ParticipantId.Value()
	}
	return
}

type OptionalParticipantInfoParams struct {
	FixtureId optional.Int
}

func (p OptionalParticipantInfoParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if p.FixtureId.IsSet() {
		out["fixture_id"] = p.FixtureId.Value()
	}
	return
}

type ParticipantExternalRef struct {
	ParticipantID int    `json:"participant_id,omitempty"`
	FeedSourceID  int    `json:"feed_source_id,omitempty"`
	FeedID        int    `json:"feed_id,omitempty"`
	FeedName      string `json:"feed_name,omitempty"`
	FeedNameShort string `json:"feed_name_short,omitempty"`
	FeedMascot    string `json:"feed_mascot,omitempty"`
	FeedAbrv      string `json:"feed_abrv,omitempty"`
}

type ParticipantMapping struct {
	FeedParticipantMappingId int64  `json:"feed_participant_mapping_id"`
	ParticipantId            int64  `json:"participant_id"`
	FeedSourceId             int64  `json:"feed_source_id"`
	FeedId                   int64  `json:"feed_id"`
	FeedName                 string `json:"feed_name"`
	FeedAbrv                 string `json:"feed_abrv"`
	InsertedOn               string `json:"inserted_on"`
}

type ParticipantMappingParam struct {
	ParticipantId int64 `json:"league_id,omitempty"`
	FeedSourceId  int64 `json:"feed_source_id"`
}

func (p *ParticipantMappingParam) String() string {
	return fmt.Sprintf("%d:%d", p.FeedSourceId, p.ParticipantId)
}
