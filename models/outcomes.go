package models

import "github.com/antihax/optional"

type Outcome struct {
	OutcomeID      int64                 `json:"outcome_id"`
	FixtureID      int64                 `json:"fixture_id"`
	SegmentID      int                   `json:"segment_id"`
	Clock          string                `json:"clock"`
	IsCurrent      bool                  `json:"is_current"`
	IsStartGame    bool                  `json:"is_start_game"`
	IsEndGame      bool                  `json:"is_end_game"`
	IsStartSegment bool                  `json:"is_start_segment"`
	IsEndSegment   bool                  `json:"is_end_segment"`
	InsertedOn     string                `json:"inserted_on"`
	RoutingKey     string                `json:"routing_key"`
	Participants   []OutcomeParticipant  `json:"outcomes_participants"`
	Scores         []OutcomeSegmentScore `json:"segment_scores"`
	States         []OutcomeState        `json:"states"`
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
	Score                int    `json:"score"`
	IsWinner             bool   `json:"is_winner"`
	InsertedOn           string `json:"inserted_on"`
}
