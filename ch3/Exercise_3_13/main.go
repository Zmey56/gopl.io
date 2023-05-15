//“Exercise 3.13:
//Write const declarations for KB, MB, up through YB as compactly
//as you can.”

package main

const (
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
