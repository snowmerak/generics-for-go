# Result

result is a type for functional programming

## How to use

### create

```go
okInt := result.Ok(10)
errInt := result.Err[int] (errors.New("anything"))
```

### unwrap

```go
okInt := result.Ok(10)
errInt := result.Err[int] (errors.New("anything"))

if okInt.Ok() {
    fmt.Println(result.Unwrap())
}

if !errInt.Ok() {
    log.Fatal(result.Err())
}
```

## map

### when succeed

```go
a := result.Ok(100)

a.Map(func (t int) (int, error)) {
    return t * 2, nil
})

if a.Ok() {
    fmt.Println(a.Unwrap())
}
```

```bash
200
```

### when failed

```go
a := result.Ok(100)

a.Map(func (t int) (int, error)) {
    return -1, errors.New("some error")
})

if !a.Ok() {
    fmt.Println(a.Err())
}
```

```bash
some error
```

## map or

second parameter is default value

```go
a := result.Ok(100)

a.MapOr(func (t int) (int, error)) {
    return t * 2, nil
}, 99)

fmt.Println(a.Unwrap())

a.MapOr(func (t int) (int, error)) {
    return -1, errors.New("some error")
}, 99)

fmt.Println(a.Unwrap())
```

```bash
99
99
```

## and then

```go
a := result.Ok(100)

a = a.AndThen(func(t int) (int, error) {
    return t * 2, nil
}).AndThen(func (t int) (int, error) {
    return t * 3, nil
})

if a.Ok() {
    fmt.Println(a.Unwrap())
}
```

```bash
600
```

## unwrap or

```go
a := result.Err[int](errors.New("Hello, World!"))

b := a.UnwrapOr(999)

fmt.Println(b)
```

```bash
999
```

## unwrap or panic

```go
a := result.Err[int](errors.New("Hello, World!"))

a.UnwrapOrPanic(errors.New("some error"))
```

```bash
some error
```
