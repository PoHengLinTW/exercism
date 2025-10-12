package lasagna

func PreparationTime(layers []string, minutePerLayer int) int {
    if minutePerLayer == 0 {
        minutePerLayer = 2
    }
    return len(layers) * minutePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
    noodles = 0
    sauce = 0.0

    for _, layer := range(layers) {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }
    
    return
}

func AddSecretIngredient(friendRecipe, ownRecipe []string) {
    ownRecipe[len(ownRecipe) - 1] = friendRecipe[len(friendRecipe) - 1]
}

func ScaleRecipe(amountForTwoPortion []float64, expectedPortion int) []float64 {
    var amountNeeded []float64
	for _, amount := range(amountForTwoPortion) {
        amountNeeded = append(amountNeeded, amount / 2.0 * float64(expectedPortion))
    }
    return amountNeeded
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
