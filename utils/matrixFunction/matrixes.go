package matrixfunction

import "fmt"

func MatrixThousandMillion(a int64) (ar []int64) {

	const oneThousandMillion = 1000000000
	var i int64
	for i = 1; i <= a; i++ {
		ar = append(ar, oneThousandMillion+i)
	}
	return ar

}

func CreateMatrix3x3FromSlice(ar []int32) [][]int32 {

	mtxLen := 3
	mtxOut := make([][]int32, len(ar))

	le := len(ar)
	fmt.Println(le)

	if len(ar) == 9 {
		for i, v := range ar {
			for j := 0; j < mtxLen; j++ {
				mtxOut[j][i] = v
			}
		}
	} else {
		return mtxOut
	}
	return mtxOut
}
