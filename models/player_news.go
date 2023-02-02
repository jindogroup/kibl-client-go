package models

import (
	"strings"

	"github.com/antihax/optional"
)

type PlayerNews struct {
	ParticipantId           int64                     `json:"participant_id"`
	Name                    string                    `json:"name"`
	PlayerNewsId            int64                     `json:"player_news_id"`
	PlayerId                int64                     `json:"player_id"`
	InjuryLocationId        int64                     `json:"injury_location_id"`
	FixtureId               int64                     `json:"fixture_id"`
	Status                  string                    `json:"status"`
	IsImportant             bool                      `json:"is_important"`
	IsHighlighted           bool                      `json:"is_highlighted"`
	IsHealed                bool                      `json:"is_healed"`
	InsertedOn              string                    `json:"inserted_on"`
	Comment                 string                    `json:"comment"`
	Impact                  string                    `json:"impact"`
	Headline                string                    `json:"headline"`
	LineId                  int64                     `json:"line_id"`
	MoreName                string                    `json:"more_name"`
	MoreLink                string                    `json:"more_link"`
	LastUpdated             string                    `json:"last_updated"`
	Participants            []PlayerInjuryParticipant `json:"participant"`
	ParticipantExternalRefs []ParticipantExternalRef  `json:"participant_reference_external_ids,omitempty"`
}

type OptionalPlayerNewsInfoParams struct {
	ParticipantId         optional.Int64
	PlayerId              optional.Int64
	PlayerNewsId          optional.Int64
	SinceLastUpdated      optional.Time
	IncludeExternalIds    []string
	ExternalParticipantId *ParticipantMappingParam
}

func (p OptionalPlayerNewsInfoParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}

	if p.ParticipantId.IsSet() {
		out["participant_id"] = p.ParticipantId.Value()
	}
	if p.PlayerId.IsSet() {
		out["player_id"] = p.PlayerId.Value()
	}
	if p.PlayerNewsId.IsSet() {
		out["player_news_id"] = p.PlayerNewsId.Value()
	}
	if p.SinceLastUpdated.IsSet() {
		out["since_last_updated"] = p.SinceLastUpdated.Value().UTC()
	}
	if len(p.IncludeExternalIds) > 0 {
		out["include_external_ids"] = strings.Join(p.IncludeExternalIds, ",")
	}
	if p.ExternalParticipantId != nil {
		out["external_participant_id"] = p.ExternalParticipantId.String()
	}
	return
}
