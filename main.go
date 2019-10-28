package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "northwind"
	dbPass = "admin"
	dbHost = "localhost"
	dbPort = "3306"
)

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func routers() *chi.Mux {
	router.Get("/products", AllProductos)
	router.Post("/products", CreateProducto)
	router.Put("/products/{id}", UpdateProducto)
	router.Delete("/products/{id}", DeleteProducto)

	return router
}

func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id,product_code,COALESCE(description,'')
				 FROM products`
	results, err := db.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}
	respondwithJSON(w, http.StatusOK, products)
}

func CreateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := db.Prepare("Insert products SET product_code=?, description=?")
	catch(err)

	_, er := query.Exec(producto.Product_Code, producto.Description)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var product Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&product)

	query, err := db.Prepare("Update products set product_code=?, description=? where id=?")
	catch(err)
	_, er := query.Exec(product.Product_Code, product.Description, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("delete from products where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dbSource)

	catch(err)
}

func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}
