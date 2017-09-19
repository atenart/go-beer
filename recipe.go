/*
Copyright (C) 2017 Antoine Tenart <antoine.tenart@ack.tf>

This file is licensed under the terms of the GNU General Public License version
2. This program is licensed "as is" without any warranty of any kind, whether
express or implied.
*/

// The package beer helps making beer recipes and provides functions to compute
// various parameters. Amounts are in grams, volumes in liter, yield and
// attenuations in percentages and times in seconds.
package beer

type Recipe struct {
	Name		string
	Description	string
	Volume		float64
	Malts		[]*Malt
	Hops		[]*Hop
	Yeasts		[]*Yeast
	Steps		[]*Step
	MeasuredOG	float64
	MeasuredFG	float64
	Settings	Settings
}

type Malt struct {
	Name		string
	EBC		float64
	PPG		float64
	Amount		float64
}

type Hop struct {
	Name		string
	Alpha		float64
	Amount		float64
	Time		float64
	DryHop		bool
}

type Yeast struct {
	Name		string
	Attenuation	float64
	Amount		float64
}

type Step struct {
	Temperature	float64
	Time		float64
}

type Settings struct {
	Efficiency	float64
	EvaporationLoss	float64
	GrainLoss	float64
}

// Start a new recipe.
func NewRecipe(name, description string) *Recipe {
	return &Recipe{
		Name:		name,
		Description:	description,
		Settings:	Settings{70, 8, 5},
		MeasuredOG:	-1,
		MeasuredFG:	-1,
	}
}

// Set the equipment settings.
func (r *Recipe) SetSettings(efficiency, evaporationLoss, grainLoss float64) {
	r.Settings.Efficiency = efficiency
	r.Settings.EvaporationLoss = evaporationLoss
	r.Settings.GrainLoss = grainLoss
}

// Set the expected recipe volume at the end of the brewing.
func (r *Recipe) SetVolume(volume float64) {
	r.Volume = volume
}

// Add a cooking step.
func (r *Recipe) AddStep(temperature, time float64) {
	r.Steps = append(r.Steps, &Step{
		Temperature:	temperature,
		Time:		time,
	})
}

// Add a malt to a recipe.
func (r *Recipe) AddMalt(name string, ebc, ppg, amount float64) {
	r.Malts = append(r.Malts, &Malt{
		Name:	name,
		EBC:	ebc,
		PPG:	ppg,
		Amount:	amount,
	})
}

// Add an hop to a recipe.
func (r *Recipe) AddHop(name string, alpha, amount, time float64, dryHop bool) {
	r.Hops = append(r.Hops, &Hop{
		Name:	name,
		Alpha:	alpha,
		Amount:	amount,
		Time:	time,
		DryHop:	dryHop,
	})
}

// Add a yeast to a recipe.
func (r *Recipe) AddYeast(name string, attenuation, amount float64) {
	r.Yeasts = append(r.Yeasts, &Yeast{
		Name:		name,
		Attenuation:	attenuation,
		Amount:		amount,
	})
}

// Set the measured original gravity
func (r *Recipe) SetOG(OG float64) {
	r.MeasuredOG = OG
}

// Set the measured final gravity
func (r *Recipe) SetFG(FG float64) {
	r.MeasuredFG = FG
}

// Get the original gravity.
func (r *Recipe) OG() float64 {
	if r.MeasuredOG != -1 { return r.MeasuredOG }
	return r.EstimatedOG()
}

// Get the final gravity.
func (r *Recipe) FG() float64 {
	if r.MeasuredFG != -1 { return r.MeasuredFG }
	return r.EstimatedFG()
}
