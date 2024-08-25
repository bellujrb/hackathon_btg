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
