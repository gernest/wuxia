package wuxia

//go:generate go-bindata -o data.go -pkg wuxia js/...

func entryScript() string {
	d, err := Asset("js/init.js")
	if err != nil {
		//FIXME: retrun error instead
		return ""
	}
	return string(d)
}
