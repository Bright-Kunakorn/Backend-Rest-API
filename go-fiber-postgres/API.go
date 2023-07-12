package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type SKU_branch struct {
	
	SKUID      *string    `json:"skuid"`
	MerchantID *string    `json:"merchantid"`
	BranchID   *string    `json:"branchid"`
	Price      *float64   `json:"price"`
	StartDate  *time.Time `json:"startdate"`
	EndDate    *time.Time `json:"enddate"`
	IsActive   *int32     `json:"isactive"`
}

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
	r.GET("/skus_branch", func(c *gin.Context) {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM backendposdata_sku_branch_price").Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		rows, err := db.Query("SELECT * FROM backendposdata_sku_branch_price")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var skus_branch []SKU_branch
		for rows.Next() {
			var sku_branch SKU_branch

			err := rows.Scan(
				&sku_branch.SKUID,
				&sku_branch.MerchantID,
				&sku_branch.BranchID,
				&sku_branch.Price,
				&sku_branch.StartDate,
				&sku_branch.EndDate,
				&sku_branch.IsActive)
			if err != nil {
				log.Fatal(err)
			}
			skus_branch = append(skus_branch, sku_branch)
		}
		log.Println(len(skus_branch) == count)
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, skus_branch)
	})
	r.GET("/skus", func(c *gin.Context) {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM backendposdatasku").Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		rows, err := db.Query("SELECT * FROM backendposdatasku")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var skus []SKU
		for rows.Next() {
			var sku SKU
			err := rows.Scan(
				&sku.SKUID,
				&sku.BarcodePOS,
				&sku.ProductName,
				&sku.BrandID,
				&sku.ProductGroupID,
				&sku.ProductCatID,
				&sku.ProductSubCatID,
				&sku.ProductSizeID,
				&sku.ProductUnit,
				&sku.PackSize,
				&sku.Unit,
				&sku.BanForPracharat,
				&sku.IsVat,
				&sku.CreateBy,
				&sku.CreateDate,
				&sku.IsActive,
				&sku.MerchantID,
				&sku.MapSKU,
				&sku.IsFixPrice)
			if err != nil {
				log.Fatal(err)
			}
			skus = append(skus, sku)
		}
		log.Println(len(skus) == count)
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, skus)
	})

	r.GET("/skus/:id", func(c *gin.Context) {
		id := c.Param("id")
		rows, err := db.Query("SELECT * FROM backendposdatasku WHERE skuid='" + id + "'")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var skus []SKU
		for rows.Next() {
			var sku SKU

			err := rows.Scan(
				&sku.SKUID,
				&sku.BarcodePOS,
				&sku.ProductName,
				&sku.BrandID,
				&sku.ProductGroupID,
				&sku.ProductCatID,
				&sku.ProductSubCatID,
				&sku.ProductSizeID,
				&sku.ProductUnit,
				&sku.PackSize,
				&sku.Unit,
				&sku.BanForPracharat,
				&sku.IsVat,
				&sku.CreateBy,
				&sku.CreateDate,
				&sku.IsActive,
				&sku.MerchantID,
				&sku.MapSKU,
				&sku.IsFixPrice)
			if err != nil {
				log.Fatal(err)
			}
			skus = append(skus, sku)
		}
		c.JSON(http.StatusOK, skus)
	})
	r.GET("/skus_branch/:id", func(c *gin.Context) {
		id := c.Param("id")
		rows, err := db.Query("SELECT * FROM backendposdata_sku_branch_price WHERE skuid='" + id + "'")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var skus_branch []SKU_branch
		for rows.Next() {
			var sku_branch SKU_branch

			err := rows.Scan(
				&sku_branch.SKUID,
				&sku_branch.MerchantID,
				&sku_branch.BranchID,
				&sku_branch.Price,
				&sku_branch.StartDate,
				&sku_branch.EndDate,
				&sku_branch.IsActive)
			if err != nil {
				log.Fatal(err)
			}
			skus_branch = append(skus_branch, sku_branch)
		}
		c.JSON(http.StatusOK, skus_branch)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
