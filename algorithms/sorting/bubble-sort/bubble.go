// Package main provides ...
package bubble

import (
//"fmt"
)

func sort(arr []int) {
	// We minus 1 because we are always comparing the current value with the next value.
	// The number of rounds will be the total length - 1, for array with length 5, we will do 4 rounds: 0 and 1, 1 and 2, 2 and 3, 3 and 4.
	for itemCount := len(arr) - 1; ; itemCount-- {
		swap := false
		// At each round, we compare the current j with the next value
		for i := 1; i <= itemCount; i++ {
			// Only swap their positions if left value < right value as we aim to move all the small values to the back
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}
