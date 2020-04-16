package gomodtake08

import (
	"testing"
)

//v.1.2.0
func TestMyTake08(t *testing.T) {
	want := "Mikki - Looki"
	if got := MyFuncTake08("Mikki"); got != want {
		t.Errorf("MyFuncTake08() = %q, want %q", got, want)
	}
}

