package PyramidArray


func Pyramid(n int) [][]int {
	// your code here

	res := [][]int{}
	for i := 1; i <= n; i++ {

		ar := []int{}
		for j := 1; j <= i; j++ {
			ar = append(ar, 1)
		}
		res = append(res, ar)
	}
	return res

}
