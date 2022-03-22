package showMe

import (
	"encoding/json"
	"fmt"
)

func LikeAJson(data interface{}, optionalTitle ...string) {
	title := ""

	if len(optionalTitle) > 0 {
		title = optionalTitle[0]
	}

	j, err := json.MarshalIndent(data, "", ident)
	if err != nil {
		msg := fmt.Sprintf("hubo un error al parsear el json: %v", err.Error())
		show(title, msg)
		return
	}
	show(title, string(j))
}
