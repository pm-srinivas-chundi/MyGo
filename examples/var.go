package main

import (
	"log"
)

func myfunc() {
	var a,b bool = true, false
	var x,y int = 1 , 2
	var str string = "WELCOME"

	i, s, f, b    :=    1, "PUBMATIC", 3.141, true

	log.Println(a, b)
	log.Println(x, y)
	log.Println(str)
	log.Println(i, s, f, b)
}


func main() {
	myfunc()
}
