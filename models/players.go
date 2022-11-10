package models

import "github.com/antihax/optional"

type Player struct {
	PlayerId      int64  `json:"player_id"`
	ParticipantId int64  `json:"participant_id,omitempty"`
	Name          string `json:"name"`
	Abrv          string `json:"abrv"`
	Mascot        string `json:"mascot"`
	InsertedOn    string `json:"inserted_on"`
}

type OptionalPlayerParams struct {
	ParticipantId optional.Int
}

func (p OptionalPlayerParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}

	if p.ParticipantId.IsSet() {
		out["participant_id"] = p.ParticipantId.Value()
	}
	return
}

type PlayerMapping struct {
	FeedPlayerMappingId int64  `json:"feed_player_mapping_id"`
	PlayerId            int64  `json:"player_id"`
	FeedSourceId        int64  `json:"feed_source_id"`
	FeedId              int64  `json:"feed_id"`
	FeedName            string `json:"feed_name"`
	InsertedOn          string `json:"inserted_on"`
}
