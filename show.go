package showMe

import (
	"fmt"

	"github.com/soni295/showme/regex"
)

func show(title string, data interface{}) {

	title = "  " + title
	const mark = "-"
	title, _ = regex.RegexReplace(title, mark, " ")

	longTitle := longLine - len(title)
	footer := ""

	for i := 0; i < longLine; i++ {
		if longTitle > i {
			title += mark
		}
		footer += mark
	}
	fmt.Printf("\n%v\n%v\n%v\n", title, data, footer)
}
