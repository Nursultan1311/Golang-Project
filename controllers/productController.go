package controllers

import (
	"Goland/database"
	"Goland/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func ShowProducts(res http.ResponseWriter, req *http.Request) { // we get all products from database with by GetProductsfubction
	tpl, _ = tpl.ParseGlob("templates/*.html")
	tpl.ExecuteTemplate(res, "showProducts.html", GetProducts())
	return
}

func GetProducts() []models.Product { // Getting all products from database
	db = database.ConnectToDB() // connecting to database

	rows, err := db.Query("SELECT name, description, price, quantity FROM products") // selecting all products
	defer rows.Close()
	var products []models.Product // declaration of slice 
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Name, &product.Description, &product.Price, &product.Quantity); err != nil {
			return nil
		}
		products = append(products, product) // adding each products for products slice
	}

	if err != nil {
		log.Fatal(err)
	}
	return products // then return products for showProducts function
}

func GetProduct(res http.ResponseWriter, req *http.Request) { // Getting all product with by search input from header
	db = database.ConnectToDB() // connecting to database

	name := req.FormValue("Target") // getting target value from request post with by name Target

	fmt.Println(name)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, description, price, quantity FROM products") // selecting all products to slice
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Name, &product.Description, &product.Price, &product.Quantity); err != nil {
			fmt.Println("error in rows")
		}
		products = append(products, product)
	}

	var result []models.Product
	for _, i := range products {
		// comparing product name with target value from search
		if strings.Contains(strings.ToLower(i.Name), strings.ToLower(name)) || strings.ToLower(i.Name) == strings.ToLower(name) {
			result = append(result, i) // if products exist with target name, then we add it to result slice
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	tpl.ExecuteTemplate(res, "showProducts.html", result) // then return all products with target value name
}

func AddProduct(res http.ResponseWriter, req *http.Request) { // adding a new product from website to database
	db = database.ConnectToDB() // connecting

	if req.Method != "POST" {
		http.ServeFile(res, req, "templates/add_product.html")
		return
	}
	// getting data from form
	description := req.FormValue("description")
	price := req.FormValue("price")
	quantity := req.FormValue("quantity")
	name := req.FormValue("name")
	// inserting a new data to database
	_, err = db.Exec("INSERT INTO products(name, description, price, quantity) VALUES(?, ?, ?, ?)", name, description, price, quantity)
	if err != nil {
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	}
	http.Redirect(res, req, "/show_products", 301) // redirecting the page
}
