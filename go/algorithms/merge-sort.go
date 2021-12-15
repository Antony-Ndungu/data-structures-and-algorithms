package algorithms

func MergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}
	return merge(MergeSort(s[:len(s)/2]), MergeSort(s[len(s)/2:]))
}

func merge(s1, s2 []int) []int {
	var i, j int
	s3 := make([]int, 0, len(s1)+len(s2))
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			s3 = append(s3, s1[i])
			i++
		} else {
			s3 = append(s3, s2[j])
			j++
		}
	}
	for i < len(s1) {
		s3 = append(s3, s1[i])
		i++
	}

	for j < len(s2) {
		s3 = append(s3, s2[j])
		j++
	}
	return s3
}
