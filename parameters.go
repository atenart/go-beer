/*
Copyright (C) 2017 Antoine Tenart <antoine.tenart@ack.tf>

This file is licensed under the terms of the GNU General Public License version
2. This program is licensed "as is" without any warranty of any kind, whether
express or implied.
*/

package beer

import "math"

var (
     MlToGallon	float64 = 0.000264172
     GrToPound	float64 = 0.002204620
     GrToOunce	float64 = 0.035274000
)

// Compute the total boiling time in seconds.
func (r *Recipe) BoilingTime() float64 {
	var tot float64 = 0
	for _, s := range r.Steps {
		tot += s.Time
	}
	for _, h := range r.Hops {
		tot += h.Time
	}
	return tot
}

// Estimate the initial volume needed to reach the targeted volume at the end of
// the brewing.
func (r *Recipe) VolumeStart() float64 {
	var maltAmount float64 = 0
	for _, m := range r.Malts {
		maltAmount += m.Amount
	}

	g := maltAmount * r.Settings.GrainLoss
	e := r.BoilingTime() / 360. * (r.Settings.EvaporationLoss / 100.) / 100.
	return r.Volume + g + (e - 1.) * r.Volume
}

// Compute the beer color in SRM.
func (r *Recipe) Color() float64 {
	var mcu float64 = 0
	for _, m := range r.Malts {
		mcu += (m.Amount * m.EBC * GrToPound) / (r.Volume * MlToGallon)
	}

	// Morey equation
	return 1.4922 * math.Pow(mcu, 0.6859) * 1.97
}

// Compute the bitterness using Tinseth formula.
func (r *Recipe) IBU() float64 {
	var tot float64 = 0
	for _, h := range r.Hops {
		if h.DryHop { continue }

		ibu := (h.Alpha / 100.) * (h.Amount * GrToOunce) * 7490.
		ibu /= r.Volume * MlToGallon
		ibu *= (1 - math.Pow(math.E, -0.04 * h.Time / 60)) / 4.15
		tot += ibu * 1.65 * math.Pow(0.000125, r.OG() / 1000. - 1.)
	}

	return tot
}

// Compute the BU:GU ratio.
func (r *Recipe) BU_GU() float64 {
	return r.IBU() / (r.OG() - 1000)
}

// Estimate the original gravity.
func (r *Recipe) EstimatedOG() float64 {
	var og float64 = 0

	for _, m := range r.Malts {
		og += m.PPG * m.Amount * GrToPound / (r.Volume * MlToGallon)
	}
	og *= r.Settings.Efficiency / 100.

	return og + 1000
}

// Estimate the yeast attenuation.
func (r *Recipe) EstimatedAttenuation() float64 {
	var attenuation, n float64 = 0, 0
	for _, y := range r.Yeasts {
		attenuation += y.Attenuation * y.Amount
		n += y.Amount
	}
	return attenuation / n
}

// Estimate the final gravity.
func (r *Recipe) EstimatedFG() float64 {
	return r.OG() - (r.EstimatedAttenuation() * r.OG()) / 100.
}

// Compute the ideal original gravity.
func (r *Recipe) IdealOG() float64 {
	savedEfficiency := r.Settings.Efficiency
	r.Settings.Efficiency = 100
	og := r.EstimatedOG()
	r.Settings.Efficiency = savedEfficiency
	return og
}

// Compute the ideal final gravity.
func (r *Recipe) IdealFG() float64 {
	savedEfficiency := r.Settings.Efficiency
	r.Settings.Efficiency = 100
	fg := r.EstimatedFG()
	r.Settings.Efficiency = savedEfficiency
	return fg
}

// Compute the brewing efficiency.
func (r *Recipe) Efficiency() float64 {
	if r.MeasuredOG == -1 {
		return r.Settings.Efficiency
	}

	return (r.MeasuredOG - 1000) / (r.IdealOG() - 1000) * 100
}

// Compute the yeast real attenuation.
func (r *Recipe) Attenuation() float64 {
	if r.MeasuredFG == -1 {
		return r.EstimatedAttenuation()
	}

	return (r.MeasuredOG - r.MeasuredFG) / (r.MeasuredOG) * 100
}

// Compute alcohol by volume.
func (r *Recipe) ABV() float64 {
	return ((1.05 * (r.OG() - r.FG())) / r.FG()) / 0.79 * 100
}
