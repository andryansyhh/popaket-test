package service

import (
	"errors"
	"fmt"
	"popaket/entity"
	"popaket/storage"
)

type LogisticService interface {
	ShowAllLogistic() ([]entity.Logistics, error)
	ShowLogisticByParam(destinationName, OriginName string) (entity.Logistics, error)
}

type logisticservice struct {
	dao storage.LogisticDao
}

func NewLogisticService(dao storage.LogisticDao) *logisticservice {
	return &logisticservice{dao}
}

func (s *logisticservice) ShowAllLogistic() ([]entity.Logistics, error) {
	logistic, err := s.dao.ShowAllLogistic()
	if err != nil {
		return logistic, err
	}

	return logistic, nil
}

func (s *logisticservice) ShowLogisticByParam(destinationName, originName string) (entity.Logistics, error) {
	logistic, err := s.dao.ShowLogisticByParam(destinationName, originName)
	if err != nil {
		return entity.Logistics{}, err
	}

	if logistic.LogisticName == "" {
		newError := fmt.Sprintf("logistic name not found")
		return entity.Logistics{}, errors.New(newError)
	}

	if logistic.DestinationName == "" {
		newError := fmt.Sprintf("destination name not found")
		return entity.Logistics{}, errors.New(newError)
	}

	return logistic, nil
}
