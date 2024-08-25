package storage

import (
	"brasa/db"
)

func CreateReceivable(receivable *db.Receivable) error {
	return db.Repo.Create(receivable).Error
}

func GetReceivableByID(id string) (*db.Receivable, error) {
	var receivable db.Receivable
	err := db.Repo.First(&receivable, id).Error
	if err != nil {
		return nil, err
	}
	return &receivable, nil
}

func GetAllReceivables() ([]db.Receivable, error) {
	var receivables []db.Receivable
	err := db.Repo.Find(&receivables).Error
	if err != nil {
		return nil, err
	}
	return receivables, nil
}

func GetContractByID(contractID string) (*db.Contract, error) {
	var contract db.Contract
	if err := db.Repo.Where("contract_id = ?", contractID).First(&contract).Error; err != nil {
		return nil, err
	}
	return &contract, nil
}

func UpdateContract(contract *db.Contract) error {
	return db.Repo.Save(contract).Error
}

func RegisterDistribution(investorID string, amount float64, contractID string) error {
	distribution := db.Distribution{
		InvestorID: investorID,
		Amount:     amount,
		ContractID: contractID,
	}
	return db.Repo.Create(&distribution).Error
}

func GetAllTokens(contractId string) ([]db.Token, error) {
	var tokens []db.Token
	err := db.Repo.Where("contract_id = ?", contractId).Find(&tokens).Error
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
