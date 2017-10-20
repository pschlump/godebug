package main

import (
	"fmt"

	"github.com/pschlump/godebug"
)

func main() {
	x := godebug.ChkEnv("YepYep")
	fmt.Printf("%v\n", x)
}
