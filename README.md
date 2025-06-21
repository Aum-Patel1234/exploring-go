# GoLang Learning Repository

This repository is my personal reference and playground for learning **Go (Golang)**.  
It includes basic programs, project experiments, and hands-on explorations.

### [Learn Go from Boot.dev](https://boot.dev/learn/learn-golang)


# Go Type Sizes: Integers, Floats, and Complex Numbers

In Go, different numeric types have different **sizes**, which determine how much memory they use.

---

## Whole Numbers (No Decimal)

| Type    | Description                                                      |
| ------- | ---------------------------------------------------------------- |
| `int`   | Signed integer (default size: 32 or 64 bits depending on system) |
| `int8`  | 8-bit signed integer                                             |
| `int16` | 16-bit signed integer                                            |
| `int32` | 32-bit signed integer                                            |
| `int64` | 64-bit signed integer                                            |

---

## Positive Whole Numbers (No Decimal)

**Unsigned integers** ("uint" stands for "unsigned integer"):
They can only store positive numbers (and zero).

| Type      | Description                             |
| --------- | --------------------------------------- |
| `uint`    | Unsigned integer (32 or 64 bits)        |
| `uint8`   | 8-bit unsigned integer                  |
| `uint16`  | 16-bit unsigned integer                 |
| `uint32`  | 32-bit unsigned integer                 |
| `uint64`  | 64-bit unsigned integer                 |
| `uintptr` | Unsigned integer for pointer arithmetic |

---

## Signed Decimal Numbers (Floats)

| Type      | Description                            |
| --------- | -------------------------------------- |
| `float32` | 32-bit floating point number           |
| `float64` | 64-bit floating point number (default) |

---

## Complex Numbers

Complex numbers have a **real** and an **imaginary** part.

| Type         | Description                                 |
| ------------ | ------------------------------------------- |
| `complex64`  | Complex number with float32 parts           |
| `complex128` | Complex number with float64 parts (default) |

---

## What's the Deal With the Sizes?

The size (8, 16, 32, 64, 128, etc.) represents **how many bits in memory** will be used to store the variable.

The "default" sizes depend on the machine architecture:

* On a **32-bit system**, `int` and `uint` are typically 32 bits.
* On a **64-bit system**, `int` and `uint` are typically 64 bits.

### Standard Types to Prefer (Unless You Need Optimizations)

* `int`
* `uint`
* `float64`
* `complex128`

These are typically the best choices for most use cases in Go.

---

## Converting Between Types

In Go, you can **convert** between types using type casting.

### Example: Float to Integer Conversion

```go
temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)
```

### Important:

* Casting a float to an integer **truncates** the decimal part (it does not round).


## Additional Core Types Explained

| Type         | Description                                                                                      |
| ------------ | ------------------------------------------------------------------------------------------------ |
| `bool`       | Represents a boolean value: `true` or `false`.                                                   |
| `string`     | Represents a sequence of characters. Strings are immutable in Go.                                |
| `byte`       | Alias for `uint8`. Represents a single 8-bit byte, commonly used in byte slices.                 |
| `rune`       | Alias for `int32`. Represents a Unicode code point, useful for handling characters beyond ASCII. |

---

# Comparing Go's Speed

Go is generally faster and more lightweight than interpreted or VM-powered languages like:

- Python
- JavaScript
- PHP
- Ruby
- Java

However, in terms of execution speed, Go does lag behind some other compiled languages like:

- C
- C++
- Rust

Go is a bit slower mostly due to its automated memory management, also known as the **Go runtime**. Slightly slower speed is the price we pay for memory safety and simple syntax!
