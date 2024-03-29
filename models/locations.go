package models

type Location struct {
	LocationId int64  `json:"location_id"`
	Name       string `json:"name"`
	Abrv       string `json:"abrv"`
	InsertedOn string `json:"inserted_on"`
}

type LocationMapping struct {
	FeedLocationMappingId int64  `json:"feed_location_mapping_id"`
	LocationId            int64  `json:"location_id"`
	FeedSourceId          int64  `json:"feed_source_id"`
	FeedId                int64  `json:"feed_id"`
	FeedName              string `json:"feed_name"`
	FeedAbrv              string `json:"feed_abrv"`
	InsertedOn            string `json:"inserted_on"`
}
