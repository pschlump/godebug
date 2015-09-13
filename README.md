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

LF() takes an optional parameter, 1 indicates the current
function, 2 is the parent, 3 the grandparent.

I will add complete documentation tomorrow.

