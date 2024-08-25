package storage

import (
	"brasa/db"
	inter "brasa/internal/tokens/interface"
	"time"

	"gorm.io/gorm"
)

func CreateTokenDb(q *gorm.DB, input inter.SaveToken) (*db.Token, error) {
	newToken := db.Token{
		UriNumber:    input.UriNumber,
		ContractID:   input.ContractID,
		AssetID:      input.AssetID,      // Adicionando AssetID
		ExpireDate:   input.ExpireDate,   // Adicionando ExpireDate
		NominalValue: input.NominalValue, // Adicionando NominalValue
		Guarantees:   input.Guarantees,   // Adicionando Guarantees
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := q.Create(&newToken).Error; err != nil {
		return nil, err
	}

	return &newToken, nil
}

func GetAllTokensDb(repo *gorm.DB, contractId string) ([]db.Token, error) {
	var tokens []db.Token
	err := repo.Where("contract_id = ?", contractId).Find(&tokens).Error
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
