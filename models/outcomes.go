package models

import "github.com/antihax/optional"

type Outcome struct {
	OutcomeId  int64  `json:"outcome_id"`
	FixtureId  int64  `json:"fixture_id"`
	SegmentId  int64  `json:"segment_id"`
	Clock      string `json:"clock"`
	IsFinal    bool   `json:"is_final"`
	InsertedOn string `json:"inserted_on"`
}

type OptionalOutcomeParticipantInfoParams struct {
	OutcomeId optional.Int64
}

func (p OptionalOutcomeParticipantInfoParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if p.OutcomeId.IsSet() {
		out["outcome_id"] = p.OutcomeId.Value()
	}
	return
}

type OptionalOutcomeInfoParams struct {
	FixtureId optional.Int64
}

func (p OptionalOutcomeInfoParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if p.FixtureId.IsSet() {
		out["fixture_id"] = p.FixtureId.Value()
	}
	return
}

type OutcomeParticipant struct {
	OutcomeParticipantId int64  `json:"outcome_participant_id"`
	OutcomeId            int64  `json:"outcome_id"`
	FixtureParticipantId int64  `json:"fixture_participant_id"`
	Score                string `json:"score"`
	IsWinner             bool   `json:"is_winner"`
	InsertedOn           string `json:"inserted_on"`
}
