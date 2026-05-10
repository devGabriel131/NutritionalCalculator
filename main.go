package main

import "fmt"

func main() {
	ns := GetNutritionalValue(NutritionalData{
		Energy:              EnergyFromKcal(),
		Sugars:              SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMilliGram(),
		Fiber:               FiberGram(),
		Fruits:              FruitsPercent(),
		Proteins:            ProteinGram(),
	}, Food)

	fmt.Printf("You Nutritional Score is: %d", ns.Value)
}
