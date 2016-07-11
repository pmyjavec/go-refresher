# Go Notes

Re-leaning Go after some time off, useful notes.

* Programs are seperated into 'packages', to use a package import it using the `import` statement
* Strings are just a sequence of bytes, an array
* Commonly used basic types, [see all](http://www.golangbootcamp.com/book/types) : `uint`, `int`, `float32`, `float64`, `byte`, `rune`,`string`
* `rune`, alias for `int32`
* Create variables using `var` or `:=`, examples: `var feet float64` or `feet := 2.0`
* Create constants using `const`, example: `cont hello string = 'Hello World!'`
* Loops, Go only has one type of loop, the `for` loop,  no while, do, until etc
* Loops can be written like `for i <= 10 {}` or `for i := 1; i <= 10; i++ {}`
* Other control strucures are `if` and `switch`
* Functions named after types perform conversions, `float64(24) == 24.0`
* Members of an array can only be of same type
* Shorthand for creating an array `x := [5]float64{1,2,3,4,5}`
* A slice is a segement of an array, unlike arrays the length of a slice may change
* Slices are always associated with an underlying array `x := make([]float64, 5)`, a slice can be shorter than the associated array but never longer.
* Creating new slices can be done using the [low : high] expression, `slc := arr[0:5]` or `allElements := arr[:]`
* Maps are hash tables, to create a map and initialize use `make()`, `people := make(map[string]string)`

