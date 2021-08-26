package main

import (
	"fmt"
	"strings"
)

func main() {
	line5 := "########"
	line4 := "#......#"
	line3 := "#.###..#"
	line2 := "#...#.##"
	line1 := "#X#....#"
	line0 := "########"
	lines := []string{line5, line4, line3, line2, line1, line0}

	findTreasure(lines)
}

func findTreasure(lines []string) {
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(lines[i]); j++ {
			symbol := line[j : j+1]
			if symbol == "." {
				coordinateY := len(lines) - (i + 1)
				coordinateX := j
				fmt.Printf("coodinate(%d,%d):\n", coordinateX, coordinateY)

				newLine := replaceAtIndex(line, '$', j)

				//ganti value lines[i] dengan line yang simbol . nya diganti $
				lines[i] = newLine
				printLayout(lines)

				//kembalikan lagi value lines[i] ke asal
				lines[i] = line
			}
		}
	}
}
func printLayout(lines []string) {
	layout := strings.Join(lines, "\n")
	fmt.Println(layout)
}

//replaceAtIndex digunakan untuk mengganti symbol . dengan $
func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
