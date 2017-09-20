/*
Copyright (C) 2017 Antoine Tenart <antoine.tenart@ack.tf>

This file is licensed under the terms of the GNU General Public License version
2. This program is licensed "as is" without any warranty of any kind, whether
express or implied.
*/

/*
The package beer helps making beer recipes and provides functions to compute
various parameters given a BeerXML formated recipe.

Units (follow the Beer XML specification):
- Weights in kilograms (kg).
- Volumes in liters (l).
- Temperatures in degree Celsius (°C).
- Times in minutes (m).
- Pressures in kilopascals (kPa).
*/
package beer

import "github.com/atenart/go-beer/beerxml"

type Recipe struct {
	Recipe		*beerxml.Recipe
	Ideal		bool
}

// Start a new beer recipe.
func NewRecipe(r *beerxml.Recipe) *Recipe {
	return &Recipe{
		Recipe:	r,
		Ideal:	false,
	}
}

// Return the original gravity.
func (r *Recipe) OG() float64 {
	if r.Recipe.OG != 0 { return r.Recipe.OG }
	return r.EstimatedOG()
}

// Return the final gravity.
func (r *Recipe) FG() float64 {
	if r.Recipe.FG != 0 { return r.Recipe.FG }
	return r.EstimatedFG()
}

// Return the brewing efficiency.
func (r *Recipe) Efficiency() float64 {
	if r.Ideal { return 1 }

	if r.Recipe.OG == 0 {
		if r.Recipe.Efficiency != 0 { return r.Recipe.Efficiency / 100 }
		return 0.7	// 70% is a safe default
	}

	return (r.Recipe.OG - 1) * 1000 / ((r.IdealOG() - 1) * 1000)
}

// Return the evaporation rate.
func (r *Recipe) EvapRate() float64 {
	if r.Recipe.Equipment.EvapRate != 0 {
		return r.Recipe.Equipment.EvapRate / 100
	}
	return 1	// 1l/h
}
