package showMe

import (
	"fmt"
	"regexp"
)

type FormatData func(data interface{}) (string, error)

type HandlerData struct {
	title        string
	Data         interface{}
	formatedData string
}

func (hd *HandlerData) Format(format FormatData) {
	var err error
	hd.formatedData, err = format(hd.Data)
	if err != nil {
		hd.formatedData = fmt.Sprintf("Has a error when try to formate %v", hd.Data)
	}
}

func (h *HandlerData) SetTitle(optionalTitle []string) {
	h.title = identTitle

	if len(optionalTitle) > 0 {
		h.title += optionalTitle[0]
	}
}

func (h HandlerData) Show() {
	title := regexp.MustCompile(" ").ReplaceAllString(h.title, mark)
	longTitle := longLine - len(title)
	footer := ""

	for i := 0; i < longLine; i++ {
		if longTitle > i {
			title += mark
		}
		footer += mark
	}
	fmt.Printf("\n%s\n%s\n%s\n", title, h.formatedData, footer)
}

func newHadlerData(data interface{}, title []string) HandlerData {
	var hd HandlerData
	hd.SetTitle(title)
	hd.Data = data
	return hd
}
