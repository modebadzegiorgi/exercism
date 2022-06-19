package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}

	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, ingredient := range layers {
		if ingredient == "noodles" {
			noodles += 50
		} else if ingredient == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, scale int) []float64 {
	returnAMounts := make([]float64, len(amounts))

	for i, am := range amounts {
		returnAMounts[i] = am / 2.0 * float64(scale)
	}

	return returnAMounts
}
