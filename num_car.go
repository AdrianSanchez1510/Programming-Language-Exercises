package main

import (
	"fmt"
	"strings"
)
//We are going to use the use the Split() function to convert the string in an array

func Count(phrase string){
	fmt.Println("\nPhrase: " + phrase + "\n") //Prints the phrase
	
	var countC, countW, countI, i int 

	cha := strings.Split(phrase, "") //This is the character array
	wor := strings.Split(phrase, " ") //This is the word array
	ide := strings.Split(phrase, "\n") //This is the line array


	countC = 0
	countW = 0
	countI = 0

	for i = 0; i < len(cha); i++{ //This for counts every single character (including spaces)
		countC = countC + 1
	}

	for i = 0; i < len(wor); i++{ //This one counts every word
		countW = countW + 1
	}

	for i = 0; i < len(ide); i++{ //And this one counts every line
		countI = countI + 1
	}


	fmt.Println("The amount of characters for this phrase is: ", countC)
	fmt.Println("The amount of words for this phrase is: ", countW)
	fmt.Println("The amount of lines for this phrase is: ", countI)
}

func main(){
	var newString string = "Sample Text \nDon't Use it Incorrectly \nBy Adrián Sánchez"

	Count(newString)

	newString = "Kick Back"

	Count(newString)
}
