package gomap

import (
	"reflect"
)

// Options is a configuration container to pass to the
// new instance methods for GoMap
type Options struct {
	MappingConfig []Mapping
}

// GoMap holds all configuration for any mappings
// registered at the startup
type GoMap struct {
	mappingConfig []Mapping
}

// Mapping between two different structs
type Mapping struct {
	Key        string
	FieldLinks map[string]MapConfig
}

// MapConfig describes rules for the destination field
type MapConfig struct {
	Ignore bool
	Source string
}

// New returns a new gomapme Confguration struct
func New(options Options) *GoMap {
	return &GoMap{
		mappingConfig: options.MappingConfig,
	}
}

// NewDefault returns a plain GoMap func with default configuration
func NewDefault() *GoMap {
	gomap := GoMap{
		mappingConfig: make([]Mapping, 0),
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
	srcType := reflect.TypeOf(s)

	hasConfig, config := g.getConfig(srcType.Name() + dstType.Name())

	// loop the desintation VM fields
	for i := 0; i < dstType.NumField(); i++ {
		var src string
		ft := dstType.Field(i)

		if hasConfig {
			if _, ok := config.FieldLinks[ft.Name]; ok {
				if config.FieldLinks[ft.Name].Ignore {
					continue
				}
				src = config.FieldLinks[ft.Name].Source
			}
		}

		if sv := srcVal.FieldByName(src); sv.IsValid() {
			fv := dstVal.FieldByName(ft.Name)
			//add logic here to cast
			fv.Set(sv)
		}
	}
}

func (g *GoMap) getConfig(key string) (bool, Mapping) {
	var m Mapping
	for i := range g.mappingConfig {
		if g.mappingConfig[i].Key == key {
			return true, g.mappingConfig[i]
		}
	}
	return false, m
}

// Add applys a new mapping between two structs to the
// global configuration
func (g *GoMap) Add(source interface{}, destination interface{}, config map[string]MapConfig) {
	key := reflect.TypeOf(source).Name() + reflect.TypeOf(destination).Name()
	g.mappingConfig = append(g.mappingConfig, Mapping{key, config})
}
