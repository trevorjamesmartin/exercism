package purchase

import "sort"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	arr := []string{option1, option2}
	s := " is clearly the better choice."
	if sort.StringsAreSorted(arr) {
		return option1 + s
	}
	return option2 + s
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// 10 years and older are worth half their original price
	if age >= 10 {
		return originalPrice * .5
	}
	// between 3 and 10 years are worth 70%
	if age >= 3 {
		return originalPrice * .7
	}
	// 3 years or less are worth 80%
	return originalPrice * .8
}
