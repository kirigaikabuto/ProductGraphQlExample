package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	schema "github.com/kirigaikabuto/ProductGraphQlExample/schemas"
	"net/http"
)

func init() {
	product1 := schema.Product{"1", "product1", 1300}
	product2 := schema.Product{"2", "product2", 1400}
	product3 := schema.Product{"3", "product3", 1500}
	schema.ProductsList = append(schema.ProductsList, product1, product2, product3)
}
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	//list
	//curl -g 'http://localhost:8080/graphql?query={productList{id,name,price}}'
	//create
	//curl -g 'http://localhost:8080/graphql?query=mutation+_{createProduct(name:"product3333",price:"123"){id,name,price}}'
	//update
	//curl -g 'http://localhost:8080/graphql?query=mutation+_{updateProduct(name:"My+new_product1231231",id:"1"){id,name,price}}'
	//get by id
	//curl -g 'http://localhost:8080/graphql?query={product(id:"1"){id,name,price}}'
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema.ProductSchema)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080", nil)
}
