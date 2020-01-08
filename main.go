package main

import (
	"bytes"
	"log"
	"reflect"
	"sort"

	"github.com/CloudyKit/jet"
)

func main() {

	var View = jet.NewHTMLSet("./views") // relative to the current working directory from where this code is run

	b := Bushel{
		{
			Flavor: "crisp",
		},
		{
			Flavor: "tart",
		},
		{
			Flavor: "mellow",
		},
	}

	t, err := View.GetTemplate("example.jet")
	if err != nil {
		panic(err)
	}
	var w bytes.Buffer // needs to conform to io.Writer interface (like gin's context.Writer for example)
	vars := make(jet.VarMap)
	vars.Set("data", b)

	vars.SetFunc("SortApples", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("SortApples", 1, 1)
		bu := a.Get(0).Interface().(Bushel)
		sort.SliceStable(bu, func(i, j int) bool {
			return bu[i].Flavor < bu[j].Flavor
		})
		return reflect.ValueOf(bu)
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

type Bushel []Apple
