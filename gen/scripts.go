package gen

//go:generate go-bindata -o data.gen.go -pkg gen js/...

func entryScript() string {
	d, err := Asset("js/init.js")
	if err != nil {
		//FIXME: retrun error instead
		return ""
	}
	return string(d)
}
