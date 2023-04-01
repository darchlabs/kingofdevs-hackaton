package eventstorage

import (
	"fmt"

	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage"
	"github.com/darchlabs/synchronizer-v2/pkg/event"
)

type Storage struct {
	storage *storage.S
}

func New(s *storage.S) *Storage {
	return &Storage{
		storage: s,
	}
}

func (s *Storage) GetEventsCount() (int64, error) {
	var totalRows int64
	query := "SELECT COUNT(*) FROM event"
	err := s.storage.DB.Get(&totalRows, query)
	if err != nil {
		return 0, err
	}

	return totalRows, nil
}

func (s *Storage) GetEventCountByAddress(address string) (int64, error) {
	var totalRows int64
	query := "SELECT COUNT(*) FROM event WHERE address = $1"
	err := s.storage.DB.Get(&totalRows, query, address)
	if err != nil {
		return 0, err
	}

	return totalRows, nil
}

func (s *Storage) GetEventDataCount(address string, eventName string) (int64, error) {
	var totalRows int64
	query := "SELECT COUNT(event_data.*) FROM event_data JOIN event ON event_data.event_id = event.id JOIN abi ON event.abi_id = abi.id WHERE event.address = $1 AND abi.name = $2"
	err := s.storage.DB.Get(&totalRows, query, address, eventName)
	if err != nil {
		return 0, err
	}

	return totalRows, nil
}

func (s *Storage) Stop() error {
	err := s.storage.DB.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) ListAllEvents() ([]*event.Event, error) {
	// define events response
	events := []*event.Event{}

	// get events from db
	eventQuery := "SELECT * FROM event"
	err := s.storage.DB.Select(&events, eventQuery)
	if err != nil {
		return nil, err
	}

	// iterate over events for getting abi and input values
	for _, e := range events {
		// query for getting event abi
		abi := &event.Abi{}
		abiQuery := "SELECT * FROM abi WHERE ID = $1"
		err = s.storage.DB.Get(abi, abiQuery, e.ID)
		if err != nil {
			return nil, err
		}
		e.Abi = abi

		// query for getting event abi inputs
		inputs := []*event.Input{}
		inputsQuery := "SELECT * FROM input WHERE abi_id = $1"
		err = s.storage.DB.Select(&inputs, inputsQuery, e.ID)
		if err != nil {
			return nil, err
		}
		e.Abi.Inputs = inputs
	}

	return events, nil
}

func (s *Storage) ListEvents(sort string, limit int64, offset int64) ([]*event.Event, error) {
	// define events response
	events := []*event.Event{}

	// get events from db
	eventQuery := fmt.Sprintf("SELECT * FROM event ORDER BY created_at %s LIMIT $1 OFFSET $2", sort)
	err := s.storage.DB.Select(&events, eventQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	// iterate over events for getting abi and input values
	for _, e := range events {
		// query for getting event abi
		abi := &event.Abi{}
		abiQuery := "SELECT * FROM abi WHERE ID = $1"
		err = s.storage.DB.Get(abi, abiQuery, e.ID)
		if err != nil {
			return nil, err
		}
		e.Abi = abi

		// query for getting event abi inputs
		inputs := []*event.Input{}
		inputsQuery := "SELECT * FROM input WHERE abi_id = $1"
		err = s.storage.DB.Select(&inputs, inputsQuery, e.ID)
		if err != nil {
			return nil, err
		}
		e.Abi.Inputs = inputs
	}

	return events, nil
}

func (s *Storage) ListEventsByAddress(address string, sort string, limit int64, offset int64) ([]*event.Event, error) {
	// define events response
	events := []*event.Event{}

	// get events from db
	eventQuery := fmt.Sprintf("SELECT * FROM event WHERE address = $1 ORDER BY created_at %s LIMIT $2 OFFSET $3", sort)
	err := s.storage.DB.Select(&events, eventQuery, address, limit, offset)
	if err != nil {
		return nil, err
	}

	// iterate over events for getting abi and input values
	for _, e := range events {
		// query for getting event abi
		abi := &event.Abi{}
		abiQuery := "SELECT * FROM abi WHERE ID = $1"
		err = s.storage.DB.Get(abi, abiQuery, e.ID)
		if err != nil {
			return nil, err
		}
		e.Abi = abi

		// query for getting event abi inputs
		inputs := []*event.Input{}
		inputsQuery := "SELECT * FROM input WHERE abi_id = $1"
		err = s.storage.DB.Select(&inputs, inputsQuery, e.ID)
		if err != nil {
			return nil, err
		}
		e.Abi.Inputs = inputs
	}

	return events, nil
}

func (s *Storage) ListEventTX(address string, eventName string, sort string, limit int64, offset int64) ([]*string, error) {
	// define events data response
	eventsData := []*string{}

	// define and make the query on db
	eventsDataQuery := fmt.Sprintf("SELECT event_data.Tx FROM event_data", sort)
	err := s.storage.DB.Select(&eventsData, eventsDataQuery, address, eventName, limit, offset)
	if err != nil {
		return nil, err
	}
	return eventsData, nil
}
