# Chain

chain is a library for function chaning.

## how to use

### From

```go
func main() {
	ch := chain.From[int, int](
		func(a int) int {
			return a + 1
		},
		func(a int) int {
			return a * 2
		},
		func(a int) int {
			return a * 3
		})

	switch ch.Ok() {
	case true:
		println(ch.Unwrap().Run(1))
	case false:
		println(ch.Err().Error())
	}
}
```

```bash
12
```

#### with wrong input & output matching

```go
func main() {
	ch := chain.From[int, string](
		func(a int) int {
			return a + 1
		},
		func(a int) int {
			return a * 2
		},
		func(a int) int {
			return a * 3
		})

	switch ch.Ok() {
	case true:
		println(ch.Unwrap().Run(1))
	case false:
		println(ch.Err().Error())
	}
}
```

```bash
function's return type is invalid
```

### WithResult

```go
func main() {
	ch := chain.WithResult[int, int](
		result.ToFunctor(func(x int) (int, error) {
			return x + 1, nil
		}),
		result.ToFunctor(func(x int) (int, error) {
			return x * 2, nil
		}),
		result.ToFunctor(func(x int) (int, error) {
			return x + 1, nil
		}),
	)

	switch ch.Ok() {
	case true:
		rs := ch.Unwrap().Run(result.Ok(10))
		switch rs.Ok() {
		case true:
			println(rs.Unwrap())
		case false:
			println(rs.Err().Error())
		}
	case false:
		println(ch.Err().Error())
	}
}
```

```bash
23
```
