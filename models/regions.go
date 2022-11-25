package models

type Region struct {
	RegionId   int64  `json:"region_id"`
	Name       string `json:"name"`
	Abrv       string `json:"abrv"`
	InsertedOn string `json:"inserted_on"`
}
