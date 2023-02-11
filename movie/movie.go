package movie

import "fmt"

func Review(name string, rating float64) {
	fmt.Printf("!!!!%s: %f\n", name, rating)
}

func FindName(imdb string) string {
	switch imdb {
	case "tt4154796":
		return "bad boy121"
	case "tt191":
		return "asd"
	}
	return "badd game"
}
