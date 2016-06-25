package gen

import (
	"encoding/json"
	"errors"

	"github.com/robertkrimen/otto"
)

func getToJSON(o interface{}, vm *otto.Otto, src string) error {
	source := "JSON.stringify( function(){" + src + "}())"
	v, err := vm.Run(source)
	if err != nil {
		return err
	}
	if v.IsUndefined() || v.IsNull() {
		return errors.New("bad script")
	}
	s, err := v.ToString()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(s), o)
	if err != nil {
		return err
	}
	return nil
}
