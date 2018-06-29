package sort

type sortable interface {
	LessEqual(i, j int) bool
	Len() int
	Swap(i, j int)
}

// Determines the order of the resulting array
var ascending bool

// QuickSort sorts using the QuickSort algorithm... who would've guessed :)
func QuickSort(s sortable, isAscending bool) {
	ascending = isAscending
	sort(s, 0, s.Len()-1)
}

func sort(s sortable, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(s, start, end)
	sort(s, start, pivot-1)
	sort(s, pivot+1, end)
}

func partition(s sortable, start, end int) int {
	pivot := end
	partitionIndex := start
	for current := start; current < end; current++ {
		if (ascending && !s.LessEqual(current, pivot)) || (!ascending && s.LessEqual(current, pivot)) {
			s.Swap(current, partitionIndex)
			partitionIndex++
		}
	}
	s.Swap(partitionIndex, pivot)
	return partitionIndex
}
