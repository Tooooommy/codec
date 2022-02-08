package codec

import (
	"fmt"
	"testing"
)

func TestNop(t *testing.T) {
	nop := new(Nop)
	out, err := nop.Encode([]byte("hello"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(out))
}
