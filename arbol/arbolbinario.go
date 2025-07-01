package arbol

import "fmt"

//nodo representa a un nodo del arbol binario
type Nodo struct {
	Valor int
	Izquierdo *Nodo
	Derecho *Nodo
}

//insertar agrega un nuevo elemento al arbol
func (n *Nodo) Insertar(valor int) {
	if valor <= n.Valor {
		if n.Izquierdo == nil {
			n.Izquierdo = &Nodo{Valor: valor}
		} else {
			n.Izquierdo.Insertar(valor)
		}
	} else {
		if n.Derecho == nil {
			n.Derecho = &Nodo{Valor: valor}
		} else {
			n.Derecho.Insertar(valor)
		}
	}
}

//InOrden imprime los valores del arbol en orden ascendente
func (n *Nodo) InOrden() {
	if n == nil {
		return
	}
	n.Izquierdo.InOrden()
	fmt.Print(n.Valor, " ")
	n.Derecho.InOrden()
}

//NuevaRaiz crea un nuevo nodo raiz del arbol
func NuevaRaiz(valor int) *Nodo {
	return &Nodo{Valor: valor}
}
