package main

import "github.com/BestCharlemagne/StoreNavigator/repository"

//quickSelect places the first numItems smallest items at the beginning of the array.
//They are not necessarily in sorted order.
func quickSelect(zipCode int, array []repository.Store, start int, end int, numItems int) {
	if end-start <= 0 {
		return
	}
	var pivot int = (end - start) / 2
	var pivotItem int = abs(zipCode - (array)[pivot].ZipCode)
	swap(array, 0, pivot) //Pivot now at front

	i := start + 1
	j := end

	for i <= j {
		for i <= j && abs(zipCode-array[i].ZipCode)-pivotItem <= 0 {
			i++
		}
		for i <= j && abs(zipCode-array[j].ZipCode)-pivotItem >= 0 {
			j--
		}
		if i <= j {
			swap(array, i, j)
			i++
			j--
		}
	}
	swap(array, start, j) //Swap pivot into correct place
	if j == numItems-1 {
		return
	} else if j > numItems-1 {
		quickSelect(zipCode, array, start, j-1, numItems)
	}
	quickSelect(zipCode, array, j+1, end, numItems)
}

//swaps the two elements of an array
func swap(array []repository.Store, to int, from int) {
	store := array[to]
	array[to] = array[from]
	array[from] = store
}

//abs value of an integer
func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
