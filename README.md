# QuickSort Variants in Go

This repository contains an implementation of various QuickSort variants in Go, along with utilities for testing and performance measurement. The project aims to compare and analyze different partitioning schemes and pivot selection methods used in QuickSort.

## Features

### Partitioning Schemes
- **Lomuto Partition**
- **Hoare Partition**
- **Three-Way Partition**
- **Dual-Pivot Partition**

### Pivot Selection Methods
- **Last Element Pivot**
- **Random Pivot**
- **Median of Three Pivot**
- **Median of Three Random Pivot**

### Data Patterns
- **Random**
- **Sorted**
- **Reversed**
- **Few Unique**
- **Triangular**

## Code Overview

### Key Functions

#### `lomutoPartition`
- Implements the Lomuto partition scheme.
- Swaps elements based on comparison with the pivot.

#### `hoarePartition`
- Implements the Hoare partition scheme.
- Ensures elements less than or equal to the pivot are on one side.

#### `threeWayPartition`
- Divides the array into three parts:
  1. Elements less than the pivot.
  2. Elements equal to the pivot.
  3. Elements greater than the pivot.

#### `dualPivotPartition`
- Uses two pivots to divide the array into three segments.
- Performs fewer comparisons in some scenarios.

#### `quickSort`
- Generalized QuickSort function.
- Accepts partition and pivot selection functions as parameters.

#### `quickSortThreeWay`
- Implements QuickSort using the three-way partitioning scheme.

#### `quickSortDualPivot`
- Implements QuickSort using the dual-pivot partitioning scheme.

#### `generateTestData`
- Generates test data for sorting.
- Supports multiple patterns, such as random and sorted arrays.

### Performance Metrics
- **Comparisons**: Number of element comparisons.
- **Swaps**: Number of element swaps.
- **Execution Time**: Time taken to sort the array.

### Example Output
```
Size: 1000 Pattern: random Method: Lomuto Last Pivot Comparisons: 11718 Swaps: 5680 Execution Time: 512400
Size: 1000 Pattern: random Method: Three-Way Median Random Pivot Comparisons: 8573 Swaps: 8212 Execution Time: 532100
Size: 1000 Pattern: random Method: Dual Pivot Comparisons: 7831 Swaps: 4780 Execution Time: 489200
```

## How to Run

1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd quicksort-analysis
   ```

2. Run the program:
   ```sh
   go run main.go
   ```

3. Analyze the output to compare performance metrics across different QuickSort variants.

## Dependencies
- Go 1.17 or higher.

