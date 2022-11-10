package models

import "github.com/antihax/optional"

type FeedSource struct {
	FeedSourceId int64  `json:"feed_source_id"`
	Name         string `json:"name"`
	Tag          string `json:"tag"`
	FeedTypeId   int64  `json:"feed_type_id"`
	FeedId       int64  `json:"feed_id"`
	FeedName     string `json:"feed_name"`
	FeedAbrv     string `json:"feed_abrv"`
	ParentId     int64  `json:"parent_id"`
	InsertedOn   string `json:"inserted_on"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
}

type FeedSourceType struct {
	FeedTypeId int64  `json:"feed_type_id"`
	Name       string `json:"name"`
	Abrv       string `json:"abrv"`
	InsertedOn string `json:"inserted_on"`
}

type OptionalFeedSourceParams struct {
	Name         optional.String
	Tag          optional.String
	FeedSourceId optional.Int
	FeedTypeId   optional.Int
	Username     optional.String
	Password     optional.String
}
