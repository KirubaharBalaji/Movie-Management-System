package main

import "fmt"

func getTheatre(theatre int) string {
	switch theatre {
	case 1:
		cinemaName = "Arnold Cinemas"
	case 2:
		cinemaName = "Rocky Cinemas"
	case 3:
		cinemaName = "Luxe Cinemas"
	case 4:
		cinemaName = "AGS Cinemas"
	case 5:
		cinemaName = "INOX Cinemas"
	case 6:
		cinemaName = "PVR Cinemas"
	case 7:
		cinemaName = "SPI Cinemas"
	default:
		fmt.Println("Invalid Choice \t Enter Only From The Given Choice")
	}
	return cinemaName
}

func getLocation(locationTheatreChoice int) string {
	switch locationTheatreChoice {
	case 1:
		locationTheatre = "Aynavaram"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 2:
		locationTheatre = "Kolathur"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 3:
		locationTheatre = "Perambur"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 4:
		locationTheatre = "Mount Road"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 5:
		locationTheatre = "Kilpauk"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 6:
		locationTheatre = "T.Nagar"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 7:
		locationTheatre = "Nungambakkam"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	default:
		fmt.Println("Invalid Choice \t Enter Only From The Given Choice")
	}
	return locationTheatre
}
func confirmTheatre(confirmtheatre int) string {
	switch confirmtheatre {
	case 1:
		cinemaName = "Arnold Cinemas"
	case 2:
		cinemaName = "Rocky Cinemas"
	case 3:
		cinemaName = "Luxe Cinemas"
	case 4:
		cinemaName = "AGS Cinemas"
	case 5:
		cinemaName = "INOX Cinemas"
	case 6:
		cinemaName = "PVR Cinemas"
	case 7:
		cinemaName = "SPI Cinemas"
	default:
		fmt.Println("Invalid Choice \t Enter Only From The Given Choice")
	}
	return cinemaName
}

func confirmLocation(confirmlocation int) string {
	switch confirmlocation {
	case 1:
		locationTheatre = "Aynavaram"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 2:
		locationTheatre = "Kolathur"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 3:
		locationTheatre = "Perambur"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 4:
		locationTheatre = "Mount Road"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 5:
		locationTheatre = "Kilpauk"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 6:
		locationTheatre = "T.Nagar"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	case 7:
		locationTheatre = "Nungambakkam"
	//	fmt.Printf("%v %v .\n", cinemaName, locationTheatre)
	default:
		fmt.Println("Invalid Choice \t Enter Only From The Given Choice")
	}
	return locationTheatre
}
