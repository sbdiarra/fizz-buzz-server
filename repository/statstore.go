package repository

import "fiz-buzz-server/model"

type StatStore interface {
	Save(param model.FizzbuzzParam)
	GetMaxHitsQuery() *model.StatResponse
	GetAll() []string
}
