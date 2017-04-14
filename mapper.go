package gomapper

import (
	"reflect"
)

// Options is a configuration container to pass to the
// new instance methods for GoMap
type Options struct {
	Maps         []Mapping
	IgnoreFields []string
	IgnoreNil    bool
}

// GoMap holds all configuration for any mappings
// registered at the startup
type GoMap struct {
	maps         []Mapping
	ignoreFields []string
}

// Mapping between two different structs
type Mapping struct {
	Source      interface{}
	Destination interface{}
	FieldLinks  map[string]string
}

// New returns a new gomapme Confguration struct
func New(options Options) *GoMap {
	return &GoMap{
		maps:         options.Maps,
		ignoreFields: options.IgnoreFields,
	}
}

// NewDefault returns a plain GoMap func with default configuration
func NewDefault() *GoMap {
	gomap := GoMap{
		maps:         make([]Mapping, 0),
		ignoreFields: make([]string, 0),
	}
	return &gomap
}

// Map transforms the input struct to the output struct. Always pass the
// destination by reference and the source by value
func (g *GoMap) Map(s interface{}, d interface{}) {
	dstPtrVal := reflect.ValueOf(d)
	dstPtrType := dstPtrVal.Type()
	dstType := dstPtrType.Elem()
	dstVal := reflect.Indirect(dstPtrVal)
	srcVal := reflect.ValueOf(s)

	checkIgnore := len(g.ignoreFields) > 0

	// loop the desintation VM fields
	for i := 0; i < dstType.NumField(); i++ {

		ft := dstType.Field(i)
		sv := srcVal.FieldByName(ft.Name)

		if checkIgnore && ignoreMatch(g.ignoreFields, ft.Name) {
			continue
		}

		if sv.IsValid() {
			fv := dstVal.FieldByName(ft.Name)
			fv.Set(sv)
		}
	}
}

func ignoreMatch(ignores []string, field string) bool {
	for i := range ignores {
		if ignores[i] == field {
			return true
		}
	}
	return false
}

// Add applys a new mapping between two structs to the
// global configuration
func (g *GoMap) Add(m Mapping) {
	g.maps = append(g.maps, m)
}
