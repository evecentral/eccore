package eccore

import (
	"github.com/theatrus/crestmarket"
	"time"
)

func CRESTToRegion(c *crestmarket.Region) Region {
	return Region{Name: c.Name, Id: c.Id}
}

func CRESTToMarketType(c *crestmarket.MarketType) MarketType {
	return MarketType{Name: c.Name, Id: c.Id}
}

func CRESTToOrder(c *crestmarket.MarketOrder, static StaticItems) MarketOrder {
	station, _ := static.StationById(c.Station.Id)
	return MarketOrder{
		Bid:        c.Bid,
		OrderId:    c.Id,
		Station:    station,
		VolRemain:  c.Volume,
		VolEnter:   c.Volume,
		MinVolume:  c.MinVolume,
		Price:      c.Price,
		Range:      c.NumericRange(),
		Type:       MarketType{Name: c.Type.Name, Id: c.Type.Id},
		ReportedAt: time.Now(),
		Issued:     c.Issued,
		Expires:    (time.Duration(c.Duration) * time.Hour * 24),
	}
}
