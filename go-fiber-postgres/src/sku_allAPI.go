package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type SKU struct {
	SKUID           *string    `json:"skuid"`
	BarcodePOS      *string    `json:"barcodepos"`
	ProductName     *string    `json:"productname"`
	BrandID         *int64     `json:"brandid"`
	ProductGroupID  *int64     `json:"productgroupid"`
	ProductCatID    *int64     `json:"productcatid"`
	ProductSubCatID *int64     `json:"productsubcatid"`
	ProductSizeID   *int64     `json:"productsizeid"`
	ProductUnit     *int64     `json:"productunit"`
	PackSize        *string    `json:"packsize"`
	Unit            *int64     `json:"unit"`
	BanForPracharat *int64     `json:"banforpracharat"`
	IsVat           *bool      `json:"isvat"`
	CreateBy        *string    `json:"createby"`
	CreateDate      *time.Time `json:"createdate"`
	IsActive        *bool      `json:"isactive"`
	MerchantID      *string    `json:"merchantid"`
	MapSKU          *string    `json:"mapsku"`
	IsFixPrice      *bool      `json:"isfixprice"`
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
		file, err := os.Create("sku_all.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		jsonData, err := json.Marshal(skus)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, skus)
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
