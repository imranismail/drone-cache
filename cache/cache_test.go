package cache

import (
	"testing"
)

func TestCacheTruth(t *testing.T) {
	if true != true {
		t.Error("everything I know is wrong")
	}
}