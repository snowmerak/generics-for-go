// tuple is a readonly data set. from 2 to 100.
package tuple

type Tuple2[T1 any, T2 any] struct {
	v1 T1
	v2 T2
}

func N2[T1 any, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{
		v1: v1,
		v2: v2,
	}
}

func (t Tuple2[T1, T2]) T1() T1 {
	return t.v1
}

func (t Tuple2[T1, T2]) T2() T2 {
	return t.v2
}

type Tuple3[T1 any, T2 any, T3 any] struct {
	v1 T1
	v2 T2
	v3 T3
}

func N3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{
		v1: v1,
		v2: v2,
		v3: v3,
	}
}

func (t Tuple3[T1, T2, T3]) T1() T1 {
	return t.v1
}

func (t Tuple3[T1, T2, T3]) T2() T2 {
	return t.v2
}

func (t Tuple3[T1, T2, T3]) T3() T3 {
	return t.v3
}

type Tuple4[T1 any, T2 any, T3 any, T4 any] struct {
	v1 T1
	v2 T2
	v3 T3
	v4 T4
}

func N4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{
		v1: v1,
		v2: v2,
		v3: v3,
		v4: v4,
	}
}

func (t Tuple4[T1, T2, T3, T4]) T1() T1 {
	return t.v1
}

func (t Tuple4[T1, T2, T3, T4]) T2() T2 {
	return t.v2
}

func (t Tuple4[T1, T2, T3, T4]) T3() T3 {
	return t.v3
}

func (t Tuple4[T1, T2, T3, T4]) T4() T4 {
	return t.v4
}

type Tuple5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
}

func N5[T1 any, T2 any, T3 any, T4 any, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{
		v1: v1,
		v2: v2,
		v3: v3,
		v4: v4,
		v5: v5,
	}
}

func (t Tuple5[T1, T2, T3, T4, T5]) T1() T1 {
	return t.v1
}

func (t Tuple5[T1, T2, T3, T4, T5]) T2() T2 {
	return t.v2
}

func (t Tuple5[T1, T2, T3, T4, T5]) T3() T3 {
	return t.v3
}

func (t Tuple5[T1, T2, T3, T4, T5]) T4() T4 {
	return t.v4
}

func (t Tuple5[T1, T2, T3, T4, T5]) T5() T5 {
	return t.v5
}

type Tuple6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any] struct {
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
}

func N6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Tuple6[T1, T2, T3, T4, T5, T6] {
	return Tuple6[T1, T2, T3, T4, T5, T6]{
		v1: v1,
		v2: v2,
		v3: v3,
		v4: v4,
		v5: v5,
		v6: v6,
	}
}

func (t Tuple6[T1, T2, T3, T4, T5, T6]) T1() T1 {
	return t.v1
}

func (t Tuple6[T1, T2, T3, T4, T5, T6]) T2() T2 {
	return t.v2
}

func (t Tuple6[T1, T2, T3, T4, T5, T6]) T3() T3 {
	return t.v3
}

func (t Tuple6[T1, T2, T3, T4, T5, T6]) T4() T4 {
	return t.v4
}

func (t Tuple6[T1, T2, T3, T4, T5, T6]) T5() T5 {
	return t.v5
}

func (t Tuple6[T1, T2, T3, T4, T5, T6]) T6() T6 {
	return t.v6
}

type Tuple7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any] struct {
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
}

func N7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return Tuple7[T1, T2, T3, T4, T5, T6, T7]{
		v1: v1,
		v2: v2,
		v3: v3,
		v4: v4,
		v5: v5,
		v6: v6,
		v7: v7,
	}
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) T1() T1 {
	return t.v1
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) T2() T2 {
	return t.v2
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) T3() T3 {
	return t.v3
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) T4() T4 {
	return t.v4
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) T5() T5 {
	return t.v5
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) T6() T6 {
	return t.v6
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) T7() T7 {
	return t.v7
}

type Tuple8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any] struct {
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
	v8 T8
}

func N8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
		v1: v1,
		v2: v2,
		v3: v3,
		v4: v4,
		v5: v5,
		v6: v6,
		v7: v7,
		v8: v8,
	}
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T1() T1 {
	return t.v1
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T2() T2 {
	return t.v2
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T3() T3 {
	return t.v3
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T4() T4 {
	return t.v4
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T5() T5 {
	return t.v5
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T6() T6 {
	return t.v6
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T7() T7 {
	return t.v7
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) T8() T8 {
	return t.v8
}

type Tuple9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any] struct {
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
	v8 T8
	v9 T9
}

func N9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
		v1: v1,
		v2: v2,
		v3: v3,
		v4: v4,
		v5: v5,
		v6: v6,
		v7: v7,
		v8: v8,
		v9: v9,
	}
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T1() T1 {
	return t.v1
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T2() T2 {
	return t.v2
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T3() T3 {
	return t.v3
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T4() T4 {
	return t.v4
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T5() T5 {
	return t.v5
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T6() T6 {
	return t.v6
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T7() T7 {
	return t.v7
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T8() T8 {
	return t.v8
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) T9() T9 {
	return t.v9
}

type Tuple10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	v7  T7
	v8  T8
	v9  T9
	v10 T10
}

func N10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
		v1:  v1,
		v2:  v2,
		v3:  v3,
		v4:  v4,
		v5:  v5,
		v6:  v6,
		v7:  v7,
		v8:  v8,
		v9:  v9,
		v10: v10,
	}
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T1() T1 {
	return t.v1
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T2() T2 {
	return t.v2
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T3() T3 {
	return t.v3
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T4() T4 {
	return t.v4
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T5() T5 {
	return t.v5
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T6() T6 {
	return t.v6
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T7() T7 {
	return t.v7
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T8() T8 {
	return t.v8
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T9() T9 {
	return t.v9
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) T10() T10 {
	return t.v10
}
