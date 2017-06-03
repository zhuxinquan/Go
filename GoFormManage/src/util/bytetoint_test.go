package util

import (
	"fmt"
	"testing"
)

func TestByteToInt(t *testing.T) {
	b := []byte{0x11, 0x00, 0x03, 0xe8}
	fmt.Println(123)
	fmt.Println(ByteToInt(b))
}