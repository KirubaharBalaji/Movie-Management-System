package main

import "fmt"

func getMovieName(movieChoice int) string {
	switch movieChoice {
	case 1:
		movieName = "Vikram"
	case 2:
		movieName = "Beast"
	case 3:
		movieName = "Mamanidhan"
	case 4:
		movieName = "K.G.F Chapter 2"
	case 5:
		movieName = "Valimai"
	case 6:
		movieName = "Don"
	case 7:
		movieName = "Mahaan"
	default:
		fmt.Println("Invalid Choice \t Enter Only From The Given Choice")
	}
	return movieName

}

func confirmMovie(confirmmovie int) string {
	switch confirmmovie {
	case 1:
		movieName = "Vikram"
	case 2:
		movieName = "Beast"
	case 3:
		movieName = "Mamanidhan"
	case 4:
		movieName = "K.G.F Chapter 2"
	case 5:
		movieName = "Valimai"
	case 6:
		movieName = "Don"
	case 7:
		movieName = "Mahaan"
	default:
		fmt.Println("Invalid Choice \t Enter Only From The Given Choice")
	}
	return movieName

}
