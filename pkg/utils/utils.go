package utils

import (
	"Test2/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseBody(r *http.Request) models.Book {
	newBook := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		fmt.Println("Error during parse of request body")
	}
	return newBook
}
