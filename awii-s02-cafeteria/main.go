// App Web II - S2 (sesión de Taller)
// Nivel y Paralelo: 6to "B"
// Autoras: Anchundia Pillasagua Karen, Moreira Garcia Marly
// Fecha: 16 de abril del 2026

package main

import "fmt"

// Primer Checkpoint
// Estructura
type Cliente struct {
	id      int
	nombre  string
	carrera string
	saldo   float64
}
type Producto struct {
	id        int
	nombre    string
	precio    float64
	stock     int
	categoria string
}
type Pedido struct {
	id          int
	clienteID   int
	productosID int
	cantidad    int
	total       float64
	fecha       string
}

// Lista de clientes
func ListarClientes(clientes []Cliente) {
	if len(clientes) == 0 {
		fmt.Println("(No hay clientes registrados)")
		return
	}
	fmt.Println("---------------------- Lista de Clientes ----------------------")
	fmt.Println("---------------------------------------------------------------")
	// Encabezado
	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Nombre", "Carrera", "Saldo")
	fmt.Println("---------------------------------------------------------------")

	// Datos
	for _, c := range clientes {
		fmt.Printf("%-5d %-20s %-15s $%-10.2f\n",
			c.id, c.nombre, c.carrera, c.saldo)
	}
}

// Main
func main() {
	var clientes []Cliente
	var productos []Producto
	var pedidos []Pedido
	// Agregar clientes
	clientes = []Cliente{
		{1, "Marly Moreira", "TI", 25.50},
		{2, "Karen Anchundia", "Sotfware", 15.00},
		{3, "Mara García", "Educación", 40.80},
	}
	// Agregar productos
	productos = []Producto{
		{1, "CocaCola", 1.50, 100, "Bebidas"},
		{2, "Papas Fritas", 2.00, 50, "Snacks"},
		{3, "Chocolate", 1.00, 200, "Dulces"},
		{4, "Sanduche", 3.50, 15, "Comida"},
	}
	// pedidos queda vacío
	pedidos = []Pedido{}
	// Datos
	ListarClientes(clientes)
	fmt.Println("\nAgregando cliente...")
	clientes = AgregarCliente(clientes, Cliente{4, "Yelena Pillasagua", "Administración", 30.00})
	ListarClientes(clientes)
	fmt.Println("\nEliminando cliente ID 2...")
	clientes = EliminarCliente(clientes, 2)
	ListarClientes(clientes)
	fmt.Println("\nBuscando cliente por ID...")
	index := BuscarClientePorID(clientes, 6)
	if index != -1 {
		fmt.Printf("Cliente encontrado: %s\n", clientes[index].nombre)
	} else {
		fmt.Println("Cliente no encontrado")
	}
	fmt.Println("Productos cargados:", len(productos))
	ListarProductos(productos)
	fmt.Println("\nAgregando producto...")
	productos = AgregarProducto(productos, Producto{5, "Galletas", 1.20, 80, "Snacks"})
	ListarProductos(productos)
	fmt.Println("\nEliminando producto ID 3...")
	productos = EliminarProducto(productos, 3)
	ListarProductos(productos)
	fmt.Println("\nBuscando producto por ID...")
	indexP := BuscarProductoPorID(productos, 4)
	if indexP != -1 {
		fmt.Printf("Producto encontrado: %s\n", productos[indexP].nombre)
	} else {
		fmt.Println("Producto no encontrado")
	}
	fmt.Println("Pedidos registrados:", len(pedidos))
}

// Segundo Checkpoint
// CRUD Cliente y Producto
func AgregarCliente(clientes []Cliente, nuevoCliente Cliente) []Cliente {
	clientes = append(clientes, nuevoCliente)
	return clientes
}
func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.id == id {
			return i
		}
	}
	return -1
}
func EliminarCliente(clientes []Cliente, id int) []Cliente {
	index := BuscarClientePorID(clientes, id)
	if index == -1 {
		fmt.Println("Cliente no encontrado")
		return clientes
	}
	return append(clientes[:index], clientes[index+1:]...)
}

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	productos = append(productos, nuevo)
	return productos
}
func BuscarProductoPorID(productos []Producto, id int) int {
	for i, p := range productos {
		if p.id == id {
			return i
		}
	}
	return -1
}
func ListarProductos(productos []Producto) {
	if len(productos) == 0 {
		fmt.Println("(No hay productos registrados)")
		return
	}
	fmt.Println("---------------------- Lista de Productos ----------------------")
	fmt.Println("----------------------------------------------------------------")
	// Encabezado
	fmt.Printf("%-5s %-20s %-10s %-5s %-15s\n", "ID", "Nombre", "Precio", "Stock", "Categoría")
	fmt.Println("----------------------------------------------------------------")

	// Datos
	for _, p := range productos {
		fmt.Printf("%-5d %-20s $%-10.2f %-5d %-15s\n",
			p.id, p.nombre, p.precio, p.stock, p.categoria)
	}
}
func EliminarProducto(productos []Producto, id int) []Producto {
	index := BuscarProductoPorID(productos, id)
	if index == -1 {
		fmt.Println("Producto no encontrado")
		return productos
	}
	return append(productos[:index], productos[index+1:]...)
}
