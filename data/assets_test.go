package data

import "testing"

func TestHTTPAsset(t *testing.T) {
	a := HTTPAsset()
	css := "css/semantic.min.css"
	f, err := a.Open(css)
	if err != nil {
		t.Fatal(err)
	}
	_ = f.Close()
}
