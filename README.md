# zx

A high order functions collection for golang

## Slice

1. Map
2. Reduce
3. Any
4. All
5. BinarySearch
6. Count
7. Filter
8. FilterIndex
9. FirstOf
10. FirstIndexOf
11. LastOf
12. LastIndexOf
13. Reverse
14. Foreach
15. GroupBy
16. Max
17. Min
18. Random
19. Shuffle
20. Sort
21. Chunk
22. Zip
23. JoinToString

## Set

1. Add
2. Remove
3. Contains
4. FromSlice
5. FromMapKey
6. FromMapValue
7. ToSlice
8. Union
9. Intersect
10. Subtract
11. Equal
12. IsSubset
13. IsSuperset
14. IsDisjoint

## Table

1. MapKey
2. MapValue
3. Reduce
4. Any
5. All
6. Count
7. Filter
8. Foreach
9. Random
10. FromSlices

## Result

1. Success(T) Result[T]
2. Failed(error) Result[T]
3. Result[T].Ok() bool
4. Result[T].Unwrap() T
5. Result[T].Err() error

## Option

1. Some(T) Result[T]
2. None() Result[T]
3. Result[T].Ok() bool
4. Result[T].Unwrap() T