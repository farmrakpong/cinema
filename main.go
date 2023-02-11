package main

import (
	"fmt"
	"github.com/farmrakpong/cinema/movie"
	ticket "github.com/farmrakpong/cinema/ticket/packageticket"
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
