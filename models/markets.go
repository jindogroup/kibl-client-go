package models

import (
	"strings"

	"github.com/antihax/optional"
)

type Market struct {
	MarketKey                   string               `json:"market_key"`
	Participants                []MarketParticipant  `json:"participants"`
	MarketReferenceExternalIds  []MarketExternalRef  `json:"market_reference_external_ids,omitempty"`
	FixtureReferenceExternalIds []FixtureExternalRef `json:"fixture_reference_external_ids,omitempty"`
	SideReferenceExternalIds    []SideExternalRef    `json:"side_reference_external_ids,omitempty"`
}

type OptionalMarketInfoParams struct {
	FixtureId            optional.Int64
	FixtureParticipantId optional.Int64
	MarketTypeId         optional.Int64
	SideId               optional.Int64
	SegmentId            optional.Int64
	FeedSourceId         optional.Int64
	MarketId             optional.Int64
	SinceLastUpdated     optional.Time
	IncludeExternalIds   []string
	ExternalLeagueId     *LeagueMappingParam
}

func (p OptionalMarketInfoParams) Params() (out map[string]interface{}) {
	out = map[string]interface{}{}

	if p.FixtureId.IsSet() {
		out["fixture_id"] = p.FixtureId.Value()
	}
	if p.FixtureParticipantId.IsSet() {
		out["fixture_participant_id"] = p.FixtureParticipantId.Value()
	}
	if p.MarketTypeId.IsSet() {
		out["market_type_id"] = p.MarketTypeId.Value()
	}
	if p.SideId.IsSet() {
		out["side_id"] = p.SideId.Value()
	}
	if p.SegmentId.IsSet() {
		out["segment_id"] = p.SegmentId.Value()
	}
	if p.FeedSourceId.IsSet() {
		out["feed_source_id"] = p.FeedSourceId.Value()
	}
	if p.MarketId.IsSet() {
		out["market_id"] = p.MarketId.Value()
	}
	if p.SinceLastUpdated.IsSet() {
		out["since_last_updated"] = p.SinceLastUpdated.Value()
	}
	if len(p.IncludeExternalIds) > 0 {
		out["include_external_ids"] = strings.Join(p.IncludeExternalIds, ",")
	}
	if p.ExternalLeagueId != nil {
		out["external_league_id"] = p.ExternalLeagueId.String()
	}
	return
}

type MarketExternalRef struct {
	MarketTypeId int64  `json:"market_type_id"`
	SegmentId    int64  `json:"segment_id"`
	FeedSourceId int64  `json:"feed_source_id"`
	FeedId       int64  `json:"feed_id"`
	FeedName     string `json:"feed_name"`
	FeedAbrv     string `json:"feed_abrv"`
}

type MarketParticipant struct {
	ParticipantId           int64                    `json:"participant_id"`
	ParticipantSideId       int64                    `json:"participant_side_id"`
	MarketId                int64                    `json:"market_id"`
	FeedSourceId            int64                    `json:"feed_source_id"`
	FixtureId               int64                    `json:"fixture_id"`
	FixtureParticipantId    int64                    `json:"fixture_participant_id"`
	MarketTypeId            int64                    `json:"market_type_id"`
	SegmentId               int64                    `json:"segment_id"`
	Point                   float64                  `json:"point"`
	PriceAmerican           int64                    `json:"price_american"`
	PriceDecimal            float64                  `json:"price_decimal"`
	PriceFraction           string                   `json:"price_fraction"`
	IsLive                  bool                     `json:"is_live"`
	IsOpener                bool                     `json:"is_opener"`
	IsPrevious              bool                     `json:"is_previous"`
	IsCurrent               bool                     `json:"is_current"`
	InsertedOn              string                   `json:"inserted_on"`
	IsMain                  bool                     `json:"is_main"`
	MarketStatusId          int64                    `json:"market_status_id"`
	SideId                  int64                    `json:"side_id"`
	BetKey                  string                   `json:"bet_key"`
	ParticipantExternalRefs []ParticipantExternalRef `json:"participant_reference_external_ids,omitempty"`
}

type MarketStatus struct {
	MarketStatusId int64  `json:"market_status_id"`
	Name           string `json:"name"`
	Abrv           string `json:"abrv"`
	InsertedOn     string `json:"inserted_on"`
}

type MarketStatusMapping struct {
	FeedMarketStatusMappingId int64  `json:"feed_market_status_mapping_id"`
	MarketStatusId            int64  `json:"market_status_id"`
	FeedSourceId              int64  `json:"feed_source_id"`
	FeedId                    int64  `json:"feed_id"`
	FeedName                  string `json:"feed_name"`
	FeedAbrv                  string `json:"feed_abrv"`
	InsertedOn                string `json:"inserted_on"`
}

type MarketType struct {
	MarketTypeId int64  `json:"market_type_id"`
	Name         string `json:"name"`
	Abrv         string `json:"abrv"`
}

type MarketTypeMapping struct {
	FeedMarketTypeMappingId int64  `json:"feed_market_type_mapping_id"`
	MarketTypeId            int64  `json:"market_type_id"`
	FeedSourceId            int64  `json:"feed_source_id"`
	FeedId                  int64  `json:"feed_id"`
	FeedName                string `json:"feed_name"`
	FeedAbrv                string `json:"feed_abrv"`
	InsertedOn              string `json:"inserted_on"`
	SegmentId               int64  `json:"segment_id"`
	SideId                  int64  `json:"side_id"`
}
