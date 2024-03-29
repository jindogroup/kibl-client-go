package models

type State struct {
	StateId    int64  `json:"state_id"`
	Name       string `json:"name"`
	Abrv       string `json:"abrv"`
	InsertedOn string `json:"inserted_on"`
}

type StateMapping struct {
	FeedStateMappingId int64  `json:"feed_state_mapping_id"`
	StateId            int64  `json:"state_id"`
	FeedSourceId       int64  `json:"feed_source_id"`
	FeedId             int64  `json:"feed_id"`
	FeedName           string `json:"feed_name"`
	FeedAbrv           string `json:"feed_abrv"`
	InsertedOn         string `json:"inserted_on"`
}
