package models

import "github.com/antihax/optional"

type Sportsbook struct {
	FeedSourceId int64  `json:"feed_source_id"`
	Name         string `json:"name"`
	Tag          string `json:"tag"`
	FeedTypeId   int64  `json:"feed_type_id"`
	InsertedOn   string `json:"inserted_on"`
}

type OptionalSportsbookParams struct {
	Name         optional.String
	Tag          optional.String
	FeedSourceId optional.Int64
	FeedTypeId   optional.Int64
}
