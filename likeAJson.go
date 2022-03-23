package showMe

func LikeAJson(data interface{}, optionalTitle ...string) {
	hd := newHadlerData(data, optionalTitle)
	hd.Format(jsonToString)
	defer hd.Show()
}

func LikeAXml(data interface{}, optionalTitle ...string) {
	hd := newHadlerData(data, optionalTitle)
	hd.Format(xmlToString)
	defer hd.Show()
}
