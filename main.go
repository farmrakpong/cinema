package main

import (
	"fmt"
	"learn/farm/cinema/movie"
	ticket "learn/farm/cinema/ticket/packageticket"
)

func init() {
	fmt.Println("vary good")
}
func main() {
	moviesName := movie.FindName("tt4154796999")
	fmt.Println(moviesName)
	movie.Review("avg", 8)
	ticket.BuyTicket(moviesName)
}
