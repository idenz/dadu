package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type ipemain struct {
	point  int
	dadu   []int
	status bool
}

func throw(input_pemain int, input_dadu int, pemain map[string]ipemain) map[string]ipemain {

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= input_pemain; i++ {
		var (
			npemain = fmt.Sprintf("Pemain " + fmt.Sprint(i))
			dadu    = []int{}
			point   = 0
		)

		if len(pemain[npemain].dadu) > 1 {
			point = pemain[npemain].point
			input_dadu = len(pemain[npemain].dadu)
		}

		for i := 0; i < input_dadu; i++ {
			dadu = append(dadu, rand.Intn(7-1)+1)
		}

		pemain[npemain] = ipemain{
			point: point,
			dadu:  dadu,
		}
	}

	return pemain
}

func findAndDelete(s []int, item int) ([]int, int) {
	index := 0
	value := 0
	for _, n := range s {
		if n != item {
			s[index] = n
			index++
		} else {
			value++
		}
	}

	return s[:index], value
}

func findMinAndMax(a map[string]ipemain) (max int, min int) {
	var minimun int
	var maximum int

	for _, value := range a {
		if value.point < minimun {
			minimun = value.point
		}
		if value.point > maximum {
			maximum = value.point
		}
	}

	return maximum, minimun

}

func findIndexMaxPoint(s map[string]ipemain, point int) string {

	for i, v := range s {
		if v.point == point {
			return i
		}
	}

	return ""
}

func main() {

	input_pemain := 3
	input_dadu := 4

	pemain := map[string]ipemain{}
	pemain_playing := input_dadu

	no := 1
	for pemain_playing != 1 {
		throw := throw(input_pemain, input_dadu, pemain)

		keys := make([]string, 0, len(throw))
		for k, _ := range throw {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		fmt.Println("Giliran " + fmt.Sprint(no) + " lempar dadu :")
		for _, v := range keys {
			fmt.Println(v + " (" + fmt.Sprint(throw[v].point) + ") : " + fmt.Sprint(throw[v].dadu))
			delete_6, value := findAndDelete(throw[v].dadu, 6)
			// after_dadu, _ := findAndDelete(delete_6, 1)

			// fmt.Println("Pemain " + fmt.Sprint(k+2))
			/* Add Value to prev player*/
			// for i := 0; i < value_1; i++ {
			// 	if k == len(keys)-1 {

			// 		// fmt.Println(after_dadu_0, "KWKWKW")
			// 		pemain["Pemain 1"] = ipemain{point: throw["Pemain 1"].point + value, dadu: append(throw["Pemain 1"].dadu, 1)}
			// 		// fmt.Println(pemain)
			// 	} else {
			// 		pemain_after := fmt.Sprintf("Pemain " + fmt.Sprint(k+2))
			// 		pemain[pemain_after] = ipemain{point: throw[pemain_after].point + value, dadu: append(throw[pemain_after].dadu, 1)}
			// 		// fmt.Println(pemain)
			// 	}
			// 	pemain[v] = ipemain{point: throw[v].point + value, dadu: after_dadu}
			// 	// fmt.Println(after_dadu, "delete angka 1")

			// }
			pemain[v] = ipemain{point: throw[v].point + value, dadu: delete_6}

		}

		fmt.Println("")
		fmt.Println("Setelah evaluasi : ")
		for _, v := range keys {
			fmt.Println(v + " (" + fmt.Sprint(pemain[v].point) + ") : " + fmt.Sprint(pemain[v].dadu))
		}

		for _, n := range pemain {
			if len(n.dadu) == 0 {
				pemain_playing--
			}
		}

		max, _ := findMinAndMax(pemain)
		winner := findIndexMaxPoint(pemain, max)

		fmt.Println("")
		fmt.Println("Winner is " + winner)
		no++
	}
}
