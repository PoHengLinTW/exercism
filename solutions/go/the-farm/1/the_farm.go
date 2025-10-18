package thefarm

import (
	"errors"
    "fmt"
)

func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
    fodderAmount, err := fc.FodderAmount(numberOfCows)
    if err != nil {
        return 0.0, err
    }
    
    fatteningFactor, err := fc.FatteningFactor()
    if err != nil {
        return 0.0, err
    }

    return fodderAmount * fatteningFactor / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
    if numberOfCows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    }
    
    return DivideFood(fc, numberOfCows)
}

type InvalidCowsError struct {
    numberOfCows int
    message string
}

func (i *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", i.numberOfCows, i.message)
}

func ValidateNumberOfCows(numberOfCows int) error {
    switch {
        case numberOfCows < 0: 
        	return &InvalidCowsError{ numberOfCows: numberOfCows, message: "there are no negative cows"}
        case numberOfCows == 0:
        	return &InvalidCowsError{ numberOfCows: numberOfCows, message: "no cows don't need food"}
        default:
        	return nil
    }
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
