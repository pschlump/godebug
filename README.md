# godebug

Small library of tools for debugging Go (golang) programs.

The most useful of these is LF().  It returns a line number
and file name.  The parameter is an optional number.
1 indicates that you want it for the current call. 2 would
be the parent of the current call.

Example:

``` golang

	package main

	import (
		"fmt"

		"github.com/pschlump/godebug"
	)

	func main() {
		fmt.Printf("I am at: %s\n", debug.LF())
	}

```

I will add complete documentation tomorrow.

