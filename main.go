package main

import (
	"os"
	"pgo/internal"
)

func main() {
	lin := os.Args[1]
	internal.GetVideoLink(lin)
}
