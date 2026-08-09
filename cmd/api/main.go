package main

import (
	"fmt"
	"log"

	"ecommerce-functional-system/internal/cart"
	"ecommerce-functional-system/internal/catalog"
)

func main() {
	fmt.Println("=== SISTEMA DE E-COMMERCE: DEMOSTRACIÓN COMPLETA ===")

	p1, err := catalog.NewProducto("P1", "Teclado Mecánico", "Periféricos", 80.0, 10)
	if err != nil {
		log.Fatalf("Error al crear producto: %v", err)
	}

	fmt.Printf("Producto creado: %s - Categoria: %s - Precio: $%.2f\n", p1.Nombre(), p1.Categoria(), p1.Precio())

	miCarrito := cart.NewCarrito(3.50)

	// Probando captura de error de cupón inválido (80%)
	err = miCarrito.SetCuponDescuento(0.80)
	if err != nil {
		fmt.Printf(" [Error Capturado Exitosamente]: %v\n", err)
	}

	_ = miCarrito.SetCuponDescuento(0.10)

	_ = miCarrito.AgregarItem(cart.ItemCarrito{
		IDProducto: p1.ID(),
		Nombre:     p1.Nombre(),
		PrecioUnit: p1.Precio(),
		Cantidad:   2,
	})

	// Polimorfismo mediante Interfaces
	impuestoEcuador := cart.ImpuestoEcuador{}
	impuestoCero := cart.ImpuestoCero{}

	resumenEcuador, _ := miCarrito.ProcesarOrden(impuestoEcuador)
	fmt.Printf("\n--- Orden con IVA Ecuador (15%%) ---\nSubtotal: $%.2f\nIVA: $%.2f\nTotal: $%.2f\n", resumenEcuador.Subtotal, resumenEcuador.Impuesto, resumenEcuador.Total)

	resumenCero, _ := miCarrito.ProcesarOrden(impuestoCero)
	fmt.Printf("\n--- Orden Exenta de Impuesto ---\nSubtotal: $%.2f\nIVA: $%.2f\nTotal: $%.2f\n", resumenCero.Subtotal, resumenCero.Impuesto, resumenCero.Total)
}
