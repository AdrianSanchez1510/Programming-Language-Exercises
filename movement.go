package main

import "fmt"

type wordList []string

var newList wordList

func (l *wordList) Rotation(Dir int, numC int){
	var temp string
	
	fmt.Println("\n Lista original: ", *l)
	var i, cant int
	cant = len(*l)

	if numC < len(*l){
		if Dir == 0{
			for i = 0 ; i < numC ; i++{
				temp = (*l)[i]
				(*l) = append(*l, temp)
				(*l)[len(*l)-(1)] = (*l)[i]
			}
			(*l) = (*l)[(numC):]
		} else {
			for i = numC ; i > 0 ; i--{
				temp = (*l)[numC - i]
				(*l)[numC-i] = (*l)[len(*l) - i]
				(*l)[i] = temp
			}
			(*l) = (*l)[:cant]
		}
	}
	fmt.Println("New List: ", *l)
}

func fillData(){
	newList = append(newList, "This")
	newList = append(newList, "is")	
	newList = append(newList, "a")	
	newList = append(newList, "sample")	
	newList = append(newList, "text")
}

func main(){
	fmt.Println("---------- Rotation --------")
	fillData()
	newList.Rotation(0,3)
}
