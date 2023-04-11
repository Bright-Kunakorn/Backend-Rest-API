package main

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	_ "github.com/lib/pq"
)

type SKU struct {
	SKUID           sql.NullString `json:"skuid"`
	BarcodePOS      sql.NullString `json:"barcodepos"`
	ProductName     sql.NullString `json:"productname"`
	BrandID         sql.NullInt64  `json:"brandid"`
	ProductGroupID  sql.NullInt64  `json:"productgroupid"`
	ProductCatID    sql.NullInt64  `json:"productcatid"`
	ProductSubCatID sql.NullInt64  `json:"productsubcatid"`
	ProductSizeID   sql.NullInt64  `json:"productsizeid"`
	ProductUnit     sql.NullInt64  `json:"productunit"`
	PackSize        sql.NullString `json:"packsize"`
	Unit            sql.NullInt64  `json:"unit"`
	BanForPracharat sql.NullInt64  `json:"banforpracharat"`
	IsVat           bool           `json:"isvat"`
	CreateBy        sql.NullString `json:"createby"`
	CreateDate      time.Time      `json:"createdate"`
	IsActive        bool           `json:"isactive"`
	MerchantID      sql.NullString `json:"merchantid"`
	MapSKU          sql.NullString `json:"mapsku"`
	IsFixPrice      bool           `json:"isfixprice"`
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

func TestGetSKUs(t *testing.T) {
	req, err := http.NewRequest("GET", "/skus", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	r := gin.Default()

	r.GET("/skus", getSKUs)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func getSKUs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUs called"})
}
