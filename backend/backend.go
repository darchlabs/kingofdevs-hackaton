package backend

import (
	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/smartcontract"
	"github.com/darchlabs/kingofdevs-hackaton/backend/pkg/transaction"
)

type SmartContractStorage interface {
	ListSmartContracts(sort string, limit int64, offset int64) ([]*smartcontract.SmartContract, error)
	InsertSmartContract(s *smartcontract.SmartContract) (*smartcontract.SmartContract, error)
	GetSmartContractByID(id string) (*smartcontract.SmartContract, error)
	GetSmartContractsCount() (int64, error)
	Stop() error
}

type TransactionStorage interface {
	ListTxs() ([]*transaction.Transaction, error)
	GetTxById(id string) (*transaction.Transaction, error)
	InsertTx(t *transaction.Transaction) (*transaction.Transaction, error)
	UpdateTx(t *transaction.Transaction) (*transaction.Transaction, error)
}
