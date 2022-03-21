package in_memory

import (
	"github.com/sekou-diarra/fiz-buzz-server/service/stats/suitetest"
	"testing"
)

func Test_InMemoryStore(t *testing.T) {
	store := NewInMemoryStore()
	suitetest.StatStore(t, store)
}

//type InMemoSuite struct {
//	testifysuite.StatStoreTestSuite
//}
//
//func (ims InMemoSuite)  Test_InMemorySuite() {
//	inMemStore := NewInMemoryStore()
//	ims.StatStore(&inMemStore)
//}
//
//
//func Test_Stats_Store_Implementation(t *testing.T) {
//	 suite.Run(t,new(InMemoSuite) )
//}
