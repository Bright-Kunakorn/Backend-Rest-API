package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://root:secret@localhost:5433?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL database!")

	rows, err := db.Query("SELECT * FROM backendposdatasku")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var skuid string
	var barcodepos string
	var productname string
	var brandid int
	var productgroupid int
	var productcatid int
	var productsubcatid int
	var productsizeid int
	var productunit int
	var packsize string
	var unit int
	var banforpracharat int
	var isvat bool
	var createby string
	var createdate time.Time
	var isactive bool
	var merchantid string
	var mapsku string
	var isfixprice bool

	for rows.Next() {
		err := rows.Scan(&skuid, &barcodepos, &productname, &brandid, &productgroupid, &productcatid, &productsubcatid, &productsizeid, &productunit, &packsize, &unit, &banforpracharat, &isvat, &createby, &createdate, &isactive, &merchantid, &mapsku, &isfixprice)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s | %s | %s | %d | %d | %d | %d | %d | %d | %s | %d | %d | %t | %s | %v | %t | %s | %s | %t\n", skuid, barcodepos, productname, brandid, productgroupid, productcatid, productsubcatid, productsizeid, productunit, packsize, unit, banforpracharat, isvat, createby, createdate, isactive, merchantid, mapsku, isfixprice)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
