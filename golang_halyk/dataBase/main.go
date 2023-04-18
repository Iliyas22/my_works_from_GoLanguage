package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	type product struct {
		product_no int
		name       string
		price      float64
	}
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "55500ilias"
		dbname   = "postgres"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Println("connected")
	//1)Добавление
	//	sqlStatement := `
	//INSERT INTO products (product_no, name, price)
	//VALUES ($1, $2, $3)
	//RETURNING product_no`
	//	product_no := 0
	//	err = db.QueryRow(sqlStatement, 7, "Salad", 550).Scan(&product_no)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("Add New product is:", sqlStatement)

	//	2)Обновление, изменение
	//sqlStatement := `update products set  name=$2, price=$3 WHERE product_no=$1 returning product_no
	//`
	//product_no := 0
	//err = db.QueryRow(sqlStatement, 2, "pizza", 1750).Scan(&product_no)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("New Products is", sqlStatement)

	//	deleteID := `
	//DELETE FROM products
	//WHERE product_no= $1;`
	//	res, err := db.Exec(deleteID, 1)
	//	if err != nil {
	//		panic(err)
	//	}

	//count, err := res.RowsAffected()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(count)
	//--- вышла ошибка
	//	req := product{
	//		product_no: 4,
	//		name: "meat",
	//		price: 1000,
	//	}
	//	if err, _ = db.QueryRow('
	//		INSERT INTO
	//		products(
	//			"product_no",
	//			"name",
	//			"price",
	//		)
	//		VALUES ($1,$2,$3),
	//	&req.product_no,
	//	&req.name,
	//	&req.price,'
	//)
	//	fmt.Println(req)

	rows, err := db.Query("select * from products")

	if err != nil && err.Error() != "sq: no rows in result set" {
		log.Println(err.Error())
		return
	}

	for rows.Next() {
		buffer := new(product)
		err := rows.Scan(&buffer.product_no, &buffer.name, &buffer.price)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Println(buffer)
	}
}
