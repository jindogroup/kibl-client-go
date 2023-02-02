package models

type Season struct {
	SeasonId   int64  `json:"season_id"`
	Name       string `json:"name"`
	Abrv       string `json:"abrv"`
	InsertedOn string `json:"inserted_on"`
}

type SeasonMapping struct {
	FeedSeasonMappingId int64  `json:"feed_season_mapping_id"`
	SeasonId            int64  `json:"season_id"`
	FeedSourceId        int64  `json:"feed_source_id"`
	FeedId              int64  `json:"feed_id"`
	FeedName            string `json:"feed_name"`
	FeedAbrv            string `json:"feed_abrv"`
	InsertedOn          string `json:"inserted_on"`
}
