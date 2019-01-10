package godebug

import "testing"

func TestPrintf(t *testing.T) {
	Printf("abc%(Red)def\n")
}
