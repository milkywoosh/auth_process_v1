package util

import (
	"testing"
)

func TestRandomInt(t *testing.T) {

	for i := 0; i < 10; i++ {
		RandomIntTest := RandomInt(10, 200)

		if RandomIntTest < 10 {
			t.Errorf("[FAILS] RandomIntTest must be less than 10, got %d", RandomIntTest)
		}
	}

}
