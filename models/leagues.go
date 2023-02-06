package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/jindogroup/kibl-client-go/utils"
	"github.com/markphelps/optional"
)

type League struct {
	LeagueId           int64  `json:"league_id"`
	SportId            int64  `json:"sport_id"`
	RegionId           int64  `json:"region_id"`
	Name               string `json:"name"`
	Abrv               string `json:"abrv"`
	InsertedOn         string `json:"inserted_on"`
	CombinedLineTypeId int64  `json:"combined_line_type_id,omitempty"`
}

type OptionalLeagueParams struct {
	SportsIds []int64
	LeagueIds []int64
}

func (p OptionalLeagueParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}
	if len(p.SportsIds) > 0 {
		out["sport_id"] = strings.Join(utils.String.FromInt64Slice(p.SportsIds), ",")
	}
	if len(p.LeagueIds) > 0 {
		out["league_id"] = strings.Join(utils.String.FromInt64Slice(p.LeagueIds), ",")
	}
	return
}

type LeagueMapping struct {
	FeedLeagueMappingId int64  `json:"feed_league_mapping_id"`
	LeagueId            int64  `json:"league_id,omitempty"`
	FeedSourceId        int64  `json:"feed_source_id"`
	FeedId              int64  `json:"feed_id"`
	FeedName            string `json:"feed_name"`
	FeedAbrv            string `json:"feed_abrv"`
	InsertedOn          string `json:"inserted_on"`
}

type LeagueMappingParam struct {
	LeagueId     int64 `json:"league_id,omitempty"`
	FeedSourceId int64 `json:"feed_source_id"`
}

func (p *LeagueMappingParam) String() string {
	return fmt.Sprintf("%d:%d", p.FeedSourceId, p.LeagueId)
}

type LeagueSetting struct {
	LeagueSettingID         int          `json:"league_setting_id"`
	LeagueID                int          `json:"league_id"`
	AcceptedStates          string       `json:"accepted_states"`
	AcceptedSegments        string       `json:"accepted_segments"`
	ClockBased              bool         `json:"clock_based"`
	ClockAscending          bool         `json:"clock_ascending"`
	AllowTies               bool         `json:"allow_ties"`
	HasHalftime             bool         `json:"has_halftime"`
	ScoreBased              bool         `json:"score_based"`
	ScoreAscending          bool         `json:"score_ascending"`
	SegmentBased            bool         `json:"segment_based"`
	SegmentAscending        bool         `json:"segment_ascending"`
	HalftimeStartEOFSegment optional.Int `json:"halftime_start_eof_segment"`
	HalftimeEndSOFSegment   optional.Int `json:"halftime_end_sof_segment"`
	FinalEOFSegment         optional.Int `json:"final_eof_segment"`
	InsertedOn              time.Time    `json:"inserted_on"`
}
