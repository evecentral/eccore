package eccore

import (
	"time"
)

type Region struct {
	Name string
	Id   int
}

type SolarSystem struct {
	Name         string
	Id           int
	Security     float64
	TrueSecurity float64
	Region       Region
}

type Station struct {
	Name        string
	Id          int
	SolarSystem SolarSystem
}

type MarketType struct {
	Name string
	Id   int
}

type MarketOrder struct {
	Type       MarketType
	OrderId    int
	Price      float64
	Bid        bool
	Station    Station
	Range      int
	VolRemain  int
	VolEnter   int
	MinVolume  int
	Issued     time.Time
	Expires    time.Duration
	ReportedAt time.Time
}
