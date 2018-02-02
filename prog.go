package mypkg

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

func checkIP(x string) bool {
	var d = 0
	s := strings.Split(x, ".")

	for i := 0; i < 4; i++ {
		i1, err := strconv.Atoi(s[i])
		if err == nil {
			//fmt.Println(i1)
		}

		if i1 >= 0 && i1 <= 255 && len(s[i]) <= 3 {
			d++
		}
	}
	return d == 4
}
func Rmspecial(x string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(x); i++ {
		r := rune(x[i])
		if unicode.IsUpper(r) || unicode.IsLower(r) || unicode.IsDigit(r) {
			buffer.WriteString(string(r))
		}
	}
	return buffer.String()
}

//func main() {
//var ip = "172.0.5.3"
//fmt.Printf("%s", Rmspecial(ip))

//}
