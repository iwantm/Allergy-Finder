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
	Brand                    string `bson:"brands,omitempty"`
	ProductName              string `bson:"product_name_en,omitempty"`
	Barcode                  string `bson:"code,omitempty"`
	Ingredients              string `bson:"ingredients_text_en,omitempty"`
	AllergensFromIngredients string `bson:"allergens_from_ingredients,omitempty"`
	Allergens                string `bson:"allergens,omitempty"`
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
	var result bson.M

	err = coll.FindOne(ctx, bson.M{"code": barcode}).Decode(&result)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	product := Product{}
	brand, brandBool := result["brands"].(string)
	name, nameBool := result["product_name_en"].(string)
	code, codeBool := result["code"].(string)
	ingredients, ingredientsBool := result["ingredients_text_en"].(string)
	allergensFromIngredients, allergensFromIngredientsBool := result["allergens_from_ingredients"].(string)
	allergens, allergensBool := result["allergens"].(string)

	switch {
	case brandBool:
		product.Brand = brand
		fallthrough
	case nameBool:
		product.ProductName = name
		fallthrough
	case codeBool:
		product.Barcode = code
		fallthrough
	case ingredientsBool:
		product.Ingredients = ingredients
		fallthrough
	case allergensFromIngredientsBool:
		product.AllergensFromIngredients = allergensFromIngredients
		fallthrough
	case allergensBool:
		product.Allergens = allergens
	}

	return &product
}
