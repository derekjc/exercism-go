package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * timePerLayer
	}
}

func Quantities(layers []string) (int, float64) {
	noodleQty := 0
	sauceQty := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleQty += 50
		}
		if layer == "sauce" {
			sauceQty += 0.2
		}
	}
	return noodleQty, sauceQty
}

func AddSecretIngredient(friendsList, myList []string) {
	lastFriendsItem := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = lastFriendsItem
}

func ScaleRecipe(quantities []float64, scale int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i, q := range quantities {
		scaledQuantities[i] = q * float64(scale) / 2.0
	}
	return scaledQuantities
}
