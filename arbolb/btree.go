package arbolb

//BNode representa un nodo del arbol B
type BNode struct {
	Leaf bool // si es hoja
	keys []int // claves del nodo 
	Children []*BNode //hijos del nodo
}

//BTree representa al arbol B completo	
type BTree struct {
	Root *BNode //nodo raiz
	T int // Grado minimo
}

//NuevaBTree crea un arbol B con grado minimo t
func NuevaBTree(t int) *BTree {
	return &BTree{
		Root: &BNode{Leaf: true},
		T: t,
	}
}

//Insertar inserta una clave en el arbol B
//logica de insercion incluye dividir el nodo si esta lleno, Insertar recursivamente en el nodo hijo adecuado
func (t *BTree) Insertar(k int){
	r := t.Root 
	if len(r.Keys) == 2 * t.T -1 {
		s := &BNode{Leaf: false}
		s.Children = append(s.Children, r)
		t.Root = s
		t.dividirHijo(s, 0)
		t.insertNoLleno(s, k)
	} else {
		t.insertNoLleno(r, k)
	}
}
//implementar insernolleno y diividr hijo
//...