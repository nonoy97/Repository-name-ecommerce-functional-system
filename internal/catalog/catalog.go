package catalog

import (
	"errors"
	"strings"

	"github.com/samber/lo"
)

var ErrProductoNoEncontrado = errors.New("el producto solicitado no existe")

// Encapsulación (U3): Campos privados
type Producto struct {
	id         string
	nombre     string
	categoria  string
	precio     float64
	stock      int
	disponible bool
}

// Constructor con validación (U2 - U3)
func NewProducto(id, nombre, categoria string, precio float64, stock int) (Producto, error) {
	if precio < 0 {
		return Producto{}, errors.New("el precio no puede ser negativo")
	}
	return Producto{
		id:         id,
		nombre:     nombre,
		categoria:  categoria,
		precio:     precio,
		stock:      stock,
		disponible: stock > 0,
	}, nil
}

// Getters Públicos (U3)
func (p Producto) ID() string        { return p.id }
func (p Producto) Nombre() string    { return p.nombre }
func (p Producto) Categoria() string { return p.categoria }
func (p Producto) Precio() float64   { return p.precio }
func (p Producto) Stock() int        { return p.stock }

// Setter con validación de error (U3)
func (p *Producto) SetPrecio(nuevoPrecio float64) error {
	if nuevoPrecio < 0 {
		return errors.New("el precio debe ser positivo")
	}
	p.precio = nuevoPrecio
	return nil
}

// Búsqueda pura con manejo de error (U3)
func BuscarPorID(productos []Producto, id string) (Producto, error) {
	prod, ok := lo.Find(productos, func(p Producto) bool {
		return p.id == id
	})
	if !ok {
		return Producto{}, ErrProductoNoEncontrado
	}
	return prod, nil
}

// Filtrado con Slices (U2)
func FiltrarPorCategoria(productos []Producto, categoria string) []Producto {
	return lo.Filter(productos, func(p Producto, _ int) bool {
		return strings.EqualFold(p.categoria, categoria)
	})
}