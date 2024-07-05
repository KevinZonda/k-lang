# Essential Internal Functions

So, before all, let's talk about some essential internal functions that are used in the language.

We will only talk about 4 functions here, but there are many more functions in the language to help you with your code. The 4 functions are: `print`, `println`, `range` and `panic`.

## `print(...)` & `println(...)`

Both print and println functions are used to output content to the console. It is very useful if you want to see what is happening in your code.

Here is an example of how to use the `print` function:

```k
print("Hello, World!")
```

The output is very straightforward:

```
Hello, World!
```

The `println` function is the same as the `print` function, but it adds a newline character at the end of the output. Here is an example:

```k
println("Hello, World!")
```

The output is:

```
Hello, World!
```

So now, you may be confused, why? They look the same? Here is the answer. If we want to output 2 strings, e.g. 'Hello' and 'World', we can use the `print` function like this:

```k
print("Hello")
print("World")
```

The output is:

```
HelloWorld
```

But if we use the `println` function like this:

```k
println("Hello")
println("World")
```

The output is:

```
Hello
World
```

The difference is that the `println` function adds a newline character at the end of the output.

If you want to print a new line with `print`, you may use `\n` inside the string, but that's a bit more complicated. We'll talk about it later.

## `range(x)`

The `range` function is used to create a range of numbers. It is very useful when you want to iterate over a list of numbers. The `range(x)` will create a range from 0 to x-1, in math we write it as $[0, x)$ (or $[0, x-1]$).

Here is an example of how to use the `range` function:

```k
x = range(5)
```

Then `x` will be a list of numbers from 0 to 4, i.e., `[0, 1, 2, 3, 4]`.

This will be very useful when you want to iterate over a list of numbers. We will talk about it later.

## `panic(err)`

The `panic(err)` function is used to stop the program immediately. It is very useful when you want to stop the program when an error occurs.