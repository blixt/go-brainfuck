package brainfuck

import (
	"bytes"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	const (
		in = "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
		out = "Hello World!\n"
	)

	buffer := new(bytes.Buffer)
	if _, err := Run(in, "", buffer); err != nil {
		t.Fatalf("Run failed: %s", err)
	} else {
		if buffer.String() != out {
			t.Errorf("%v != %v", buffer.String(), out)
		}
	}
}
