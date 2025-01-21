package sortingquestions

func Partition01(arr []int, size int) int {
	p1 := 0
	p2 := size - 1
	count := 0
	for p1 < p2 {
		if arr[p1] == 1 {
			for arr[p2] != 0 {
				p2--
			}
			if p1 < p2 {
				arr[p1], arr[p2] = arr[p2], arr[p1]
				count++
			}
		}
		p1++
	}
	return count
}
