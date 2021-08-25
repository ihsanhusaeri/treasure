package main

import (
	"fmt"
	"strings"
)

func main() {
	y5 := "########"
	y4 := "#......#"
	y3 := "#.###..#"
	y2 := "#...#.##"
	y1 := "#X#....#"
	y0 := "########"
	layout := []string{y5, y4, y3, y2, y1, y0}

	findTreasure(layout)
}

func findTreasure(layout []string) {
	for i := 0; i < len(layout); i++ {
		y := layout[i]
		for j := 0; j < len(layout[i]); j++ {
			symbol := y[j : j+1]
			if symbol == "." {
				coordinateY := len(layout) - (i + 1)
				coordinateX := j
				fmt.Printf("coodinate(%d,%d):\n", coordinateX, coordinateY)
				newY := replaceAtIndex(y, '$', j)
				layout[i] = newY
				printLayout(layout)
				layout[i] = y
			}
		}
	}
}
func printLayout(layout []string) {
	layoutString := strings.Join(layout, "\n")
	fmt.Println(layoutString)
}
func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
