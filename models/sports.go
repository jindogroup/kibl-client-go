package models

import "github.com/antihax/optional"

type OptionalSportsParams struct {
	SportId optional.Int64
}

func (p OptionalSportsParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}

	if p.SportId.IsSet() {
		out["sport_id"] = p.SportId.Value()
	}
	return
}

type Sport struct {
	SportId    int64  `json:"sport_id"`
	Name       string `json:"name"`
	Abrv       string `json:"abrv"`
	InsertedOn string `json:"inserted_on"`
}

type SportMapping struct {
	FeedSportMappingId int64  `json:"feed_sport_mapping_id"`
	SportId            int64  `json:"sport_id"`
	FeedSourceId       int64  `json:"feed_source_id"`
	FeedId             int64  `json:"feed_id"`
	FeedName           string `json:"feed_name"`
	FeedAbrv           string `json:"feed_abrv"`
	IsActive           bool   `json:"is_active"`
	InsertedOn         string `json:"inserted_on"`
}
