package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {

	var i int
	var newie *producto
	newie, i = l.buscarProducto(nombre)

	var newProduct producto
	newProduct.nombre = nombre; newProduct.cantidad = cantidad; newProduct.precio = precio

	if i != 0{
		*l = append((*l), newProduct)
	} else{
		newie.cantidad = newie.cantidad + cantidad
	}
}


func (l *listaProductos) agregarProductos(item ...producto){
	for _, objeto := range(item){
		lProductos.agregarProducto(objeto.nombre, 1, 1)
	} 
}


func (l *listaProductos) buscarProducto(nombre string) (*producto ,int) { //el retorno es el índice del producto encontrado y -1 si no existe
	var i int
	var err int = -1
	var item *producto = nil
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			err = 0
			item = &((*l)[i])
		}
	}

	return item, err
}


func (l *listaProductos) venderProducto(nombre string, cant int) {
	var p, prod = l.buscarProducto(nombre)
	if prod != 1 {
		p.cantidad = p.cantidad - cant
	} else {
		fmt.Println("This item is out of stock")
	}
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 2, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 7, 1200)
	lProductos.agregarProducto("café", 23, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	return nil
}

//------------------------------ Here lies the Practice --------------------------"
//I did the comments in English so there can be a way to differentiate between the original comments

func (l listaProductos) Len() int{ //The three following functions are needed to apply the Sort() function
	return len(l)
}

func (l listaProductos) Less(i, j int) bool {
    return l[i].precio > l[j].precio
}

func (l listaProductos) Swap(i, j int) {
    (l)[i], (l)[j] = (l)[j], (l)[i]
}
func (l *listaProductos) aumentarInventarioDeMinimos(){ //Is better to no use any list

	fmt.Println("\n")
	for i := 0;  i < len((*l)); i++{
		if (*l)[i].cantidad < 8{
			fmt.Println("The following product's stock has been restored: ", (*l)[i].nombre)
			(*l)[i].cantidad = 10 //Increases the quantity of the product
		}
	}
	fmt.Println("List of Products (Updated): ", (*l))
}

func (l *listaProductos) ordenarLista(){

	sort.Sort((*l))
	fmt.Println("\nThe list has been sorted from the most expensive to the cheapest")
	fmt.Println(lProductos)
}


func main() {
	llenarDatos()
	lProductos.agregarProducto("leche", 8, 12200)
	fmt.Println("The original list:")
	fmt.Println(lProductos)
	lProductos.ordenarLista()
	lProductos.aumentarInventarioDeMinimos()
}
