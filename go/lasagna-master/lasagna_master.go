package lasagna

func PreparationTime(layers []string, averageLayerTime int) int {
	if averageLayerTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * averageLayerTime
}

func Quantities(layers []string) (int, float64) {
	n := 0
	s := 0.0
	for step := 0; step < len(layers); step++ {
		switch layers[step] {
		case "sauce":
			s += 0.2 // liters
		case "noodles":
			n += 50 // grams
		default:
			// not applicable
		}
	}
	return n, s
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(recipe []float64, portions int) []float64 {
	scaledRecipe := make([]float64, len(recipe))
	scaler := float64(portions)
	for item := 0; item < len(recipe); item++ {
		value := scaler * recipe[item] / 2
		scaledRecipe[item] = value
	}
	return scaledRecipe
}
