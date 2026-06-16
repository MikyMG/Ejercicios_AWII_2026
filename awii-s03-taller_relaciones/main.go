package main

import (
	"fmt"

	"github.com/MikyMG/Ejercicios_AWII_2026/awii-s03-taller_relaciones/internal/cafeteria"
)

func main() {

	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()

	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Ana"})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Pepe"})

	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Agua", Precio: 1, Stock: 60, Categoria: "Bebida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "Pan", Precio: 2, Stock: 40, Categoria: "Comida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Cafe", Precio: 3, Stock: 90, Categoria: "Bebida"})

	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Printf("Error al buscar cliente: %s\n", err.Error())
	} else {
		fmt.Printf("\nEncontrado: %s\n", c.Nombre)
	}

	fantasma, err := repo.ObtenerCliente(3)
	if err != nil {
		fmt.Printf("Error al buscar cliente: %s\n", err.Error())
	} else {
		fmt.Println(fantasma)
	}

	fmt.Printf("Listando Productos:\n")
	for _, p := range repo.ListarProductos() {
		fmt.Printf("%v\n", p)
	}
}
