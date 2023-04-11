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
	SKUID      sql.NullString  `json:"skuid"`
	MerchantID sql.NullString  `json:"merchantid"`
	BranchID   sql.NullString  `json:"branchid"`
	Price      sql.NullFloat64 `json:"price"`
	StartDate  time.Time       `json:"startdate"`
	EndDate    sql.NullTime    `json:"enddate"`
	IsActive   sql.NullInt32   `json:"isactive"`
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
		rows, err := db.Query("SELECT * FROM backendposdata_sku_branch_price")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var skus_branch []SKU_branch

		for rows.Next() {
			var sku_branch SKU_branch

			err := rows.Scan(&sku_branch.SKUID, &sku_branch.MerchantID, &sku_branch.BranchID, &sku_branch.Price, &sku_branch.StartDate, &sku_branch.EndDate, &sku_branch.IsActive)
			if err != nil {
				log.Fatal(err)
			}
			skus_branch = append(skus_branch, sku_branch)
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, skus_branch)
	})

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
