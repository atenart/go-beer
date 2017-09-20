/*
Copyright (C) 2017 Antoine Tenart <antoine.tenart@ack.tf>

This file is licensed under the terms of the GNU General Public License version
2. This program is licensed "as is" without any warranty of any kind, whether
express or implied.
*/

package beer

import "math"

var (
     SrmToEbc	float64 = 1.97
     EbcToSrm	float64 = 0.508
     LToGallon	float64 = 0.264172
     KgToPound	float64 = 2.204620
     KgToOunce	float64 = 35.274
)

// Compute the total boiling time in minutes.
func (r *Recipe) BoilingTime() float64 {
	var bt float64 = 0
	for _, s := range r.Recipe.Mash.MashSteps {
		bt += s.StepTime + s.RampTime
	}
	for _, h := range r.Recipe.Hops {
		if h.Use == "Dry Hop" || h.Use == "Aroma" { continue }
		bt += h.Time
	}
	return bt
}

// Estimate the initial volume needed to reach the targeted volume at the end of
// the brewing.
func (r *Recipe) VolumeStart() float64 {
	var maltAmount float64 = 0
	for _, f := range r.Recipe.Fermentables {
		if f.Type != "Grain" { continue }
		maltAmount += f.Amount
	}

	g := maltAmount * 1	// Percentage loss in the wort. 1l/kg.
	e := r.BoilingTime() / 60 * r.EvapRate()

	return r.Recipe.BatchSize + g + e
}

// Compute the beer color in SRM using the Morey equation.
func (r *Recipe) Color() float64 {
	var mcu float64 = 0
	for _, f := range r.Recipe.Fermentables {
		if f.Type != "Grain" { continue }
		gal := r.Recipe.BatchSize * LToGallon
		mcu += (f.Amount * KgToPound) * (f.Color * SrmToEbc) / gal
	}

	return 1.4922 * math.Pow(mcu, 0.6859) * 1.97
}

// Compute the bitterness using Tinseth formula.
func (r *Recipe) IBU() float64 {
	var tot float64 = 0
	for _, h := range r.Recipe.Hops {
		if h.Use == "Dry Hop" || h.Use == "Aroma" { continue }

		ibu := (h.Alpha / 100) * (h.Amount * KgToOunce) * 7490
		ibu /= r.Recipe.BatchSize * LToGallon
		ibu *= (1 - math.Pow(math.E, -0.04 * h.Time)) / 4.15
		tot += ibu * 1.65 * math.Pow(0.000125, r.OG() - 1)
	}

	return tot
}

// Compute the BU:GU ratio.
func (r *Recipe) BU_GU() float64 {
	return r.IBU() / ((r.OG() - 1)* 1000)
}

// Estimate the original gravity.
func (r *Recipe) EstimatedOG() float64 {
	var og float64 = 0

	for _, f := range r.Recipe.Fermentables {
		gal := r.Recipe.BatchSize * LToGallon
		og += f.Yield * (f.Amount * KgToPound) / gal
	}
	og *= r.Efficiency()

	return (og + 1000) / 1000
}

// Estimate the yeast attenuation.
func (r *Recipe) EstimatedAttenuation() float64 {
	var attenuation, n float64 = 0, 0
	for _, y := range r.Recipe.Yeasts {
		attenuation += y.Attenuation * y.Amount
		n += y.Amount
	}
	return attenuation / n / 100
}

// Estimate the final gravity.
func (r *Recipe) EstimatedFG() float64 {
	return r.OG() - (r.EstimatedAttenuation() * r.OG())
}

// Compute the ideal original gravity.
func (r *Recipe) IdealOG() float64 {
	saved := r.Ideal
	r.Ideal = true
	og := r.EstimatedOG()
	r.Ideal = saved
	return og
}

// Compute the ideal final gravity.
func (r *Recipe) IdealFG() float64 {
	saved := r.Ideal
	r.Ideal = true
	fg := r.EstimatedFG()
	r.Ideal = saved
	return fg
}

// Compute the yeast real attenuation.
func (r *Recipe) Attenuation() float64 {
	if r.Recipe.FG == 0 {
		return r.EstimatedAttenuation()
	}

	return (r.Recipe.OG - r.Recipe.FG) / (r.Recipe.OG)
}

// Compute alcohol by volume.
func (r *Recipe) ABV() float64 {
	return ((1.05 * (r.OG() - r.FG())) / r.FG()) / 0.79 * 100
}
