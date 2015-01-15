package eccore

import (
	"database/sql"
	"fmt"
	"log"
)

type dbOrderStore struct {
	db     *sql.DB
	static StaticItems
}

type DBOrderStore interface {
	UpdateOrders(region Region, mt MarketType, orders []MarketOrder) error
}

func NewOrderStore(db *sql.DB, static StaticItems) (DBOrderStore, error) {
	log.Println("Building new order store")
	return &dbOrderStore{db: db, static: static}, nil
}

func (d *dbOrderStore) UpdateOrders(region Region, mt MarketType, orders []MarketOrder) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM current_market WHERE regionid = $1 AND typeid = $2", region.Id, mt.Id)
	if err != nil {
		tx.Rollback()
		log.Printf("Orders can't delete due to %s", err)
		return err
	}

	insCurrent, err := tx.Prepare(`
INSERT INTO current_market
(regionid, systemid, stationid, typeid, bid, price, orderid,
minvolume, volremain, volenter, issued, duration, range, reportedby, reportedtime)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, CAST ($12 AS INTERVAL), $13, 0, $14)`)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	insArchive, err := tx.Prepare(`
INSERT INTO archive_market
(regionid, systemid, stationid, typeid, bid, price, orderid,
minvolume, volremain, volenter, issued, duration, range, reportedby, source)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, CAST ($12 AS INTERVAL), $13, 0, 'evec_upload_cqache')`)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	for _, order := range orders {
		duration := fmt.Sprintf("%d", int(order.Expires.Hours()/24))
		// Historically, this is an integer column. welp.
		bid := 0
		if order.Bid {
			bid = 1
		}
		_, err = insCurrent.Exec(region.Id, order.Station.SolarSystem.Id, order.Station.Id,
			mt.Id, bid, order.Price, order.OrderId, order.MinVolume, order.VolRemain, order.VolEnter,
			order.Issued, duration, order.Range, order.ReportedAt)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return err
		}
		_, err = insArchive.Exec(region.Id, order.Station.SolarSystem.Id, order.Station.Id,
			mt.Id, bid, order.Price, order.OrderId, order.MinVolume, order.VolRemain, order.VolEnter,
			order.Issued, duration, order.Range)

		if err != nil {
			tx.Rollback()
			log.Println(err)
			return err
		}
	}

	insArchive.Close()
	insCurrent.Close()
	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	return err
}
