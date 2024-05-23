package matematica

import (
	"fmt"
	"strconv"
)

func Media(x ...float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}
	media := total / float64(len(x))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
