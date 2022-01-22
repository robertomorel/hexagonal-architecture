package db

import (
	"database/sql"

	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

/*
	Este adapter precisa implementar a interface ProductPersistenceInterface
	Estamos aqui trabalhando com o lado "direito" da arquitetura hexagonal, para trabalhar com o DB
*/

// Iniciando o struct para este adaptador
type ProductDb struct {
	db *sql.DB
}

// Criando nova instância do Product DB
func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}

	// O Scan relaciona cada item que estou buscando com o atributo do produto application.Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("Select id from products where id=?", product.GetID()).Scan(&rows)
	if rows == 0 {
		// Se o produto não existir, cria
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		// Se existir, atualiza
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

// "create" está como letra minúscula para tornar ele privado
func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

// "update" está como letra minúscula para tornar ele privado
func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec("update products set name = ?, price=?, status=? where id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil
}
