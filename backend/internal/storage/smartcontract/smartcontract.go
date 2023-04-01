package smartcontractstorage

import (
	"errors"

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
	return nil, errors.New("method InsertSmartContract not implemented")
}

func (s *Storage) ListSmartContracts(sort string, limit int64, offset int64) ([]*smartcontract.SmartContract, error) {
	return nil, errors.New("method ListEvents not implemented")
}

func (s *Storage) GetSmartContractsCount() (int64, error) {
	return 0, errors.New("method GetSmartContractsCount not implemented")
}

func (s *Storage) Stop() error {
	err := s.storage.DB.Close()
	if err != nil {
		return err
	}

	return nil
}
