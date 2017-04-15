# GoMap [![Build Status](https://travis-ci.org/gdperkins/gomap.svg?branch=master)](https://travis-ci.org/gdperkins/gomap)

Simple package to map structs together

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

## Todo

* Nested struct mapping
* Field type casting
* Pre map func invocation