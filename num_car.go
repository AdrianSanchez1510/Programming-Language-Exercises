package main

import (
	"fmt"
	"strings"
)

func characterCount(phrase string){

	fmt.Println("\nPhrase: " + phrase + "\n")
	
	var countC, countW, countI, i int 

	cha := strings.Split(phrase, "")
	wor := strings.Split(phrase, " ")
	ide := strings.Split(phrase, "\n")


	countC = 0
	countW = 0
	countI = 0

	for i = 0; i < len(cha); i++{
		countC = countC + 1
	}

	for i = 0; i < len(wor); i++{
		countW = countW + 1
	}

	for i = 0; i < len(ide); i++{
		countI = countI + 1
	}


	fmt.Println("The amount of characters for this phrase is: ", countC)
	fmt.Println("The amount of words for this phrase is: ", countW)
	fmt.Println("The amount of lines for this phrase is: ", countI)
}

func main(){
	var newString string = "Sample Text \nDon't Use it Incorrectly \nESSSS"

	characterCount(newString)

	newString = "Kick Back"

	characterCount(newString)
}
