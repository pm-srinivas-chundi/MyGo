package main

import (
	"log"
)

const MAX = 10
func main() {
	flag := true

	if flag == true {
		log.Println("YES")
	}
	
	i:=1
	for  i < MAX {
		log.Printf("Counter: %v\n", i)
		i += 1
	}
}

