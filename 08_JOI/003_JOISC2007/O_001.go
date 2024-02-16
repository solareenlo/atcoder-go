package main

import "fmt"

// http://hydra.nat.uni-magdeburg.de/packing/csq/csq90.html
var RESULT = [][]int{{5374995, 5374995}, {23994523, 5374995}, {42614051, 5374995}, {61233580, 5374995}, {79853108, 5374995}, {93732505, 6267495}, {14684759, 10749990}, {33304287, 10749990}, {51923816, 10749990}, {70543344, 10749990}, {84952463, 14838549}, {5374995, 16124984}, {23994523, 16124984}, {42614051, 16124984}, {61233580, 16124984}, {94625005, 19529413}, {75642699, 20213544}, {14684759, 21499979}, {33304287, 21499979}, {51923816, 21499979}, {66332935, 25588539}, {84952463, 25588539}, {23994523, 26874974}, {5374995, 26874974}, {42614051, 26874974}, {94625005, 30279402}, {57023170, 30963533}, {75642699, 30963533}, {14684759, 32249969}, {33304287, 32249969}, {47713406, 36338528}, {66332935, 36338528}, {84952463, 36338528}, {5374995, 37624964}, {23994523, 37624964}, {94625005, 41029392}, {38403642, 41713523}, {57023170, 41713523}, {75642699, 41713523}, {14684759, 42999959}, {29093878, 47088518}, {47713406, 47088518}, {66332935, 47088518}, {84952463, 47088518}, {5374995, 48374953}, {94625005, 51779382}, {19784114, 52463513}, {38403642, 52463513}, {57023170, 52463513}, {75642699, 52463513}, {29093878, 57838508}, {47713406, 57838508}, {66332935, 57838508}, {84952463, 57838508}, {5374995, 59124943}, {15621905, 62375036}, {94625005, 62529371}, {38403642, 63213502}, {57023170, 63213502}, {75642699, 63213502}, {24931669, 67750031}, {47713406, 68588497}, {66332935, 68588497}, {84952463, 68588497}, {5374995, 69874933}, {15621905, 73125026}, {34241434, 73125026}, {94625005, 73279361}, {57023170, 73963492}, {75642699, 73963492}, {24931669, 78500021}, {43551198, 78500021}, {66332935, 79338487}, {84952463, 79338487}, {5374995, 80624922}, {15621905, 83875016}, {34241434, 83875016}, {52860962, 83875016}, {94625005, 84029351}, {75642699, 84713482}, {24931669, 89250010}, {43551198, 89250010}, {62170726, 89250010}, {5374995, 91374912}, {82145225, 93273822}, {15621905, 94625005}, {34241434, 94625005}, {52860962, 94625005}, {71480490, 94625005}, {92809961, 94625005}}

func main() {
	for _, line := range RESULT {
		for i := 0; i < len(line); i++ {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(line[i])
		}
		fmt.Println()
	}
}
