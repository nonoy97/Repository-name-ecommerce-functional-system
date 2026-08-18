package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ecommerce-functional-system/internal/cart"
	"ecommerce-functional-system/internal/catalog"
)

// Variables globales para simular la base de datos en memoria
var (
	inventario []catalog.Producto
	miCarrito  cart.Carrito
)

func init() {
	// Inicializamos datos de prueba
	p1, _ := catalog.NewProducto("P1", "Teclado Mecánico", "Periféricos", 80.0, 10)
	p2, _ := catalog.NewProducto("P2", "Monitor 144Hz", "Pantallas", 250.0, 5)
	inventario = append(inventario, p1, p2)
	miCarrito = cart.NewCarrito(3.50)
}

// 1. GET /productos
func getProductosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventario)
}

// 2. GET /productos/buscar?id=P1
func buscarProductoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	prod, err := catalog.BuscarPorID(inventario, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod)
}

// 3. POST /productos/nuevo
func nuevoProductoHandler(w http.ResponseWriter, r *http.Request) {
	// Simplificado para la demostración
	p3, _ := catalog.NewProducto("P3", "Mouse Gamer", "Periféricos", 45.0, 20)
	inventario = append(inventario, p3)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Producto P3 agregado exitosamente"})
}

// 4. GET /carrito
func getCarritoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(miCarrito.Items())
}

// 5. POST /carrito/agregar?id=P1&cant=1
func agregarCarritoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	prod, err := catalog.BuscarPorID(inventario, id)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	item := cart.ItemCarrito{
		IDProducto: prod.ID(),
		Nombre:     prod.Nombre(),
		PrecioUnit: prod.Precio(),
		Cantidad:   1, // Fijo por simplicidad del ejemplo
	}
	miCarrito.AgregarItem(item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Item agregado al carrito"})
}

// 6. POST /carrito/cupon?desc=0.10
func aplicarCuponHandler(w http.ResponseWriter, r *http.Request) {
	// Fijo al 10% por simplicidad de la API sin leer body
	err := miCarrito.SetCuponDescuento(0.10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Cupón del 10% aplicado"})
}

// 7. GET /carrito/procesar/ecuador
func procesarEcuadorHandler(w http.ResponseWriter, r *http.Request) {
	calculador := cart.ImpuestoEcuador{}
	resumen, err := miCarrito.ProcesarOrden(calculador)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resumen)
}

// 8. GET /carrito/procesar/exento
func procesarExentoHandler(w http.ResponseWriter, r *http.Request) {
	calculador := cart.ImpuestoCero{}
	resumen, err := miCarrito.ProcesarOrden(calculador)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resumen)
}

func main() {
	http.HandleFunc("/productos", getProductosHandler)
	http.HandleFunc("/productos/buscar", buscarProductoHandler)
	http.HandleFunc("/productos/nuevo", nuevoProductoHandler)
	http.HandleFunc("/carrito", getCarritoHandler)
	http.HandleFunc("/carrito/agregar", agregarCarritoHandler)
	http.HandleFunc("/carrito/cupon", aplicarCuponHandler)
	http.HandleFunc("/carrito/procesar/ecuador", procesarEcuadorHandler)
	http.HandleFunc("/carrito/procesar/exento", procesarExentoHandler)

	fmt.Println("=== Servidor de E-Commerce Iniciado en http://localhost:8080 ===")
	log.Fatal(http.ListenAndServe(":8080", nil))
}