package main

import (
	"log"
)

const (
	x = 1 << 8
)
func myfunc() {
	log.Println(x)
}


func main() {
	myfunc()
}
