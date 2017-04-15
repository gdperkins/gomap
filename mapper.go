package gomap

import (
	"errors"
	"reflect"
)

// Mapping between two different structs
type Mapping struct {
	Key        string
	FieldLinks map[string]FieldConfig
}

// FieldConfig describes rules for the destination field
type FieldConfig struct {
	Ignore bool
	Source string
}

// GoMap holds all configuration for any mappings
// registered at the startup
type GoMap struct {
	mappingConfig []Mapping
}

// New returns a plain GoMap func with default configuration
func New() *GoMap {
	gomap := GoMap{
		mappingConfig: make([]Mapping, 0),
	}
	return &gomap
}

// Map transforms the input struct to the output struct. s (source) has to
// be passed by value and d (destination) needs to be passed by reference.
// Both parameters need to be of type struct or a error will be returned.
func (g *GoMap) Map(s interface{}, d interface{}) error {
	dstVal, dstType, srcVal, srcType, err := validInput(s, d)
	if err != nil {
		return err
	}

	hasConfig, config := g.getConfig(srcType.Name() + dstType.Name())

	// loop the desintation struct fields
	for i := 0; i < dstType.NumField(); i++ {
		ft := dstType.Field(i)
		src := ft.Name

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
	return nil
}

func validInput(s interface{}, d interface{}) (reflect.Value, reflect.Type, reflect.Value, reflect.Type, error) {
	dstPtr := reflect.ValueOf(d)
	dstVal := reflect.Indirect(dstPtr)
	dstType := dstPtr.Type().Elem()
	srcVal := reflect.ValueOf(s)
	srcType := reflect.TypeOf(s)

	if srcType.Kind().String() != "struct" {
		return dstVal, dstType, srcVal, srcType, errors.New("Invalid source parameter type")
	}

	if dstPtr.Kind().String() != "ptr" {
		return dstVal, dstType, srcVal, srcType, errors.New("Invalid destination parameter type")
	}

	return dstVal, dstType, srcVal, srcType, nil
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
func (g *GoMap) Add(source interface{}, destination interface{}, config map[string]FieldConfig) {
	key := reflect.TypeOf(source).Name() + reflect.TypeOf(destination).Name()
	g.mappingConfig = append(g.mappingConfig, Mapping{key, config})
}
