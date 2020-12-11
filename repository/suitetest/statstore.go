package suitetest

import (
	"fiz-buzz-server/model"
	"fiz-buzz-server/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

// suitetest.StatStore
func StatStore(t *testing.T, store repository.StatStore) {
	var stubQuery = model.FizzbuzzParam{FstMultiple: 3, SecMultiple: 9, Limit: 5, Label1: "label1", Label2: "label2"}
	var stubQuery2 = model.FizzbuzzParam{FstMultiple: 4, SecMultiple: 3, Limit: 5, Label1: "label1", Label2: "label2"}

	t.Run("result_should_be_empty", func(t *testing.T) {
		emptyResult := store.GetAll()
		assert.Len(t, emptyResult, 0)
	})
	t.Run("should_return_nil", func(t *testing.T) {
		emptyResult := store.GetMaxHitsQuery()
		assert.Nil(t, emptyResult)
	})

	t.Run("should_correctly_save_the_request", func(t *testing.T) {
		store.Save(stubQuery)
		assert.Len(t, store.GetAll(), 1)
	})

	t.Run("should_return_the_request_with_the_most_hit", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			store.Save(stubQuery2)
		}
		got := store.GetMaxHitsQuery()
		want := &model.StatResponse{
			Hits:            5,
			AssociateParams: stubQuery2.String(),
		}
		assert.Equal(t, want, got)
	})
}
