package mem

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []string{
		"one",
		"two",
		"three",
		"four",
	}

	for _, test := range tests {
		fmt.Printf("%x = %s\n", Hash(test, len(test)), test)
	}
}
