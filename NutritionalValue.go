package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalData struct {
	Energy             float64
	Sugar              float64
	SaturatedFattyAcid float64
	Sodium             float64
	Fiber              float64
	Protein            float64
	IsWater            bool
}

type NutritionalScore struct {
	Value     int
	Posititve int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyAcidsGram float64

type SodiumMilliGram float64

type FiberGram float64

type ProteinGram float64

func GetNutritionalValue()
