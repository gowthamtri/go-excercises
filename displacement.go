package main

import (
	"fmt"
)

func main() {
	var initialAcceleration float64;
	var initialVelocity float64;
	var initialDisplacement float64;

	var inputTime float64;

	fmt.Printf("Enter Initial Acceleration : ");
	fmt.Scan(&initialAcceleration);

	fmt.Printf("Enter Initial Velocity : ");
	fmt.Scan(&initialVelocity);

	fmt.Printf("Enter Initial Displacement : ");
	fmt.Scan(&initialDisplacement);

	fmt.Printf("Acceleration : %f Velocity : %f Initial displacement : %f", initialAcceleration, initialVelocity, initialDisplacement);

	calculateDisplacementFn := GetDisplacementFn(initialAcceleration, initialVelocity, initialDisplacement);

	for {
		fmt.Printf("\nEnter the time to calculate displacement : ");
		num, err := fmt.Scan(&inputTime);
		if num > 0 && err == nil {
			fmt.Printf("Calculated Displacement : %f", calculateDisplacementFn(inputTime));
		} else {
			break;
		}
	}
}

func GetDisplacementFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v * t) + s;
	}
}