package cafeteria

import (
	"errors"
)

type RepoMemoria struct {
	clientes  []Cliente
	productos []Producto
	pedidos   []Pedido
}

type Cliente struct {
	ID     int
	Nombre string
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID       int
	Cliente  Cliente
	Producto Producto
	Cantidad int
	Total    float64
}

var (
	ErrCliente  = errors.New("Cliente no encontrado")
	ErrProducto = errors.New("Producto no encontrado")
	ErrPedido   = errors.New("Pedido no encontrado")
	ErrStock    = errors.New("Stock insuficiente")
	ErrSaldo    = errors.New("Saldo insuficiente")
	ErrTotal    = errors.New("El total debe ser mayor a cero")
)

type Repository interface {
	GuardarCliente(cliente Cliente) error
	ObtenerCliente(id int) (Cliente, error)
	ListarClientes() []Cliente

	GuardarProducto(producto Producto) error
	ObtenerProducto(id int) (Producto, error)
	ListarProductos() []Producto
}

func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{}
}
func (r *RepoMemoria) GuardarCliente(c Cliente) error {
	r.clientes = append(r.clientes, c)
	return nil
}

func (r *RepoMemoria) ObtenerCliente(id int) (Cliente, error) {
	for _, c := range r.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return Cliente{}, ErrCliente
}

func (r *RepoMemoria) ListarClientes() []Cliente {
	return r.clientes
}

func (r *RepoMemoria) GuardarProducto(p Producto) error {
	r.productos = append(r.productos, p)
	return nil
}

func (r *RepoMemoria) ObtenerProducto(id int) (Producto, error) {
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProducto
}

func (r *RepoMemoria) ListarProductos() []Producto {
	return r.productos
}

// Verificación en tiempo de compilación:
// Si RepoMemoria NO cumple Repository, esto da error al compilar.
var _ Repository = (*RepoMemoria)(nil)
