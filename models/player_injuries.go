package models

import (
	"strings"

	"github.com/antihax/optional"
)

type PlayerInjury struct {
	ParticipantId          int64                     `json:"participant_id"`
	Name                   string                    `json:"name"`
	PlayerInjuryId         int64                     `json:"player_injury_id"`
	PlayerId               int64                     `json:"player_id"`
	InjuryLocationId       int64                     `json:"injury_location_id"`
	FixtureId              int64                     `json:"fixture_id"`
	Status                 string                    `json:"status"`
	InsertedOn             string                    `json:"inserted_on"`
	Injury                 string                    `json:"injury"`
	ReturnEstimate         string                    `json:"return_estimate"`
	StartDate              string                    `json:"start_date"`
	LineId                 int64                     `json:"line_id"`
	Participant            []PlayerInjuryParticipant `json:"participant"`
	ParticipantExternalRef []ParticipantExternalRef  `json:"participant_reference_external_ids,omitempty"`
}

type OptionalPlayerInjuryInfoParams struct {
	ParticipantId         optional.Int64
	PlayerId              optional.Int64
	IncludeExternalIds    []string
	ExternalParticipantId *ParticipantMappingParam
}

func (p OptionalPlayerInjuryInfoParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}

	if p.ParticipantId.IsSet() {
		out["participant_id"] = p.ParticipantId.Value()
	}
	if p.PlayerId.IsSet() {
		out["player_id"] = p.PlayerId.Value()
	}
	if len(p.IncludeExternalIds) > 0 {
		out["include_external_ids"] = strings.Join(p.IncludeExternalIds, ",")
	}
	if p.ExternalParticipantId != nil {
		out["external_participant_id"] = p.ExternalParticipantId.String()
	}
	return
}

type PlayerInjuryParticipant struct {
	ParticipantId     int64  `json:"participant_id"`
	LeagueId          int64  `json:"league_id"`
	ParticipantTypeId int64  `json:"participant_type_id"`
	Name              string `json:"name"`
	Abrv              string `json:"abrv"`
	Mascot            string `json:"mascot"`
	InsertedOn        string `json:"inserted_on"`
}

type InjuryLocation struct {
	InjuryLocationId int64  `json:"injury_location_id"`
	Name             string `json:"name"`
	Abrv             string `json:"abrv"`
	InsertedOn       string `json:"inserted_on"`
}
