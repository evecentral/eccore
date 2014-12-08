package eccore

import (
	"time"
)

type Region struct {
	Name string
	Id int
}

type Station struct {
	Name string
	Id int
}

type SolarSystem struct {
	Name string
	Id int
}

type MarketType struct {
	Name string
	Id int
}

type MarketOrder struct {
	Type MarketType
	OrderId int
	Price float64
	Bid bool
	Station Station
	SolarSystem SolarSystem
	Region Region
	Range int
	VolRemain int
	VolEnter int
	MinVolume int
	Expires time.Duration
	ReportedAt time.Time
}
