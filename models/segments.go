package models

import "time"

type Segment struct {
	SegmentId  int64  `json:"segment_id"`
	Name       string `json:"name"`
	Abrv       string `json:"abrv"`
	InsertedOn string `json:"inserted_on"`
}

type SegmentMapping struct {
	FeedSegmentMappingId int64  `json:"feed_segment_mapping_id"`
	SegmentId            int64  `json:"segment_id"`
	FeedSourceId         int64  `json:"feed_source_id"`
	FeedId               int64  `json:"feed_id"`
	FeedName             string `json:"feed_name"`
	FeedAbrv             string `json:"feed_abrv"`
	InsertedOn           string `json:"inserted_on"`
}

type OutcomeSegmentScore struct {
	FeedSegmentMappingId int64     `json:"outcome_segment_score_id"`
	FixtureParticipantId int64     `json:"fixture_participant_id"`
	FixtureId            int64     `json:"fixture_id"`
	SegmentId            int64     `json:"segment_id"`
	Score                int       `json:"score"`
	InsertedOn           time.Time `json:"inserted_on"`
}
