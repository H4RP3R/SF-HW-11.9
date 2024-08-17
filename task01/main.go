package main

import (
	"github.com/H4RP3R/queue"
)

func main() {
	q := queue.NewQueue[int]()
	q.Poll()
}
