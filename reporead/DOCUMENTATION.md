# Gomap Technical Documentation

`gomap` is a lightweight, reflection-based object-to-object mapping library for Go. It is designed to simplify the process of transferring data between different struct types (e.g., Domain Models to DTOs or ViewModels) by providing a configurable mapping engine.

---

## 1. Project Architecture Overview

The library follows a **Registry-based Mapping Pattern**. Instead of manual assignment, developers register configurations that define how one type relates to another.

### High-Level System Design
1.  **Registry Phase**: The user registers a `Mapping` between a source struct and a destination struct.
2.  **Configuration Phase**: During registration, specific field behaviors (ignoring fields or renaming sources) are defined.
3.  **Execution Phase**: The `Map` function uses Go's `reflect` package to traverse the source object and populate the destination object based on the registered configuration.

### Data Flow
1.  **Input**: Source struct (populated) and Destination struct pointer (empty).
2.  **Lookup**: The engine generates a unique key based on the reflect types of both objects.
3.  **Transformation**: The engine iterates through the destination fields, applying logic from `FieldConfig` if present, otherwise performing a direct name-match copy.
4.  **Output**: The destination struct is mutated in place, and an error is returned if the mapping fails.

---

## 2. Core Components and Modules

### `GoMap` (The Engine)
The central coordinator of the library. It maintains the internal state of all registered mappings.
- **Responsibilities**: Storing mapping configurations, validating types, and executing the reflection logic.

### `Mapping` (The Configuration)
A struct that defines the relationship between two specific types.
- **Key Fields**:
    - `FieldConfigs`: A map of destination field names to `FieldConfig` objects.

### `FieldConfig` (Field-Level Rules)
Defines how a specific field in the destination struct should be treated.
- **Responsibilities**:
    - **Ignore**: Prevents a field from being mapped.
    - **Source**: Allows mapping a destination field from a source field with a different name.

---

## 3. API Reference

### Initialization
```go
func New() *GoMap
```
Creates and returns a new instance of the `GoMap` engine.

### Registering Mappings
```go
func (gm *GoMap) Add(source interface{}, destination interface{}, config Mapping)
```
Registers a relationship between two types.
- **source/destination**: Instances of the structs (used to infer types).
- **config**: A `Mapping` struct defining custom rules.

### Executing Mappings
```go
func (gm *GoMap) Map(s interface{}, d interface{}) error
```
Performs the data transfer from `s` to `d`.
- **s**: The source data (must be a struct or pointer to struct).
- **d**: The destination (must be a pointer to a struct).
- **Returns**: `error` if types are invalid or mapping fails.

---

## 4. Data Models and Schemas

### Internal Type Definitions
While the library uses `interface{}` for its public API, it internally manages data using these structures:

```go
type FieldConfig struct {
    Ignore bool
    Source string
}

type Mapping struct {
    FieldConfigs map[string]FieldConfig
}
```

### Key Generation
Mappings are indexed internally using a string key generated from the reflect types:
`reflect.TypeOf(source).String() + " -> " + reflect.TypeOf(destination).String()`

---

## 5. Configuration Examples

### Basic Mapping
By default, `gomap` matches fields by name.

```go
gm := gomap.New()
gm.Add(User{}, UserDTO{}, gomap.Mapping{})
```

### Advanced Field Configuration
Use `FieldConfig` to handle naming discrepancies or sensitive data.

```go
config := gomap.Mapping{
    FieldConfigs: map[string]gomap.FieldConfig{
        // Destination field "FullName" will take data from Source field "Name"
        "FullName": { Source: "Name" },
        // Destination field "Password" will be skipped
        "Password": { Ignore: true },
    },
}
gm.Add(User{}, UserViewModel{}, config)
```

---

## 6. Development Guide

### Setup Instructions
1.  **Requirement**: Go 1.13+
2.  **Installation**:
    ```bash
    go get github.com/gdperkins/gomap
    ```

### Build and Test
The project uses standard Go tooling.

**Running Tests**:
```bash
go test -v ./...
```

**Testing Approach**:
The `mapper_test.go` file contains comprehensive test cases covering:
- **TestCanMap**: Basic field-to-field transfer.
- **TestCanIgnoreField**: Ensuring "Ignore" flags are respected.
- **TestCanChangeFieldSource**: Verifying custom source naming logic.
- **TestMapTypeAcceptance**: Ensuring the library handles pointers and structs correctly.
- **TestDuplicateMapConfigsIgnored**: Verifying registry idempotency.

### Best Practices
- **Pre-register**: Always call `Add()` during application initialization (e.g., in `init()` or a global setup) to ensure mappings are ready before `Map()` is called.
- **Pointers**: Always pass a pointer to the destination object in the `Map()` function to allow the library to mutate the object.