package models

type CombinedLineType struct {
	CombinedLineTypeId int64 `json:"combined_line_type_id"`
	MarketTypeId1      int64 `json:"market_type_id_1"`
	MarketTypeId2      int64 `json:"market_type_id_2"`
}
