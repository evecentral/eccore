package eccore

import (
	"database/sql"
	"log"
)

// StaticItems provides utilities to load
// common item maps, such as Item IDs to names,
// solar systems to maps, and more.
type StaticItems interface {
	StationById(id int) (Station, bool)
	SystemById(id int) (SolarSystem, bool)
	RegionById(id int) (Region, bool)
}

// mapProvider provides a static lookup
// map which fulfills the StaticItems interface
type mapProvider struct {
	stationsById map[int]Station
	systemsById  map[int]SolarSystem
	regionsById  map[int]Region
}

func (m *mapProvider) StationById(id int) (Station, bool) {
	st, ok := m.stationsById[id]
	return st, ok
}

func (m *mapProvider) SystemById(id int) (SolarSystem, bool) {
	st, ok := m.systemsById[id]
	return st, ok
}

func (m *mapProvider) RegionById(id int) (Region, bool) {
	st, ok := m.regionsById[id]
	return st, ok
}

// NewStaticItemsFromDatabase returns a StaticItems
// which has been sourced from tables in the given
// database object.
func NewStaticItemsFromDatabase(db *sql.DB) (StaticItems, error) {

	provider := &mapProvider{stationsById: make(map[int]Station),
		systemsById: make(map[int]SolarSystem),
		regionsById: make(map[int]Region)}
	res, err := db.Query("SELECT regionid, regionname FROM regions")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for res.Next() {
		region := Region{}
		res.Scan(&region.Id, &region.Name)
		provider.regionsById[region.Id] = region
	}

	res, err = db.Query("SELECT systemid, systemname, regionid, security, truesec FROM systems")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for res.Next() {
		ss := SolarSystem{}
		var regionId int
		res.Scan(&ss.Id, &ss.Name, &regionId, &ss.Security, &ss.TrueSecurity)
		ss.Region = provider.regionsById[regionId]
		provider.systemsById[ss.Id] = ss
	}

	res, err = db.Query("SELECT stationid, stationname, systemid FROM stations")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for res.Next() {
		st := Station{}
		var systemId int
		res.Scan(&st.Id, &st.Name, &systemId)
		st.SolarSystem = provider.systemsById[systemId]
		provider.stationsById[st.Id] = st
	}
	return provider, nil
}
