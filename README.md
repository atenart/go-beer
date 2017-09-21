# Go-Beer

Go-Beer provides two GO packages: *beer* and *beerxml*. The former provides
functions to compute various beer recipe parameters given a recipe expressed in
the Beer XML format; and the later provides a description of the Beer XML format
with helpers to export and import XML files.

## Code example

```go
package main

import (
	"fmt"
	"os"
	"github.com/atenart/go-beer"
	"github.com/atenart/go-beer/beerxml"
)

func main() {
	data, err := beerxml.Import("./example-recipe.xml")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(0)
	}

	r := beer.NewRecipe(&data.Recipe[0])

	// Get or estimate the recipe parameters.
	fmt.Printf("Starting volume: %.1f liters\n", r.VolumeStart())
	fmt.Printf("Final volume: %.1f liters\n", r.Recipe.BatchSize)
	fmt.Printf("Color: %.0f EBC\n", r.Color())
	fmt.Printf("IBU: %.0f\n", r.IBU())
	fmt.Printf("Estimated attenuation: %.1f\n", r.Attenuation() * 100)
	fmt.Printf("Estimated OG: %.3f\n", r.OG())
	fmt.Printf("Estimated FG: %.3f\n", r.FG())
	fmt.Printf("Estimated ABV: %.1f°\n", r.ABV())
	fmt.Printf("Estimated BU:GU: %.2f\n", r.BU_GU())
	fmt.Printf("\n")

	// Set measured values.
	r.Recipe.OG = 1.056
	r.Recipe.FG = 1.010

	// Compute the real parameters.
	fmt.Printf("ABV: %.1f°\n", r.ABV())
	fmt.Printf("Efficiency: %.1f\n", r.Efficiency() * 100)
	fmt.Printf("Attenuation %.1f\n", r.Attenuation() * 100)

	beerxml.Export(&beerxml.BeerXML{
		Recipes: []beerxml.Recipe{ *r.Recipe, },
	}, "./example-batch.xml")
}
```
