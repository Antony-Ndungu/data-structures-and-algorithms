package algorithms

func SelectionSort(s []int) {
	var indexOfMin int
	for i := 0; i < len(s)-1; i++ {
		indexOfMin = i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				indexOfMin = j
			}
		}
		s[i], s[indexOfMin] = s[indexOfMin], s[i]
	}
}
