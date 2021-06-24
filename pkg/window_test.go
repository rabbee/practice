package pkg

import (
	"fmt"
	"testing"
)

func TestWindow(t *testing.T) {
	fmt.Println(Window("abcdacbabc", "acd"))
}
