/*
Copyright (C) 2017 Antoine Tenart <antoine.tenart@ack.tf>

This file is licensed under the terms of the GNU General Public License version
2. This program is licensed "as is" without any warranty of any kind, whether
express or implied.
*/

// The package beerxml implements the BeerXML standard used to exchange and
// represent brewing data. Do not support extensions.
// See http://www.beerxml.com.
package beerxml

type BeerXML struct {
	Hops         []Hop         `xml:"HOPS>HOP"`
	Fermentables []Fermentable `xml:"FERMENTABLES>FERMENTABLE"`
	Yeasts       []Yeast       `xml:"YEASTS>YEAST"`
	Miscs        []Misc        `xml:"MISCS>MISC"`
	Waters       []Water       `xml:"WATERS>WATER"`
	Equipments   []Equipment   `xml:"EQUIPMENTS>EQUIPMENT"`
	Styles       []Style       `xml:"STYLES>STYLE"`
	MashSteps    []MashStep    `xml:"MASH_STEPS>MASH_STEP"`
	Mashs        []Mash        `xml:"MASHS>MASH"`
	Recipes      []Recipe      `xml:"RECIPES>RECIPE"`
}

type Hop struct {
	Name          string  `xml:"NAME"`
	Version       int32   `xml:"VERSION"`
	Alpha         float64 `xml:"ALPHA"`
	Amount        float64 `xml:"AMOUNT"`
	Use           string  `xml:"USE"`
	Time          float64 `xml:"TIME"`
	Notes         string  `xml:"NOTES"`
	Type          string  `xml:"TYPE"`
	Form          string  `xml:"FORM"`
	Beta          float64 `xml:"BETA"`
	Hsi           float64 `xml:"HSI"`
	Origin        string  `xml:"ORIGIN"`
	Substitutes   string  `xml:"SUBSTITUTES"`
	Humulene      float64 `xml:"HUMULENES"`
	Caryophyllene float64 `xml:"CARYOPHYLLENE"`
	Cohumulone    float64 `xml:"COHUMULONE"`
	Myrcene       float64 `xml:"MYRCENE"`
}

type Fermentable struct {
	Name           string  `xml:"NAME"`
	Version        int32   `xml:"VERSION"`
	Type           string  `xml:"TYPE"`
	Amount         float64 `xml:"AMOUNT"`
	Yield          float64 `xml:"YIELD"`
	Color          float64 `xml:"COLOR"`
	AddAfterBoil   bool    `xml:"ADD_AFTER_BOIL"`
	Origin         string  `xml:"ORIGIN"`
	Supplier       string  `xml:"SUPPLIER"`
	Notes          string  `xml:"NOTES"`
	CoarseFineDiff float64 `xml:"COARSE_FINE_DIFF"`
	Moisture       float64 `xml:"MOISTURE"`
	DiastaticPower float64 `xml:"DIASTATIC_POWER"`
	Protein        float64 `xml:"PROTEIN"`
	MaxInBatch     float64 `xml:"MAX_IN_BATCH"`
	RecommendMash  bool    `xml:"RECOMMEND_MASH"`
	IbuGalPerLb    float64 `xml:"IBU_GAL_PER_LB"`
}

type Yeast struct {
	Name           string  `xml:"NAME"`
	Version        int32   `xml:"VERSION"`
	Type           string  `xml:"TYPE"`
	Form           string  `xml:"FORM"`
	Amount         float64 `xml:"AMOUNT"`
	AmountIsWeight bool    `xml:"AMOUNT_IS_WEIGHT"`
	Laboratory     string  `xml:"LABORATORY"`
	ProductId      string  `xml:"PRODUCT_ID"`
	MinTemperature float64 `xml:"MIN_TEMPERATURE"`
	MaxTemperature float64 `xml:"MAX_TEMPERATURE"`
	Flocculation   string  `xml:"FLOCCULATION"`
	Attenuation    float64 `xml:"ATTENUATION"`
	Notes          string  `xml:"NOTES"`
	BestFor        string  `xml:"BEST_FOR"`
	TimesCultured  int32   `xml:"TIMES_CULTURED"`
	MaxReuse       int32   `xml:"MAX_REUSE"`
	AddToSecondary bool    `xml:"ADD_TO_SECONDARY"`
}

type Misc struct {
	Name           string  `xml:"NAME"`
	Version        int32   `xml:"VERSION"`
	Type           string  `xml:"TYPE"`
	Use            string  `xml:"USE"`
	Amount         float64 `xml:"AMOUNT"`
	AmountIsWeight bool    `xml:"AMOUNT_IS_WEIGHT"`
	UseFor         string  `xml:"USE_FOR"`
	Notes          string  `xml:"NOTES"`
}

type Water struct {
	Name        string  `xml:"NAME"`
	Version     int32   `xml:"VERSION"`
	Amount      float64 `xml:"AMOUNT"`
	Calcium     float64 `xml:"CALCIUM"`
	Bicarbonate float64 `xml:"BICARBONATE"`
	Sulfate     float64 `xml:"SULFATE"`
	Chloride    float64 `xml:"CHLORIDE"`
	Sodium      float64 `xml:"SODIUM"`
	Magnesium   float64 `xml:"MAGNESIUM"`
	Ph          float64 `xml:"PH"`
	Notes       string  `xml:"NOTES"`
}

type Equipment struct {
	Name            string  `xml:"NAME"`
	Version         int32   `xml:"VERSION"`
	BoilSize        float64 `xml:"BOIL_SIZE"`
	BatchSize       float64 `xml:"BATCH_SIZE"`
	TunVolume       float64 `xml:"TUN_VOLUME"`
	TunWeight       float64 `xml:"TUN_WEIGHT"`
	TunSpecificHeat float64 `xml:"TUN_SPECIFIC_HEAT"`
	TopUpWater      float64 `xml:"TOP_UP_WATER"`
	TrubChillerLoss float64 `xml:"TRUB_CHILLER_LOSS"`
	EvapRate        float64 `xml:"EVAP_RATE"`
	BoilTime        float64 `xml:"BOIL_TIME"`
	CalcBoilVolume  bool    `xml:"CALC_BOIL_VOLUME"`
	LauterDeadspace float64 `xml:"LAUTER_DEADSPACE"`
	TopUpKettle     float64 `xml:"TOP_UP_KETTLE"`
	HopUtilization  float64 `xml:"HOP_UTILIZATION"`
	Notes           string  `xml:"NOTES"`
}

type Style struct {
	Name           string  `xml:"NAME"`
	Category       string  `xml:"CATEGORY"`
	Version        int32   `xml:"VERSION"`
	CategoryNumber string  `xml:"CATEGORY_NUMBER"`
	StyleLetter    string  `xml:"STYLE_LETTER"`
	StyleGuide     string  `xml:"STYLE_GUIDE"`
	Type           string  `xml:"TYPE"`
	OgMin          float64 `xml:"OG_MIN"`
	OgMax          float64 `xml:"OG_MAX"`
	FgMin          float64 `xml:"FG_MIN"`
	FgMax          float64 `xml:"FG_MAX`
	IbuMin         float64 `xml:"IBU_MIN"`
	IbuMax         float64 `xml:"IBU_MAX"`
	ColorMin       float64 `xml:"COLOR_MIN"`
	ColorMax       float64 `xml:"COLOR_MAX"`
	CarbMin        float64 `xml:"CARB_MIN"`
	CarbMax        float64 `xml:"CARB_MAX"`
	AbvMax         float64 `xml:"ABV_MAX"`
	AbvMin         float64 `xml:"ABV_MIN"`
	Notes          string  `xml:"NOTES"`
	Profile        string  `xml:"PROFILE"`
	Ingredients    string  `xml:"INGREDIENTS"`
	Examples       string  `xml:"EXAMPLES"`
}

type MashStep struct {
	Name         string  `xml:"NAME"`
	Version      int32   `xml:"VERSION"`
	Type         string  `xml:"TYPE"`
	InfuseAmount float64 `xml:"INFUSE_AMOUNT"`
	StepTemp     float64 `xml:"STEP_TEMP"`
	StepTime     float64 `xml:"STEP_TIME"`
	RampTime     float64 `xml:"RAMP_TIME"`
	EndTemp      float64 `xml:"END_TEMP"`
}

type Mash struct {
	Name            string     `xml:"NAME"`
	Version         int32      `xml:"VERSION"`
	GrainTemp       float64    `xml:"GRAIN_TEMP"`
	MashSteps       []MashStep `xml:"MASH_STEPS>MASH_STEP"`
	Notes           string     `xml:"NOTES"`
	TunTemp         float64    `xml:"TUN_TEMP"`
	SpargeTemp      float64    `xml:"SPARGE_TEMP"`
	Ph              float64    `xml:"PH"`
	TunWeight       float64    `xml:"TUN_WEIGHT"`
	TunSpecificHeat float64    `xml:"TUN_SPECIFIC_HEAT"`
	EquipAdjust     bool       `xml:"EQUIP_ADJUST"`
}

type Recipe struct {
	Name               string        `xml:"NAME"`
	Version            int32         `xml:"VERSION"`
	Type               string        `xml:"TYPE"`
	Style              Style         `xml:"STYLE"`
	Equipment          Equipment     `xml:"EQUIPMENT"`
	Brewer             string        `xml:"BREWER"`
	AsstBrewer         string        `xml:"ASST_BREWER"`
	BatchSize          float64       `xml:"BATCH_SIZE"`
	BoilSize           float64       `xml:"BOIL_SIZE"`
	BoilTime           float64       `xml:"BOIL_TIME"`
	Efficiency         float64       `xml:"EFFICIENCY"`
	Hops               []Hop         `xml:"HOPS>HOP"`
	Fermentables       []Fermentable `xml:"FERMENTABLES>FERMENTABLE"`
	Miscs              []Misc        `xml:"MISCS>MISC"`
	Yeasts             []Yeast       `xml:"YEASTS>YEAST"`
	Waters             []Water       `xml:"WATERS>WATER"`
	Mash               Mash          `xml:"MASH"`
	Notes              string        `xml:"NOTES"`
	TasteNotes         string        `xml:"TASTE_NOTES"`
	TasteRating        float64       `xml:"TASTE_RATING"`
	OG                 float64       `xml:"OG"`
	FG                 float64       `xml:"FG"`
	FermentationStages int32         `xml:"FERMENTATION_STAGES"`
	PrimaryAge         float64       `xml:"PRIMARY_AGE"`
	PrimaryTemp        float64       `xml:"PRIMARY_TEMP"`
	SecondaryAge       float64       `xml:"SECONDARY_AGE"`
	SecondaryTemp      float64       `xml:"SECONDARY_TEMP"`
	TertiaryAge        float64       `xml:"TERTIARY_AGE"`
	TertiaryTemp       float64       `xml:"TERTIARY_TEMP"`
	Age                float64       `xml:"AGE"`
	AgeTemp            float64       `xml:"AGE_TEMP"`
	Date               string        `xml:"DATE"`
	Carbonation        float64       `xml:"CARBONATION"`
	ForcedCarbonation  bool          `xml:"FORCED_CARBONATION"`
	PrimingSugarName   string        `xml:"PRIMING_SUGAR_NAME"`
	CarbonationTemp    float64       `xml:"CARBONATIOn_TEMP"`
	PrimingSugarEquiv  float64       `xml:"PRIMING_SUGAR_EQUIV"`
	KegPrimingFactor   float64       `xml:"KEG_PRIMING_FACTOR"`
}
