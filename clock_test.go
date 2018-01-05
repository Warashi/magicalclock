package magicalclock_test

import (
	"testing"
	"time"

	"github.com/Warashi/magicalclock"
)

func TestApply(t *testing.T) {
	ex := []int{2, 3, 4, 5}
	h, m, s := time.Now().Clock()
	for _, e := range ex {
		a := magicalclock.New(e)
		ha, ma, sa := a.Apply(h, m, s)
		if ha == h && ma == m && sa == s {
			t.Errorf("apply %d wrong", e)
		}
	}
}
