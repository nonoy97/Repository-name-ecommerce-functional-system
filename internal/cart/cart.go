package cart

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
)

// Manejo de Errores
var (
	ErrCarritoVacio     = errors.New("el carrito está vacío")
	ErrCantidadInvalida = errors.New("la cantidad debe ser mayor a cero")
	ErrCuponInvalido    = errors.New("el cupón sobrepasa el límite permitido (0% - 50%)")
)

// Interfaz para Impuestos
type CalculadorImpuestos interface {
	Calcular(monto float64) float64
}

// Implementación IVA Ecuador 15%
type ImpuestoEcuador struct{}

func (e ImpuestoEcuador) Calcular(monto float64) float64 {
	return monto * 0.15
}

// Implementación Exento de Impuestos
type ImpuestoCero struct{}

func (c ImpuestoCero) Calcular(monto float64) float64 {
	return 0.0
}

type ItemCarrito struct {
	IDProducto string
	Nombre     string
	PrecioUnit float64
	Cantidad   int
}

// Encapsulación
type Carrito struct {
	items          []ItemCarrito
	cuponDescuento float64
	tarifaEnvio    float64
}

type ResumenOrden struct {
	Subtotal  float64
	Descuento float64
	Impuesto  float64
	Envio     float64
	Total     float64
}

func NewCarrito(tarifaEnvio float64) Carrito {
	return Carrito{
		items:          make([]ItemCarrito, 0),
		cuponDescuento: 0.0,
		tarifaEnvio:    tarifaEnvio,
	}
}

// Getter público para leer los items encapsulados (U3)
func (c Carrito) Items() []ItemCarrito {
	return c.items
}

// Setter con validación de error
func (c *Carrito) SetCuponDescuento(porcentaje float64) error {
	if porcentaje < 0.0 || porcentaje > 0.50 {
		return fmt.Errorf("%w: %.2f%% no es válido", ErrCuponInvalido, porcentaje*100)
	}
	c.cuponDescuento = porcentaje
	return nil
}

func (c *Carrito) AgregarItem(item ItemCarrito) error {
	if item.Cantidad <= 0 {
		return ErrCantidadInvalida
	}
	c.items = append(c.items, item)
	return nil
}

// Uso polimórfico de la Interfaz
func (c Carrito) ProcesarOrden(calcImpuesto CalculadorImpuestos) (ResumenOrden, error) {
	if len(c.items) == 0 {
		return ResumenOrden{}, ErrCarritoVacio
	}

	subtotal := lo.Reduce(c.items, func(acc float64, item ItemCarrito, _ int) float64 {
		return acc + (item.PrecioUnit * float64(item.Cantidad))
	}, 0.0)

	montoDescuento := subtotal * c.cuponDescuento
	subtotalConDescuento := subtotal - montoDescuento
	montoImpuesto := calcImpuesto.Calcular(subtotalConDescuento)
	total := subtotalConDescuento + montoImpuesto + c.tarifaEnvio

	return ResumenOrden{
		Subtotal:  subtotal,
		Descuento: montoDescuento,
		Impuesto:  montoImpuesto,
		Envio:     c.tarifaEnvio,
		Total:     total,
	}, nil
}
