package showMe

import "fmt"

func LikeAString(title string, data ...interface{}) {
	msg := ""
	for _, item := range data {
		msg += fmt.Sprintf("\n%v\n", item)
	}
	show(title, msg)
}
