package productsrepo

import(
	"database/sql"
)

func ProductById(id int) (Product, error){
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
	stmt, err := db.Prepared("SELECT * FROM products WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var products Product

    // Execute the prepared statement, passing in an id value for the
    // parameter whose placeholder is ?
	err := stmt.QueryRow(id).Scan(
		&products.productName, 
		&products.stock,
		&products.price,
		&products.manufacturer)
	if err != nil{
		if err == sql.ErrNoRows {
			// Handle the case of no rows returned.
			//Returnera en sida för felaktig sökning
			//Eller en sida med varor och ett felmeddelande
		}
		return product, err
	}
	return product, nil
}

//https://go.dev/doc/database/querying
//https://go.dev/doc/database/prepared-statements