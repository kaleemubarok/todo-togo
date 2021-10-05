package service

import (
	"log"
	"todo-togo/entity"
	"todo-togo/repository"
)

type IStatusService interface {
	PrepareAllStatus()
}

var MapStatus = make(map[int]string)

type StatusService struct {
	stat repository.IStatusRepo
}

func NewStatusService(repo *repository.IStatusRepo) IStatusService  {
	return &StatusService{stat: *repo}
}

func (s *StatusService) PrepareAllStatus() {
	var allStat []entity.Status
	allStat, err := s.stat.SelectAllStatus()
	if err != nil {
		log.Println("error on prepare all status map,", err)
	}

	for _, s := range allStat {
		MapStatus[s.StatusID] = s.StatusTxt
	}
}
