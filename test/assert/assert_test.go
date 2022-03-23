package assert_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/test/assert"
)

func Example() {
	assertions := assert.Of(
		assert.If(true, "true is not true"),
		assert.Less(20, 10),
		assert.Greater(10, 20),
		assert.Equal(10, 10),
		assert.DeepEqual("string", "string"),
		assert.Equal("string", "string"),
		assert.NotZero(0),
		assert.Zero("1"),
		assert.NotNil(new(int)),
		assert.Nil[int](nil),
	)
	fmt.Println(assertions.Result())
	// Output:
	// 0: 20 is not less than 10
	// 1: 10 is not greater than 20
	// 2: 0 is zero value
	// 3: 1 is not zero value
}
