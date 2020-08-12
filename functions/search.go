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
	Traces                   string
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
	nameEn, nameEnBool := result["product_name_en"].(string)
	name, nameBool := result["product_name"].(string)
	code, codeBool := result["code"].(string)
	ingredientsEn, ingredientsEnBool := result["ingredients_text_en"].(string)
	ingredients, ingredientsBool := result["ingredients_text"].(string)
	allergensFromIngredients, allergensFromIngredientsBool := result["allergens_from_ingredients"].(string)
	allergens, allergensBool := result["allergens"].(string)
	traces, tracesBool := result["traces_from_ingredients"].(string)
	fmt.Println(nameEnBool)

	if brandBool {
		product.Brand = brand
	}
	if nameEnBool {
		product.ProductName = nameEn
	} else if !nameEnBool && nameBool {
		product.ProductName = name
	}
	if codeBool {
		product.Barcode = code
	}
	if ingredientsEnBool {
		product.Ingredients = ingredientsEn
	} else if !ingredientsEnBool && ingredientsBool {
		product.Ingredients = ingredients
	}
	if allergensFromIngredientsBool {
		product.AllergensFromIngredients = allergensFromIngredients
	}
	if allergensBool {
		product.Allergens = allergens
	}
	if tracesBool {
		product.Traces = traces
	}

	return &product
}
