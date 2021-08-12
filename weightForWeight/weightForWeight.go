
package weightForWeight

import (
	"sort"
	"strconv"
	"strings"
)

type People struct {
	Weight       string
	SumDigits int
}

func OrderWeight(strng string) string {
	// your code


	var weights []People

	ar := strings.Split(strng, " ")
	for _, value := range ar {
		weights = append(weights, People{
			Weight:       value,
			SumDigits: sumDigits(value)	})
	}

	sort.Slice(weights, func(i, j int) bool {


		if weights[i].SumDigits == weights[j].SumDigits {
			return weights[i].Weight < weights[j].Weight
		} else {
			return weights[i].SumDigits < weights[j].SumDigits}

	})


	res:=""

	for _, value := range weights {
		res = res + value.Weight + " "
	}

	return strings.Trim(res, " ")

}



func sumDigits(str string) (int) {

	ar:= strings.Split(str, "")

	var  res int

	for _, value := range ar {
			// о нет, что-то пошло не так

		value_, err := strconv.Atoi(value)

		if err==nil {
			res = res + value_

		}

	}
	println(str)
	println(res)

	return res

}
