package main
import (
	"fmt"
	"strings"
)
func trimExp(s string) string {
	s = strings.Replace(s, "e+0", "e+", 1)
	s = strings.Replace(s, "e-0", "e-", 1)
	return s
}
func main() {
	fmt.Println(trimExp(fmt.Sprintf("%.1e", 1234.0)))
	fmt.Println(trimExp(fmt.Sprintf("%.1e", 0.0)))
	fmt.Println(trimExp(fmt.Sprintf("%.1e", 1e100)))
    fmt.Println(trimExp(fmt.Sprintf("%.1g", 1234.0)))
}
