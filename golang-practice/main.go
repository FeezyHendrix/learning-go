package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(5) + 1
		stars := strings.Repeat("*", r)
		fmt.Println(stars)
	}

}
