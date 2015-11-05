package main


import (
  "fmt"
)


func main() {
	
	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16, 7, 5, 2)

	fmt.Println(bestFinish)
}

/**
 * [func_name description]
 * @param  {[type]} finishes ...int        [description]
 * @return {[type]}          [description]
 */
func bestLeagueFinishes(finishes ...int) int{
	
	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}