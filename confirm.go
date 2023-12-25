package main

import (
	"fmt"
)

func afterConfirmation() {
	fmt.Printf("Changes After confirmation:\n")
	fmt.Printf("1.Theatre Name:%v\n2.Location:%v\n3.Movie:%v\n", cinemaName, locationTheatre, movieName)
}
func confirmator() {
	var choice int
	fmt.Printf("What do you want to make changes with :\n1.Theatre:%v\n2.Location:%v\n3.Movie:%v\n", cinemaName, locationTheatre, movieName)
	fmt.Scan(&choice)
	switch choice {
	case 1:
		var confirmtheatre int
		fmt.Printf("Please Enter the Correct choice:\n")
		fmt.Printf("The Following Theatres are available :\n 1.Arnold Cinemas\n 2.Rocky Cinemas\n 3.LUXE Cinemas\n 4.AGS Cinemas\n 5.INOX Cinemas\n 6.PVR Cinemas\n 7.SPI Cinemas\n")
		fmt.Println("Enter Your Choice...:")
		fmt.Scan(&confirmtheatre)
		confirmTheatre(confirmtheatre)
	case 2:
		var confirmlocation int
		fmt.Printf("PLease Enter the Correct Choice\n")
		fmt.Printf("The Following Theatres are available in the Following Locations :\n 1.Aynavaram\n 2.Kolathur\n 3.Perambur\n 4.Mount Road\n 5.Kilpauk\n 6.T.Nagar\n 7.Numngambakkam\n")
		fmt.Println("Enter Your Choice...:")
		fmt.Scan(&confirmlocation)
		confirmLocation(confirmlocation)

	case 3:
		var confirmmovie int
		fmt.Printf("PLease Enter the Correct Choice\n")
		fmt.Printf("The Following Movies are available :\n 1.Vikram\n 2.Beast\n 3.Mamanidhan\n 4.K.G.F Chapter 2\n 5.Valimai\n 6.Don\n 7.Mahaan\n")
		fmt.Println("Enter Your Choice...:")
		fmt.Scan(&confirmmovie)
		confirmMovie(confirmmovie)
	default:
		fmt.Printf("Please Enter Valid Choices !\n")
	}
}
