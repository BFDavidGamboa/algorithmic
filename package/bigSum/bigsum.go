package bigsum

func AVeryBigSum(ar []int64) (a int64) {
	for _, v := range ar {
		a += int64(v)
	}
	return a
}
