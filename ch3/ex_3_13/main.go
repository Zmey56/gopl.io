// “Exercise 3.13:
// Write const declarations for KB, MB, up through YB as compactly
// as you can.”
package main

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

const (
	DELTA = 1024
	KiB
	MiB = KiB * DELTA
	GiB = MiB * DELTA
	TiB = GiB * DELTA
	PiB = TiB * DELTA
	EiB = PiB * DELTA
	ZiB = EiB * DELTA
	YiB = ZiB * DELTA
)

func main() {
	fmt.Fprintf(stdout, "ZiB/EiB = %v\n", ZiB/EiB)
}
