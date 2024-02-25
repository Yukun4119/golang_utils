package set

import (
	"github.com/Yukun4119/golang_utils/log"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	st := New()
	log.Info("%+v", st.IsExist(1))
	log.Info("%+v", st)
	st.Add(1)
	log.Info("%+v", st.IsExist(1))
	log.Info("%+v", st)
}

func TestHashSet_Empty(t *testing.T) {

}

func TestHashSet_IsExist(t *testing.T) {

}

func TestHashSet_Remove(t *testing.T) {

}

func TestHashSet_Size(t *testing.T) {

}
