package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type SKU struct {
	SKUID           string    `json:"skuid"`
	BarcodePOS      string    `json:"barcodepos"`
	ProductName     string    `json:"productname"`
	BrandID         int       `json:"brandid"`
	ProductGroupID  int       `json:"productgroupid"`
	ProductCatID    int       `json:"productcatid"`
	ProductSubCatID int       `json:"productsubcatid"`
	ProductSizeID   int       `json:"productsizeid"`
	ProductUnit     int       `json:"productunit"`
	PackSize        string    `json:"packsize"`
	Unit            int       `json:"unit"`
	BanForPracharat int       `json:"banforpracharat"`
	IsVat           bool      `json:"isvat"`
	CreateBy        string    `json:"createby"`
	CreateDate      time.Time `json:"createdate"`
	IsActive        bool      `json:"isactive"`
	MerchantID      string    `json:"merchantid"`
	MapSKU          string    `json:"mapsku"`
	IsFixPrice      bool      `json:"isfixprice"`
}

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
	log.Println("Successfully connected to PostgreSQL database!")

	r := gin.Default()

	r.GET("/skus", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM backendposdatasku")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var skus []SKU

		for rows.Next() {
			var sku SKU
			err := rows.Scan(&sku.SKUID, &sku.BarcodePOS, &sku.ProductName, &sku.BrandID, &sku.ProductGroupID, &sku.ProductCatID, &sku.ProductSubCatID, &sku.ProductSizeID, &sku.ProductUnit, &sku.PackSize, &sku.Unit, &sku.BanForPracharat, &sku.IsVat, &sku.CreateBy, &sku.CreateDate, &sku.IsActive, &sku.MerchantID, &sku.MapSKU, &sku.IsFixPrice)
			if err != nil {
				log.Fatal(err)
			}
			skus = append(skus, sku)
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, skus)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
