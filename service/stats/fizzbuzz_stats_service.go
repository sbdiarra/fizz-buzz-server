package stats

import (
	"fiz-buzz-server/model"
	"fiz-buzz-server/repository"
)

type StatService interface {
	SaveQuery(param model.FizzbuzzParam)
	GetMostAskedQuery() *model.StatResponse
}

type statService struct {
	store repository.StatStore
}

func NewStatService(store repository.StatStore) *statService {
	return &statService{store: store}
}

func (s *statService) SaveQuery(param model.FizzbuzzParam) {
	s.store.Save(param)
}

func (s *statService) GetMostAskedQuery() *model.StatResponse {
	return s.store.GetMaxHitsQuery()
}
