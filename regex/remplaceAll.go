package regex

import (
	re "regexp"
)

func RegexReplace(txt, replace, regex string) (sus string, err error) {
	r, err := re.Compile(regex)
	if err != nil {
		return
	}
	sus = r.ReplaceAllString(txt, replace)
	return
}
