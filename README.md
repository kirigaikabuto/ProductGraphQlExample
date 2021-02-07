# ProductGraphQlExample
list
curl -g 'http://localhost:8080/graphql?query={productList{id,name,price}}'
create
curl -g 'http://localhost:8080/graphql?query=mutation+_{createProduct(name:"product3333",price:"123"){id,name,price}}'
update
curl -g 'http://localhost:8080/graphql?query=mutation+_{updateProduct(name:"My+new_product1231231",id:"1"){id,name,price}}'
get by id
curl -g 'http://localhost:8080/graphql?query={product(id:"1"){id,name,price}}'