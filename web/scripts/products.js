fetch("http://localhost:8080/products")
    .then(res => res.json())
    .then(products => {
        console.log(products);

    });