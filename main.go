package main

import (
	"fmt"
	"github.com/jerryandersonh/b-tree-go/arbol"
)


func main(){
	raiz := arbol.NuevaRaiz(10)

	raiz.insertar(5)
	raiz.insertar(12)
	raiz.insertar(13)
	raiz.insertar(12)
	raiz.insertar(2)
	raiz.insertar(9)
	raiz.insertar(15)

	//mostrar recorrido en inorden
	fmt.Print("Recorrido inorder: ")
	raiz.inOrden()
	fmt.Println()
}