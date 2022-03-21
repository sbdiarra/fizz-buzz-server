package stats

import (
	"github.com/sekou-diarra/fiz-buzz-server/model"
)

type StatService interface {
	SaveQuery(param model.FizzbuzzParam)
	GetMostAskedQuery() *model.StatResponse
}

type StatStorer interface {
	Save(param model.FizzbuzzParam)
	GetMaxHitsQuery() *model.StatResponse
	GetAll() []string
}

type statService struct {
	store StatStorer
}

func NewStatService(store StatStorer) *statService {
	return &statService{store: store}
}

func (s *statService) SaveQuery(param model.FizzbuzzParam) {
	s.store.Save(param)
}

func (s *statService) GetMostAskedQuery() *model.StatResponse {
	return s.store.GetMaxHitsQuery()
}
