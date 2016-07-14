# Go Notes

Re-leaning Go after some time off, useful notes.

* Programs are seperated into 'packages', to use a package import it using the `import` statement
* Strings are just a sequence of bytes, an array
* Commonly used basic types, [see all](http://www.golangbootcamp.com/book/types) : `uint`, `int`, `float32`, `float64`, `byte`, `rune`,`string`
* `rune`, alias for `int32`
* Create variables using `var` or `:=`, examples: `var feet float64` or `feet := 2.0`
* Create constants using `const`, example: `cont hello string = 'Hello World!'`
* Loops, Go only has one type of loop, `for`,  no while, do, until etc
* Loops can be written like `for i <= 10 {}` or `for i := 1; i <= 10; i++ {}`
* Other control strucures are `if` and `switch`
* Functions named after types perform conversions, `float64(24) == 24.0`
* Members of an array can only be of same type
* Shorthand for creating an array `x := [5]float64{1,2,3,4,5}`
* A slice is a segement of an array, unlike arrays the length of a slice may change
* Slices are always associated with an underlying array `x := make([]float64, 5)`, a slice can be shorter than the associated array but never longer.
* Creating new slices can be done using the [low : high] expression, `slc := arr[0:5]` or `allElements := arr[:]`
* Maps are hash tables, to create a map and initialize use `make()`, `people := make(map[string]string)`
* `defer` schedules a function call to be run upon completion of the containing block
* Functions can be called recursively
* Closures or functions within functions are supported, closures have access to variables in the containing block
* Pointers can be created using `new()`, dereferenced with `*`, the `&` operator will return a pointer to an existing variable.
* A `struct` type allows us to easily represent real world objects by groupd of different types using named fields.
* Structs also support Embeded types, this is analoguous to mixins in Ruby. A Dog type can easily behave like an Animal by emdedding the Animal type and receiving all it's methods.
* Go supports methods and can be defined similar to functions, just a receiver (usually a `struct`) needs to be specified `func (a *Animal) Eat()`, call a method with the `.` operator. `dog := new(Animal); dog.Eat()`
* Interfaces are supported, easily ensuring anything pased into a function or method meets particular requirements.
* Goroutines are functions capable of running concurrently
* Channels allow Go-routines to communicate
* `select` is similar to switch but for channels
* Channels may be buffered, that is can receive only receive a specified amount of data and will block until the buffer channel is free
* Channel directions may be specified, for example a function may wish to only read from a channel or write to one.
