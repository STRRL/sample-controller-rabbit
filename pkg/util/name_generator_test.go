package util

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for i := 0; i < 50; i++ {
		fmt.Println(GetName())
	}
}
