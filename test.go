
package main

import(
	"fmt"
	"strings"
)


func parseFlagData(flagDat string) []string {
	return strings.Split(flagDat, ",") //split on commas
}

func main() {
	x := "one,two,three"
	y := "testfile.txt"

	fmt.Println(strings.Split(x, ","))
	fmt.Println(parseFlagData(x))
	fmt.Println(strings.Split(y, ","))

	z := strings.Split(y, ",")
	fmt.Println(z[0][len(z[0])-4:])
}
