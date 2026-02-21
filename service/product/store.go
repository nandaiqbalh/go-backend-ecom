package product

import (
	"database/sql"

	"github.com/nandaiqbalh/go-backend-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// function to query products in db (list, get by id, create, update, delete) would be defined here, following similar patterns to the user store: execute SQL queries and scan results into types.Product structs.
func (s *Store) ListProducts() ([]*types.Product, error) {
	// Implementation for listing products from the database
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*types.Product

	for rows.Next() {
        product := new(types.Product)
        var img sql.NullString
        err := rows.Scan(
            &product.ID,
            &product.Name,
            &product.Description,
            &img,
            &product.Price,
            &product.Quantity,
            &product.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        if img.Valid {
            product.Image = img.String
        }
        products = append(products, product)
    }
	return products, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	// Implementation for getting a single product by ID from the database
	row := s.db.QueryRow("SELECT * FROM products WHERE id = ?", id)

	product := new(types.Product)
    var img sql.NullString
	err := row.Scan(
        &product.ID,
        &product.Name,
        &product.Description,
        &img,
        &product.Price,
        &product.Quantity,
        &product.CreatedAt,
    )
    if img.Valid {
        product.Image = img.String
    }
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No product found with the given ID
		}
		return nil, err
	}

	return product, nil
}

func (s *Store) CreateProduct(product *types.Product) error {
	// Implementation for creating a new product in the database
	result, err := s.db.Exec(
		"INSERT INTO products (name, description, price, quantity) VALUES (?, ?, ?, ?)",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = int(id)

	return nil
}

func (s *Store) UpdateProduct(product *types.Product) error {
	// Implementation for updating an existing product in the database
	_, err := s.db.Exec(
		"UPDATE products SET name = ?, description = ?, price = ?, quantity = ? WHERE id = ?",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.ID,
	)
	return err
}

func (s *Store) DeleteProduct(id int) error {
	// Implementation for deleting a product from the database
	_, err := s.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}