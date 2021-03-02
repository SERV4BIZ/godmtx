package main

import (
	"fmt"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/godmtx"
)

func main() {
	buffer, err := files.ReadFile("0123456789.png")
	if err != nil {
		panic(err)
	}
	val, err := godmtx.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
