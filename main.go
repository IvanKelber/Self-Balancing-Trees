package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ivankelber/sbt/redblack"
)

func main() {
	rb := redblack.RedBlackTree(50)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		for !rb.Insert(r.Intn(100)) {
			fmt.Println("Trying again...")
		}
	}
	fmt.Println(rb)
}
