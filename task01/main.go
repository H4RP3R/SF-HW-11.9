package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/H4RP3R/queue"
)

func parseOrders(orders []string) (orderInts []uint64) {
	for _, order := range orders {
		ordInt, err := strconv.ParseUint(order, 10, 64)
		if err != nil {
			log.Fatal("invalid order type")
		}
		orderInts = append(orderInts, ordInt)
	}
	return
}

// removeOrder removes the order from the order queue by its number, returns true
// if such an order exists, otherwise returns false.
func removeOrder(number uint64, orders queue.Queue[uint64]) bool {
	return orders.Remove(number)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Enter not negative order numbers separated by spaces as arguments.\n")
	}
	flag.Parse()

	ordersInput := flag.Args()
	if len(ordersInput) == 0 {
		fmt.Println("No orders provided")
		return
	}

	orderNumbers := parseOrders(ordersInput)

	orderQueue := queue.NewQueue[uint64](orderNumbers...)
	fmt.Printf("Orders: %v\n", orderQueue)

	var orderToRemove uint64
	fmt.Println("Which order to cancel?")
	_, err := fmt.Scanln(&orderToRemove)
	if err != nil {
		log.Fatal(err)
	}

	ok := removeOrder(orderToRemove, orderQueue)
	if ok {
		fmt.Printf("Order %d canceled\n", orderToRemove)
	} else {
		fmt.Printf("Order %d not found\n", orderToRemove)
	}
	fmt.Printf("Orders: %v\n", orderQueue)
}
