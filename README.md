# GoMap

Simple package to map structs together

## Getting Started

```

type Employee struct {
    Firstname string
    Salary float32
    //other sensitive fields
}

type PublicEmployee struct {
    Firstname string
}

source := Employee{"Bob", 1006}
viewModel := PublicEmployee{}

fmt.Println(vm)

gm := NewDefault()
gm.Map(source, &viewModel)

fmt.Println(vm)

```

## Todo

* Nested struct mapping
* Field type casting