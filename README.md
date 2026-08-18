# Sistema de Gestión de E-Commerce en Go

## Datos del Proyecto
* **Fecha:** Agosto 2026
* **Estudiante:** Johan de la Bastida 
* **Materia:** Programacion orientada a objetos 

## Objetivo del Programa
Desarrollar un sistema de comercio electrónico implementando los paradigmas de programación funcional y orientada a objetos en Go (Golang). El sistema gestiona un catálogo de productos inmutables y un carrito transaccional mediante la exposición de Servicios Web (API REST) que responden en formato JSON.

## Principales Funcionalidades (8 Servicios Web)
1. **Listar Catálogo:** `GET /productos` - Retorna el inventario completo.
2. **Buscar Producto:** `GET /productos/buscar?id=X` - Encuentra un ítem específico.
3. **Agregar Producto:** `POST /productos/nuevo` - Añade nuevos ítems al catálogo.
4. **Ver Carrito:** `GET /carrito` - Visualiza los ítems actuales.
5. **Agregar al Carrito:** `POST /carrito/agregar?id=X` - Introduce productos a la orden.
6. **Aplicar Cupón:** `POST /carrito/cupon` - Aplica un descuento con validación estricta de errores.
7. **Procesar Orden (IVA 15%):** `GET /carrito/procesar/ecuador` - Polimorfismo aplicando impuestos locales.
8. **Procesar Orden (Exenta):** `GET /carrito/procesar/exento` - Polimorfismo sin cobro de impuestos.

## Tecnologías y Conceptos Aplicados (Unidades 1 a 4)
* **Unidad 1 y 2:** Sintaxis, condicionales, iteraciones, slices, maps y estructuras.
* **Unidad 3:** Encapsulamiento, manejo explícito de errores y polimorfismo mediante interfaces.
* **Unidad 4:** Implementación de servidor HTTP y serialización de datos en formato JSON.