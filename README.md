# ProductGraphQlExample
<p>list</p>
<p>curl -g 'http://localhost:8080/graphql?query={productList{id,name,price}}'</p>
<p>create</p>
<p>curl -g 'http://localhost:8080/graphql?query=mutation+_{createProduct(name:"product3333",price:"123"){id,name,price}}'</p>
<p>update</p>
<p>curl -g 'http://localhost:8080/graphql?query=mutation+_{updateProduct(name:"My+new_product1231231",id:"1"){id,name,price}}'</p>
<p>get by id</p>
<p>curl -g 'http://localhost:8080/graphql?query={product(id:"1"){id,name,price}}'</p>