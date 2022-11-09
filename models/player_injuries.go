/*
 * Kibl Sports API - Outbound
 *
 * Kibl Sports API - Outbound
 *
 * API version: 1.0.0
 * Contact: developers@jindogroup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

import "github.com/antihax/optional"

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
