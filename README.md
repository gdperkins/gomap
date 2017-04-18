# GoMap [![Build Status](https://travis-ci.org/gdperkins/gomap.svg?branch=master)](https://travis-ci.org/gdperkins/gomap?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/gdperkins/gomap?branch=master)](https://goreportcard.com/report/github.com/gdperkins/gomap?branch=master) [![Coverage Status](https://coveralls.io/repos/github/gdperkins/gomap/badge.svg?branch=master)](https://coveralls.io/github/gdperkins/gomap?branch=master)

Simple package to map a source struct to a destination struct. 

## Examples

Structs used for all examples:

```go
type Employee struct {
    FirstName string
    LastName  string
    Salary    int
    //other sensitive fields
}

type PublicEmployee struct {
    FirstName string
    LastName  string
}

```

### Default mapping

No configuration required, fields are mapped from the source fields to the target fields in a case sensitive manner.

```go
viewModel := PublicEmployee{}

gm := gomap.New()
gm.Map(Employee{"John", "Doe",  1006}, &viewModel)
fmt.Println(viewModel)
//output: {John Doe}
```
### Custom mapping

Sets the target FirstName fields source to the source LastName

```go
source := Employee{"John", "Doe",  1006}
viewModel := PublicEmployee{}

gm := gomap.New()
gm.Add(source, destination, map[string]FieldConfig{
    "FirstName": {
        Source: "LastName",
    },
})

gm.Map(source, &destination)
fmt.Println(viewModel)
//output: {Doe Doe}
```

### Explicitly Ignore Field

```go
source := Employee{"John", "Doe",  1006}
viewModel := PublicEmployee{}

gm := gomap.New()
gm.Add(source, destination, map[string]FieldConfig{
    "LastName": {
        Ignore: true,
    },
})

gm.Map(source, &destination)
fmt.Println(viewModel)
//output: {Doe }
```

## Roadmap

* Nested struct mapping
* Field type casting
* Pre map func invocation
* Slice to slice mapping