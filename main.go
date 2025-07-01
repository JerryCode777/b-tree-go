package main

import (
	"fmt"
	"github.com/JerryCode777/b-tree-go/arbol"
)


func main(){
	raiz := arbol.NuevaRaiz(10)

	raiz.Insertar(5)
	raiz.Insertar(12)
	raiz.Insertar(13)
	raiz.Insertar(12)
	raiz.Insertar(2)
	raiz.Insertar(9)
	raiz.Insertar(15)

	//mostrar recorrido en inorden
	fmt.Print("Recorrido inorder: ")
	raiz.InOrden()
	fmt.Println()
}