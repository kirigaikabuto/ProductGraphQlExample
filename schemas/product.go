package schemas

import (
	"errors"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"strconv"
)

var (
	ProductsList []Product
)

type Product struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty'"`
}

var productGraphQlType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createProduct": &graphql.Field{
			Type:        productGraphQlType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["name"].(string)
				price, _ := strconv.Atoi(params.Args["price"].(string))
				newProduct := Product{
					Id:    uuid.New().String(),
					Name:  name,
					Price: price,
				}
				ProductsList = append(ProductsList, newProduct)
				return newProduct, nil
			},
		},
		"updateProduct": &graphql.Field{
			Type:        productGraphQlType,
			Description: "Update product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["name"].(string)
				price, _ := params.Args["price"].(string)
				id, _ := params.Args["id"]
				updatedProduct := Product{}
				for i, v := range ProductsList {
					if v.Id == id {
						if name != "" {
							ProductsList[i].Name = name
						}
						if price != "" {
							intPrice, _ := strconv.Atoi(price)
							ProductsList[i].Price = intPrice
						}
						updatedProduct = ProductsList[i]
						break
					}
				}
				return updatedProduct, nil
			},
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"productList": &graphql.Field{
			Type:        graphql.NewList(productGraphQlType),
			Description: "List of products",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return ProductsList, nil
			},
		},
		"product": &graphql.Field{
			Type:        productGraphQlType,
			Description: "Get product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(string)
				for _, v := range ProductsList {
					if v.Id == id {
						return v, nil
					}
				}
				return nil, errors.New("Product not found")
			},
		},
	},
})

var ProductSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
