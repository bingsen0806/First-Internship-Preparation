# Overview of Go (with comparison against Java)

## Reference

References were put at the beginning because this document heavily relies on the references, and is only meant as a summary of the references.

Should more details are required, the references should be read again.

1. [Golang for Java Developers Part 1](https://levelup.gitconnected.com/go-for-java-devs-the-basics-348fa57f4100)
1. [Golang for Java Developers Part 2](https://levelup.gitconnected.com/go-for-java-devs-pointers-error-handling-and-concurrency-493dad0c5129)

## Syntax

### Type and Function Declarations

- Strongly typed. But has type inference
- `var name string = "Pat"` or `name := "Pat"`
- `func countVowels(s string) int8 {}`
- no need write `void`
- types behind variable name
- functions are first class. can have nested func like python

  - ```
    func main() {
      var add mathOp := func(a int, b int) int { //note here
        return a + b
      }
      var mult mathOp := func(a int, b int) int {
        return a * b
      }
      op(add, 2, 4)
      op(mult, 2, 4)
      func(a float32, b float32) {
        log.Printf("Op result is %f", a / b)
      }(2.25, 4.25)
    }
    ```

- multiple function return values and value assignment

  - ```
    value1, value2 = funcName()
    func name(paramName string) (string, string) {
    return value1, value2
    }
    ```

### Some Java stuff that Go doesn't have

- no semicolon needed
- no ternary operators
- have switch case but no switch expressions
- no access modifiers, capital letter for `public`, lower for `private`
- for the reason above, it's common to name lower case variables then have Uppercase getters

  - ```
    type MyStruct struct {
    var secret string // not accessible outside of MyStruct's module
    }
    func (ms MyStruct) Secret() string { // accessible everywhere
    return ms.string
    }
    ```

- no `final`, use `const` instead
- no `enums`, use `iota` instead, although it cannot be distinguished from int, and resets to 0 whenever `const` appears

  - ```
    type Animal int // we've created a type alias
    const (
    Dog Animal = iota
    Cat
    Fish
    Bird
    )
    ```

- only has Beta version of **Generics** released in Dec 2021

### Go is ~~OO~~ ~~Functional~~ Procedural language

- Go mainly works with primitive types, structs and functions `type StructName struct {...}`
- Like C, structs are **passed-by-copy** unless pointers of it is passed
- structs can have behaviours **attached** to it, but not encapsulated in it

  - ```
    func countMemebers (t Team) int {
    return len(t.members)
    }
    var c = countMembers(t) # need pass Team object t explicitly
    ```

- **Receiver functions** = methods in Java

  - ```
    func (t Team) countMemebers int {
    return len(t.members)
    }
    var c = t.countMembers() # need pass Team object t explicitly
    ```

- no Extension Methods

### Interface

```
type WithMembers interface {
countMembers() int
}
```

- But, if `Team` wants to implement `WithMembers`, no explicit **implements** keyword required. All is required is for `Team` to implement all functions stated in `WithMembers`, and compiler can check for that

### Constructors

- No explicit constructors
- `t := Team{123, “My Team”, abcInc, [],}`
- or `t := Team{id: 123, name: “My Team”, company: abcInc, members: [],}` (named parameters)
- named parameters allows constructor overloading. Non-initialized fields will be zero or nil

### Error Handling

#### Error

- Error is an interface defined as `type error interface { Error() string }`
- Typical function that throws error looks like this - `func doSomething(arg string) (string, error)`
- No distinction between runtime vs checked exceptions like Java, no `try` `catch` `throw` and therefore forces developer to handle error as soon as they arise
- Errors cannot bubble up higher methods like `main` class, use `panic()` instead for this
- Typical case of handling multiple error types:
  - ```
    func Parse(fname string) (File, error)
    file, err := Parse("file.txt")
    if (err != nil) {
        switch err.(type) {
        case *NotFoundError:
            // handle not found errors
        case *EncodingError:
            // handle encoding error
        }
    }
    ```

#### Panic
