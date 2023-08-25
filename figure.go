package main

import 	"fmt"

func printFigure(quantity int){ //The parameter of how big the diamond is going to be
	var i, j, lines, blanks int //the variables needed

	lines = 1; blanks = quantity

	for i = 0; i < quantity * 2 + 1 ; i++{ //The quantity formula is there to lower the amount of blanks and lines
		for j = 0; j < blanks; j++{ //Prints the blank spaces
			fmt.Print(" ")
		}

		for j = 0; j < lines; j++{ //Prints the visible part of the shape
			fmt.Print("-")
		}

		fmt.Println()

		if i < quantity{  //The i can now be bigger than the quatity based on the formula instead of breaking the cycle
			blanks--; lines +=2
		} else {
			blanks++; lines -=2
		}
	}
}

func main(){
	fmt.Println("-------Figure-------")
	printFigure(3)
	printFigure(8)
}
