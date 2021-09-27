package app

import "strconv"

func ValidateID(index string) int {

	wantedIndex, _ := strconv.Atoi(index)

	if wantedIndex < 1 || wantedIndex > 151 {
		panic("Please introduce a valid pokemon ID from first gen. (1-151)")
	}

	return wantedIndex
}
