// App Web II - S1 (sesión de Taller)
// Nivel y Paralelo: 6to "B"
// Tema:Variables, tipos, control de flujo y bucles en Go
// Autoras: Anchundia Pillasagua Karen, Moreira Garcia Marly
// Fecha: 09 de abril del 2026

// Calculadora Cíentifica CLI en Go
package main

import "fmt"

// Paso 4 (No se puede poner funciones dentro de main, se deben declarar fuera)
// Potencia y Factorial
// Función para calcular la potencia
func potencia(base int, exponente int) int {
	var resultado int = 1
	for i := 0; i < exponente; i++ { // For clasico
		resultado *= base
	}
	return resultado
}

// Función para calcular el factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	resultado := 1 // Acumulador
	i := 1

	for i <= n { // For while
		resultado *= i
		i++
	}
	return resultado
}

// Historial de operaciones
func mostrarHistorial(historial string, contador int) {
	fmt.Println("\n-------------------- Historial de Operaciones --------------------")
	fmt.Print(historial)
	fmt.Printf("Total de operaciones realizadas: %d\n", contador)
}

func main() {
	// Paso 1 y 2
	fmt.Println("-------------------- Calculadora Científica --------------------")
	// Operaciones básicas
	//suma := a + b
	//resta := a - b
	//multiplicacion := a * b
	//division := float64(a) / float64(b)
	// Operaciones avanzadas
	//resultadopotencia := potencia(float64(a), b)
	//factorialA := factorial(a)
	//factorialB := factorial(b)

	// Resultados
	//fmt.Println("a + b =", suma)
	//fmt.Println("a - b =", resta)
	//fmt.Println("a * b =", multiplicacion)
	//fmt.Printf("a / b = %.2f\n", division)
	//fmt.Printf("a^b = %.2f\n", resultadopotencia)
	//fmt.Printf("Factorial de a: %d! = %d\n", a, factorialA)
	//fmt.Printf("Factorial de b: %d! = %d\n", b, factorialB)

	// Paso 5
	historial := ""
	contador := 0

	// Paso 3 (For Infinito)
	// Variables para operaciones
	for {

		var a int
		var b int
		var operacion string

		fmt.Print("\nSeleccione la operación (+, -, *, /, ^, !): ")
		fmt.Scanln(&operacion)

		var resultadoTexto string

		// Factorial
		if operacion == "!" {
			var n int
			fmt.Print("Ingrese el número: ")
			fmt.Scanln(&n)

			if n < 0 {
				fmt.Println("Error: No hay factorial de negativos")
				continue
			}

			res := factorial(n)
			resultadoTexto = fmt.Sprintf("%d! = %d", n, res)
			fmt.Println("Resultado:", resultadoTexto)

		} else {
			// Operaciones básicas y potencia
			fmt.Print("Ingrese el primer número: ")
			fmt.Scanln(&a)

			fmt.Print("Ingrese el segundo número: ")
			fmt.Scanln(&b)

			// Switch para operaciones
			switch operacion {
			case "+":
				res := int(a + b)
				resultadoTexto = fmt.Sprintf("%d + %d = %.d", a, b, res)

			case "-":
				res := int(a - b)
				resultadoTexto = fmt.Sprintf("%d - %d = %.d", a, b, res)

			case "*":
				res := float64(a * b)
				resultadoTexto = fmt.Sprintf("%d * %d = %.2f", a, b, res)

			case "/":
				if b == 0 {
					fmt.Println("Error: división por cero")
					continue
				}
				res := float64(a) / float64(b)
				resultadoTexto = fmt.Sprintf("%d / %d = %.2f", a, b, res)

			case "^":
				res := potencia(int(a), b)
				resultadoTexto = fmt.Sprintf("%d ^ %d = %d", a, b, res)

			default:
				fmt.Println("Operación no válida")
				continue
			}

			fmt.Println("Resultado:", resultadoTexto)
		}

		// Paso 5 Historial de operaciones
		contador++
		historial += fmt.Sprintf("%d) %s\n", contador, resultadoTexto)

		var opcion string
		fmt.Print("¿Otra operación? (s/n): ")
		fmt.Scanln(&opcion)

		if opcion == "n" {
			break
		}
	}

	// Mostrar historial al final
	mostrarHistorial(historial, contador)
}
