package gomodtake08

import (
	"testing"
)

func TestMyTake08(t *testing.T) {
	want := "Mikki - Looki"
	if got := myFuncTake08("Mikki"); got != want {
		t.Errorf("myFuncTake08() = %q, want %q", got, want)
	}
}
