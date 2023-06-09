package smartcontractstorage

import (
	"fmt"

	"github.com/darchlabs/kingofdevs-hackaton/backend/internal/storage"
	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/smartcontract"
)

type Storage struct {
	storage *storage.S
}

func New(s *storage.S) *Storage {
	return &Storage{
		storage: s,
	}
}

func (s *Storage) InsertSmartContract(sc *smartcontract.SmartContract) (*smartcontract.SmartContract, error) {
	// insert new smartcontract in database
	var smartcontractId string
	query := "INSERT INTO smartcontract (id, name, network, node_url, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := s.storage.DB.Get(&smartcontractId, query, sc.ID, sc.Name, sc.Network, sc.NodeURL, sc.Address, sc.CreatedAt, sc.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// get created smartcontract
	createdSmartcontract, err := s.GetSmartContractByID(smartcontractId)
	if err != nil {
		return nil, err
	}

	return createdSmartcontract, nil
}

func (s *Storage) GetSmartContractByID(id string) (*smartcontract.SmartContract, error) {
	// get smartcontract from db
	sc := &smartcontract.SmartContract{}
	err := s.storage.DB.Get(sc, "SELECT * FROM smartcontract WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return sc, nil
}

func (s *Storage) ListSmartContracts(sort string, limit int64, offset int64) ([]*smartcontract.SmartContract, error) {
	// define smartcontracts response
	smartcontracts := []*smartcontract.SmartContract{}

	// get smartcontracts from db
	scQuery := fmt.Sprintf("SELECT * FROM smartcontract ORDER BY created_at %s LIMIT $1 OFFSET $2", sort)
	err := s.storage.DB.Select(&smartcontracts, scQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	return smartcontracts, nil
}

func (s *Storage) GetSmartContractsCount() (int64, error) {
	var totalRows int64
	query := "SELECT COUNT(*) FROM smartcontract"
	err := s.storage.DB.Get(&totalRows, query)
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
