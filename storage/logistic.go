package storage

import (
	"popaket/entity"

	"gorm.io/gorm"
)

type LogisticDao interface {
	ShowAllLogistic() ([]entity.Logistics, error)
	ShowLogisticByParam(destinationName, originName string) (entity.Logistics, error)
}

func NewLogisticDao(db *gorm.DB) *dao {
	return &dao{db}
}

func (r *dao) ShowAllLogistic() ([]entity.Logistics, error) {
	var logistic []entity.Logistics

	err := r.db.Find(&logistic).Error
	if err != nil {
		return logistic, err
	}

	return logistic, nil
}

func (r *dao) ShowLogisticByParam(destinationName, originName string) (entity.Logistics, error) {
	var logistic entity.Logistics

	err := r.db.Find(&logistic).Where("destination_name = ? AND origin_name = ?").Error

	if err != nil {
		return logistic, err
	}

	return logistic, nil
}
