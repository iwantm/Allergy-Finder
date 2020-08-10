package functions

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Name of the database.
	DBName = "off"
	URI    = "mongodb://localhost:27017"
)

type Product struct {
	Brand                    string
	ProductName              string
	Barcode                  string
	Ingredients              string
	AllergensFromIngredients string
	Allergens                string
}

func (p Product) Error() string {
	panic("implement me")
}

func SearchProduct(barcode string) *Product {
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	coll := client.Database(DBName).Collection("products")
	fmt.Println(coll.Name())
	var result bson.M

	err = coll.FindOne(ctx, bson.M{"code": barcode}).Decode(&result)

	if err != nil {
		return nil
	}

	product := Product{
		result["brands"].(string),
		result["product_name_en"].(string),
		result["code"].(string),
		result["ingredients_text_en"].(string),
		//add ingredients_hierarchy
		result["allergens_from_ingredients"].(string),
		result["allergens"].(string),
	}

	return &product

}
