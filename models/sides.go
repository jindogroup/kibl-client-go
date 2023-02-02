package models

import "github.com/antihax/optional"

type Side struct {
	SideId     int64  `json:"side_id"`
	Name       string `json:"name"`
	Abrv       string `json:"abrv"`
	InsertedOn string `json:"inserted_on"`
}

type OptionalSideParams struct {
	SideId optional.Int
}

func (p OptionalSideParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}

	if p.SideId.IsSet() {
		out["side_id"] = p.SideId.Value()
	}
	return
}

type SideExternalRef struct {
	SideId       int64  `json:"side_id"`
	FeedSourceId int64  `json:"feed_source_id"`
	FeedId       int64  `json:"feed_id"`
	FeedName     string `json:"feed_name"`
	FeedAbrv     string `json:"feed_abrv"`
}

type SideMapping struct {
	FeedSideMappingId int64  `json:"feed_side_mapping_id"`
	SideId            int64  `json:"side_id"`
	FeedSourceId      int64  `json:"feed_source_id"`
	FeedId            int64  `json:"feed_id"`
	FeedName          string `json:"feed_name"`
	FeedAbrv          string `json:"feed_abrv"`
	InsertedOn        string `json:"inserted_on"`
}
