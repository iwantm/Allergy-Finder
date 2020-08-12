package functions

import (
	"github.com/openfoodfacts/openfoodfacts-go"
)

type Product struct {
	Barcode       string
	Brand         string
	ProductName   string
	Ingredients   []string
	Allergens     string
	AllergensTags []string
	Traces        []string
}

func SearchApi(barcode string) (*Product, error) {
	api := openfoodfacts.NewClient("world", "", "")
	item, err := api.Product(barcode)
	if err != nil {
		return nil, err
	}
	var allergenTags []string
	for _, s := range item.AllergensTags {
		allergenTags = append(allergenTags, s.(string))
	}
	product := Product{
		Barcode:       item.Code,
		Brand:         item.Brands,
		ProductName:   item.ProductNameEn,
		Ingredients:   item.IngredientsTags,
		Allergens:     item.Allergens,
		AllergensTags: allergenTags,
		Traces:        item.TracesTags,
	}

	return &product, nil
}
