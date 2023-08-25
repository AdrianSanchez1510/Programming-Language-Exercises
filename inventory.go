package main

import "fmt"

type calzado struct{
	model string
	price int
	size int
	quantity int
}

type shoelist []calzado

var shoes shoelist

func (list *shoelist) findShoe(model string, size int) int{

	for i := 0; i < len(*list); i++{
		if(*list)[i].model == model{
			if(*list)[i].size == size{
				return i
			}
		}
	}

	return 100
}


func (list *shoelist) addShoe(model string, price int, size int, quantity int){
	var shoe calzado

	fi := (*list).findShoe(model, size)

	if fi > len(*(list)){
		shoe.model = model; shoe.price = price; shoe.size = size; shoe.quantity = quantity
		(*list) = append((*list), shoe)
	} else{
		(*list)[fi].quantity ++
	}
}

func (list *shoelist) sellShoe(model string, size int, quantity int){
	fi := (*list).findShoe(model, size)
	
	if fi < len(*list){
		if (*list)[fi].quantity < quantity{
			fmt.Println("This item is out of stock")
			return
		}

		(*list)[fi].quantity = (*list)[fi].quantity - quantity
	}
}


func fillData(){
	shoes.addShoe("Nike Air Jordan", 120000, 35, 3)
	shoes.addShoe("Adiddas", 90000, 35, 32)
	shoes.addShoe("New Balance", 120000, 35, 2)
	shoes.addShoe("Levis", 120000, 41, 5)
	shoes.addShoe("Crocs", 120000, 39, 8)
	shoes.addShoe("Levis", 120000, 33, 9)
	shoes.addShoe("Nike Air Jordan", 120000, 34, 3)
	shoes.addShoe("Eagle", 120000, 37, 13)
	shoes.addShoe("Nike Air Jordan", 12000, 35, 2)
	shoes.addShoe("Adiddas", 45000, 33, 4 )
}

func main(){
	fmt.Println("------- Shoe Inventory ------")
	fillData()
	fmt.Println(shoes)
	shoes.sellShoe("Adiddas", 33, 2)
	shoes.sellShoe("Nike Air Jordan", 35, 2)
	fmt.Println(shoes)
}
