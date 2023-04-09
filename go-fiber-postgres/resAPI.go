package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewProduct() (*Product, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	collection := client.Database("product").Collection("product")

	return &Product{client, collection}, nil
}

func (e *Product) GetAllProducts() ([]map[string]interface{}, error) {
	filter := bson.M{}
	cur, err := e.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	results := []map[string]interface{}{}

	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (e *Product) SaveProductsToFile(filePath string) error {
	products, err := e.GetAllProducts()
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(products)
	if err != nil {
		return err
	}
	jsonStr := string(jsonBytes)

	err = ioutil.WriteFile(filePath, []byte(jsonStr), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Updated documents")
	return nil
}

func main() {
	product, err := NewProduct()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		products, err := product.GetAllProducts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		jsonBytes, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})

	http.HandleFunc("/product/file", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := product.SaveProductsToFile("product.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("Product saved to file"))
	})

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
