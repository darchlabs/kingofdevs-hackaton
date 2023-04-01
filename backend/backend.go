package backend

import (
	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/smartcontract"
)

type SmartContractStorage interface {
	ListSmartContracts(sort string, limit int64, offset int64) ([]*smartcontract.SmartContract, error)
	InsertSmartContract(s *smartcontract.SmartContract) (*smartcontract.SmartContract, error)
	GetSmartContractsCount() (int64, error)
	Stop() error
}
