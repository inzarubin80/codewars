package ProductOfConsecutiveFibNumbers

func ProductFib(prod uint64) [3]uint64 {

	var f1 uint64 = 1
	var f2 uint64 = 0
	var multiplication uint64 = 0

	for n := 0; multiplication < prod; n++ {

		f2, f1 = f1, f1+f2
		multiplication = f1 * f2

	}

	var verifying uint64 = 0

	if multiplication == prod {
		verifying = 1
	}

	res := [3]uint64{f1, f2, verifying}
	return res

}
