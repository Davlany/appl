package main

import (
	"math/rand"
	"sort"
	"time"
)

type QuickSortStats struct {
	Comparisons   int
	Swaps         int
	MemoryUsage   int
	ExecutionTime time.Duration
}

// Partition schemes and pivot selection methods
func lomutoPartition(arr []int, low, high int) (int, *QuickSortStats) {
	stats := &QuickSortStats{}
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		stats.Comparisons++
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			stats.Swaps++
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	stats.Swaps++
	return i, stats
}

func hoarePartition(arr []int, low, high int) (int, *QuickSortStats) {
	stats := &QuickSortStats{}
	pivot := arr[low]
	i, j := low-1, high+1
	for {
		for {
			stats.Comparisons++
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			stats.Comparisons++
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j, stats
		}
		arr[i], arr[j] = arr[j], arr[i]
		stats.Swaps++
	}
}

func threeWayPartition(arr []int, low, high int) (int, int, *QuickSortStats) {
	stats := &QuickSortStats{}
	pivot := arr[low]
	lt, i, gt := low, low+1, high
	for i <= gt {
		stats.Comparisons++
		if arr[i] < pivot {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
			stats.Swaps++
		} else if arr[i] > pivot {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
			stats.Swaps++
		} else {
			i++
		}
	}
	return lt, gt, stats
}

func dualPivotPartition(arr []int, low, high int) (int, int, *QuickSortStats) {
	stats := &QuickSortStats{}
	if low >= high {
		return low, high, stats
	}

	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
		stats.Swaps++
	}
	pivot1, pivot2 := arr[low], arr[high]
	lt, gt := low+1, high-1
	i := lt

	for i <= gt {
		stats.Comparisons++
		if arr[i] < pivot1 {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
			stats.Swaps++
		} else if arr[i] > pivot2 {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
			stats.Swaps++
		} else {
			i++
		}
	}

	lt--
	gt++
	arr[low], arr[lt] = arr[lt], arr[low]
	arr[high], arr[gt] = arr[gt], arr[high]
	stats.Swaps += 2
	return lt, gt, stats
}

// Pivot selection methods
func choosePivotLast(arr []int, low, high int) int {
	return high
}

func choosePivotRandom(arr []int, low, high int) int {
	randomIndex := rand.Intn(high-low+1) + low
	return randomIndex
}

func choosePivotMedianOfThree(arr []int, low, high int) int {
	mid := (low + high) / 2
	indices := []int{low, mid, high}
	sort.Slice(indices, func(i, j int) bool { return arr[indices[i]] < arr[indices[j]] })
	return indices[1]
}

func choosePivotMedianOfThreeRandom(arr []int, low, high int) int {
	indices := []int{
		rand.Intn(high-low+1) + low,
		rand.Intn(high-low+1) + low,
		rand.Intn(high-low+1) + low,
	}
	sort.Slice(indices, func(i, j int) bool { return arr[indices[i]] < arr[indices[j]] })
	return indices[1]
}

func quickSort(arr []int, low, high int, partition func([]int, int, int) (int, *QuickSortStats), pivotSelector func([]int, int, int) int) *QuickSortStats {
	stats := &QuickSortStats{}
	if low < high {
		pivotIndex := pivotSelector(arr, low, high)
		arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
		partitionIndex, partStats := partition(arr, low, high)

		stats.Comparisons += partStats.Comparisons
		stats.Swaps += partStats.Swaps

		leftStats := quickSort(arr, low, partitionIndex-1, partition, pivotSelector)
		rightStats := quickSort(arr, partitionIndex+1, high, partition, pivotSelector)

		stats.Comparisons += leftStats.Comparisons + rightStats.Comparisons
		stats.Swaps += leftStats.Swaps + rightStats.Swaps
	}
	return stats
}

func quickSortThreeWay(arr []int, low, high int, pivotSelector func([]int, int, int) int) *QuickSortStats {
	stats := &QuickSortStats{}
	if low < high {
		pivotIndex := pivotSelector(arr, low, high)
		arr[pivotIndex], arr[low] = arr[low], arr[pivotIndex]
		lt, gt, partStats := threeWayPartition(arr, low, high)

		stats.Comparisons += partStats.Comparisons
		stats.Swaps += partStats.Swaps

		leftStats := quickSortThreeWay(arr, low, lt-1, pivotSelector)
		rightStats := quickSortThreeWay(arr, gt+1, high, pivotSelector)

		stats.Comparisons += leftStats.Comparisons + rightStats.Comparisons
		stats.Swaps += leftStats.Swaps + rightStats.Swaps
	}
	return stats
}

func quickSortDualPivot(arr []int, low, high int) *QuickSortStats {
	stats := &QuickSortStats{}
	if low < high {
		lt, gt, partStats := dualPivotPartition(arr, low, high)

		stats.Comparisons += partStats.Comparisons
		stats.Swaps += partStats.Swaps

		leftStats := quickSortDualPivot(arr, low, lt-1)
		middleStats := quickSortDualPivot(arr, lt+1, gt-1)
		rightStats := quickSortDualPivot(arr, gt+1, high)

		stats.Comparisons += leftStats.Comparisons + middleStats.Comparisons + rightStats.Comparisons
		stats.Swaps += leftStats.Swaps + middleStats.Swaps + rightStats.Swaps
	}
	return stats
}

func generateTestData(size int, pattern string) []int {
	arr := make([]int, size)
	switch pattern {
	case "random":
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(size)
		}
	case "sorted":
		for i := 0; i < size; i++ {
			arr[i] = i
		}
	case "reversed":
		for i := 0; i < size; i++ {
			arr[i] = size - i
		}
	case "fewUnique":
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(5)
		}
	case "triangular":
		half := size / 2
		for i := 0; i < half; i++ {
			arr[i] = i
		}
		for i := half; i < size; i++ {
			arr[i] = size - i - 1
		}
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())

	sizes := []int{1000, 5000, 10000, 50000} // Example sizes
	patterns := []string{"random", "sorted", "reversed", "fewUnique", "triangular"}

	for _, size := range sizes {
		for _, pattern := range patterns {
			data := generateTestData(size, pattern)
			start := time.Now()
			stats := quickSort(data, 0, len(data)-1, lomutoPartition, choosePivotLast)
			stats.ExecutionTime = time.Since(start)
			println("Size:", size, "Pattern:", pattern, "Method: Lomuto Last Pivot", "Comparisons:", stats.Comparisons, "Swaps:", stats.Swaps, "Execution Time:", stats.ExecutionTime)

			data = generateTestData(size, pattern)
			start = time.Now()
			stats = quickSortThreeWay(data, 0, len(data)-1, choosePivotMedianOfThreeRandom)
			stats.ExecutionTime = time.Since(start)
			println("Size:", size, "Pattern:", pattern, "Method: Three-Way Median Random Pivot", "Comparisons:", stats.Comparisons, "Swaps:", stats.Swaps, "Execution Time:", stats.ExecutionTime)

			data = generateTestData(size, pattern)
			start = time.Now()
			stats = quickSortDualPivot(data, 0, len(data)-1)
			stats.ExecutionTime = time.Since(start)
			println("Size:", size, "Pattern:", pattern, "Method: Dual Pivot", "Comparisons:", stats.Comparisons, "Swaps:", stats.Swaps, "Execution Time:", stats.ExecutionTime)
		}
	}
}
