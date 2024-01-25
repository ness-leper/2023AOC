package findnumbers

import (
	"fmt"
	"strings"
	"unicode"
)

type PartNumber struct {
	Number string
	xpt    int
	ypt    int
}

func FindNumbers(file []string) {

  var parts []string

	for i := 0; i < len(file); i++ {
		temp := strings.Replace(file[i], ".", " ", -1)
		split := strings.Split(temp, " ")
		for x := 0; x < len(split); x++ {
			if split[x] != "" {
				var build strings.Builder

				for _, r := range split[x] {
					if unicode.IsDigit(r) {
						build.WriteRune(r)
					}
				}
				output := build.String()
				if len(output) > 0 {
          parts = append(parts, output)
				}
			}
		}
	}

  fmt.Println(parts)
}
