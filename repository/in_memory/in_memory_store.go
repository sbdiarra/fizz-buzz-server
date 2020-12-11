package in_memory

import (
	"fiz-buzz-server/model"
	"fiz-buzz-server/repository"
	"sort"
)

type internalModel struct {
	query string
	hits  int
}

type inMemoryStore struct {
	storage map[string]internalModel
}

func (i *inMemoryStore) Save(param model.FizzbuzzParam) {
	im, ok := i.storage[param.String()]
	if ok {
		im.hits++
		i.storage[param.String()] = im
		return
	}

	i.storage[param.String()] = internalModel{query: param.String(), hits: 1}

}

func (i *inMemoryStore) GetMaxHitsQuery() *model.StatResponse {
	var internals []internalModel
	if len(i.storage) == 0 {
		return nil
	}
	for _, models := range i.storage {
		internals = append(internals, models)
	}
	sort.SliceStable(internals, func(i, j int) bool {
		return internals[i].hits > internals[j].hits
	})

	return &model.StatResponse{
		Hits:            internals[0].hits,
		AssociateParams: internals[0].query,
	}
}

func (i *inMemoryStore) GetAll() []string {
	var elements []string
	for key := range i.storage {
		elements = append(elements, key)
	}
	return elements
}

func NewInMemoryStore() repository.StatStore {
	return &inMemoryStore{storage: make(map[string]internalModel)}
}
