package main

import (
	"bytes"
	"log"
	"reflect"

	"github.com/CloudyKit/jet"
)

func main() {

	var View = jet.NewHTMLSet("./views")

	apples := map[string]*Apple{
		"honeycrisp": {
			Flavor: "crisp",
		},
		"red-delicious": {
			Flavor: "poor",
		},
		"granny-smith": {
			Flavor: "tart",
		},
	}

	t, err := View.GetTemplate("example.jet")
	if err != nil {
		panic(err)
	}

	var w bytes.Buffer
	vars := make(jet.VarMap)

	vars.SetFunc("GetAppleByName", func(a jet.Arguments) reflect.Value {
		name := a.Get(0).String()
		return reflect.ValueOf(apples[name])
	})

	vars.SetFunc("TellFlavor", func(a jet.Arguments) reflect.Value {
		apple := a.Get(0).Interface().(*Apple)
		flav := apple.GetFlavor()
		return reflect.ValueOf(flav)
	})

	if err = t.Execute(&w, vars, nil); err != nil {
		panic(err)
	}

	log.Println(w.String())
}

type Apple struct {
	Flavor string
}

func (a *Apple) GetFlavor() string {
	return a.Flavor
}
