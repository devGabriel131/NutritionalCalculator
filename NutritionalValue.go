package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalData struct {
	Energy             EnergyKJ
	Sugar              SugarGram
	SaturatedFattyAcid SaturatedFattyAcidsGram
	Sodium             SodiumMilliGram
	Fiber              FiberGram
	Fruit              FruitGram
	Protein            ProteinGram
	IsWater            bool
}

type NutritionalScore struct {
	Value     int
	Posititve int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64

func (e EnergyKJ) GetPoints(st ScoreType) int {
	return
}

type FruitGram float64

type SugarGram float64

type SaturatedFattyAcidsGram float64

type SodiumMilliGram float64

type FiberGram float64

type ProteinGram float64

func GetNutritionalValue(n NutritionalData, st ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	if st != Water {
		// bad most of the time
		energyPoints := n.Energy.GetPoints(st)
		sugarPoints := n.Sugar.GetPoints(st)
		satFatPoints := n.SaturatedFattyAcidsGram.GetPoints(st)
		sodiumPoints := n.Sodium.GetPoints(st)

		// good most of the time
		fiberPoints := n.Fiber.GetPoints(st)
		fruitsPoints := n.Fruit.GetPoints(st)
		proteinPoints := n.Fiber.GetPoints(st)

		positive += fiberPoints + fruitsPoints + proteinPoints
		negative += energyPoints + sugarPoints + satFatPoints + sodiumPoints
	}

	return NutritionalScore{
		Value:     value,
		Posititve: positive,
		Negative:  negative,
		ScoreType: st,
	}
}
