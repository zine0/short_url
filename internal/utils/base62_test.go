package utils

import (
	"fmt"
	"testing"
)

func TestBase62(t *testing.T) {
	var i int64 = 56800235583
	base := IntToBase62(i)
	fmt.Printf("num = %d\nbase62 = %s\n", i, base)
	if i != Base62ToInt(base) {
		t.Fail()
	}
}
