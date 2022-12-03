package models

import (
	"strings"

	"github.com/jindogroup/kibl-client-go/utils"
)

type Player struct {
	PlayerId      int64  `json:"player_id"`
	ParticipantId int64  `json:"participant_id,omitempty"`
	Name          string `json:"name"`
	Abrv          string `json:"abrv"`
	Mascot        string `json:"mascot"`
	InsertedOn    string `json:"inserted_on"`
}

type OptionalPlayerParams struct {
	ParticipantIds []int64
}

func (p OptionalPlayerParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if len(p.ParticipantIds) > 0 {
		out["participant_id"] = strings.Join(utils.String.FromInt64Slice(p.ParticipantIds), ",")
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
