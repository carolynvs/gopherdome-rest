package gopherest

import (
	"github.com/carolynvs/gopherdome-json"
)

type Gopher struct {
	Name string `json:"name"`
}

func GetResource(resource, name string) Gopher {
	var gopher Gopher
	result := gopherson.ToJson(Gopher{Name: name})
	gopherson.FromJson(result, &gopher)
	return gopher
}
