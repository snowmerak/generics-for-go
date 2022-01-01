# Option

option is optional value

## How to use

### create

```go
someInt := option.Some(100)
noneInt := option.None[int]()

someString := option.Some("Hello, World!")
noneString := option.None[string]()
```

### unwrap

```go
someInt := option.Some(100)
noneInt := option.None[int]()

if someInt.Ok() {
    fmt.Println(someInt.Unwrap())
}

if !noneInt.Ok() {
    log.Fatal(errors.New("noneInt is empty"))
}
```

## Default Parameter

```go
func Add(a *option.Option[int], b *option.Option[int]) int {
	return a.UnwrapOr(100) + b.UnwrapOr(100)
}

func main() {
	fmt.Println(Add(option.Some(1), option.Some(2)))
	fmt.Println(Add(nil, option.Some(2)))
	fmt.Println(Add(nil, nil))
}
```

```bash
3
102
200
```

