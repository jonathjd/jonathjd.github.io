---
layout: post
title:  Learning Go through Linear Algebra
date: 2024-03-23 10:40:16
description: Performing operations on vectors and matrices using Go
tags: technical data programming
categories: mathematics code
giscus_comments: true
---

# Intro

My first programming language was Python which is what I use every day for work.

My second was C++ and I liked it more than Python. The analogy I always use is C++ feels like I'm driving a nice, reliable, fuel effecient honda civic and Python feels like a Lamborghini, cool, sleek, and stylish, but much more difficult to handle. 

I say that because languages like C++ are more verbose and force you to explicitly type out what you want to do and how you are going to do it, so you gotta go a little slower to make it work.

Python abstracts a lot of that away so you can write code faster, but some times it can be overwhelming because so much is taken care of "under the hood".

That's why I'm going to _Go_ with a compromise. 

For today's project we will be using **Go 🎉**! 

It's not quite as verbose as C++, but doesn't abstract as much away as Python. Plus it looks cool so why not?

More specifically, we will be performing operations on scalars, vectors, and matrices, and learning Go fundamentals along the way.

Let's get started 😎

# Setting up Go

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/proj_banner.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

The great thing about Go, is that it is SUPER easy to install and get going with. To install head over to https://go.dev/dl/ and install the version that matches your machine.

Then type the following in your terminal

```bash
$ go version
```

If your Go version was printed, you're good to _Go_!

Open up a terminal and create a .go file. I named mine **linear_algebra.go**

(The .go extension is all you need to tell VS code that you want to write some Go code.)

Let's start with a simple program that print's "Let's do some math!" to the terminal.

```go
package main

import "fmt"

func main() {
	fmt.Println("Let's do some math!")
}
```

Firstly, "_package main_" tells the computer that this is where our main function is.

Second, the import "fmt" statement brings in the fmt package, which provides functions for I/O operations, including printing to the terminal.

Third, we define a main function that when executed, prints "Let's do some math!" to the terminal.

Go is what is called a **compiled language** meaning that first the computer compiles the program to machine code. After compilation, the resulting executable program can be run directly by the computer.

To compile the program we can run-

```bash
$ go build linear_algebra.go  
```

Then to execute it, we run-

```bash
$ ./linear_algebra
```

Nice, now let's do some math.

# Adding, subtracting, multiplying, and dividing scalars

A **scalar** is another word for a number that is fully described by a single value. Just think of it as a number.

The first thing we are going to do is assign our variables.

```go
	var x int16 = 5
	var y int16 = 10
```

**var** indicates that we are declaring a variable.

Then the variable name (x and y).

Then we tell the computer what kind of variable it is (in this case int16).

Finally we assign the numbers 5 and 10 to the variables!

Adding these variables and printing them out to the screen is fairly straightforward-

First we declare a variable (addition) that is the result of variables x + y. Then we print it out to the terminal.

```go
var addition int16 = (x + y)
fmt.Println("x + y =", addition)
```

We can continue this with the rest of the operations!

```go
package main

import "fmt"

func main() {

	fmt.Println("Let's do some math!")

	var x int16 = 5
	var y int16 = 10

	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)

	// addition
	var addition int16 = (x + y)
	fmt.Println("x + y =", addition)

	// subtraction
	var subtraction int16 = (x - y)
	fmt.Println("x - y =", subtraction)

	// division
	var division float32 = (float32(x) / float32(y))
	fmt.Println("x / y =", division)

	// multiplication
	var multiplication int16 = (x * y)
	fmt.Println("x * y =", multiplication)

}
```

If you noticed, because we declared x and y as int16 if we divide them we would not get the remainder. To get a fraction returned, we can convert the variables into floating point numbers!

This was pretty easy, let's move on to vectors and matrices.

# Adding, subtracting, and calculating the Hadamard and dot product of vectors

First, let's define our vectors *x* and *y*.

```go
var xArray = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
var yArray = [5]float32{6.0, 7.0, 8.0, 9.0, 10.0}
```

First we declare our variable (xArray and yArray).

The **[5]** tells the computer how much memory to allocate for this array.

We then tell the computer what kind of elements will be contained within our array (floating point numbers).

Finally, we define our array with a list of values.

Now, let's add and subtract them! To do this, we will iterate through the array and add the corresponding element in each place.

### Addition

$$
\mathbf{C} = \mathbf{A} + \mathbf{B} = (A_1 + B_1, A_2 + B_2)
$$

```go
// addition
var arrayAddition [5]float32
for i := 0; i < len(xArray); i++ {
    arrayAddition[i] = xArray[i] + yArray[i]
}
fmt.Println("x + y =", arrayAddition)
```

The same with subtraction-

### Subtraction

$$
\mathbf{D} = \mathbf{A} - \mathbf{B} = (A_1 - B_1, A_2 - B_2)
$$

```go
// subtraction
var arraySubtraction [5]float32
for i := 0; i < len(xArray); i++ {
    arraySubtraction[i] = xArray[i] - yArray[i]
}
fmt.Println("x - y =", arraySubtraction)
```
> Technically, you cannot perform division on vectors or matrices. So instead we will calculate the Hadamard and dot product

### Hadamard Product

The Hadamard Product is element-wise multiplication of a vector/matrix. This is often used in image processing which I hope we can get to in another project!

$$
\mathbf{C} = \mathbf{A} \circ \mathbf{B} = (A_1 \cdot B_1, A_2 \cdot B_2, \ldots, A_n \cdot B_n)
$$

```go
// Hadamard Product
var arrayMultiplication [5]float32
for i := 0; i < len(xArray); i++ {
    arrayMultiplication[i] = xArray[i] * yArray[i]
}
fmt.Println("x * y =", arrayMultiplication)
```

### Dot Product

$$
\mathbf{A} \cdot \mathbf{B} = A_1B_1 + A_2B_2 + \ldots + A_nB_n
$$

```go
// dot product
var dotProduct float32
for i := 0; i < len(xArray); i++ {
    dotProduct += xArray[i] / yArray[i]
}
fmt.Println("x * y =", dotProduct)
```
Finally, let's multiply two matrices!

# Multiplying matrices

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/matrix_multiplication.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

Adding and subtracting matrices and fairly straightforward, so we will focus on multipling matrices.

Multiplying matrices is not the same as multiplying each corresponding element together, as that would return the Hadamard product that we calculated earlier. 

To multiply to matrices, the element in the i<sup>th</sup> row and j<sup>th</sup> column of the resulting matrix is obtained by taking the dot product of the i<sup>th</sup> row of the first matrix with the j<sup>th</sup> column of the second matrix. Here is the equation-

$$
C_{ij} = \sum_{k=1}^{n} A_{ik} \times B_{kj}
$$

First, we are going to declare a variable with a new data type called a **matrix**.

```go
type Matrix [2][2]float32
```

Then, we will declare our matrices _A_ and _B_.

```go
var aMatrix = Matrix{{1, 2}, {3, 4}}
var bMatrix = Matrix{{5, 6}, {7, 8}}
```

We'll stick with a 2x2 matrix for today.

```go
var matrixMultiplication Matrix
for i := 0; i < len(aMatrix); i++ {
    for j := 0; j < len(aMatrix); j++ {
        for k := 0; k < len(aMatrix); k++ {
            matrixMultiplication[i][j] += aMatrix[i][k] * bMatrix[k][j]
        }
    }
}
```

This looks a little confusing but let's break it down a bit.

The first Loop (i loop) iterates over each row of **aMatrix**.

The second Loop (j loop) iterates over each column of bMatrix. Because both matrices are 2x2, it does not matter if we use `len(aMatrix)` or `len(bMatrix)`

The third Loop (k loop) iterates over each element within the current row of aMatrix and the current column of bMatrix. This loop performs the multiplication of elements and their accumulation. 

The variable k represents the position within the row of aMatrix and the column of bMatrix being multiplied together.

`matrixMultiplication[i][j] += aMatrix[i][k] * bMatrix[k][j]` performs the matrix multiplication. 

For each position (i, j) in the resulting matrix, it accumulates the product of the k<sup>th</sup> element of the ith row of aMatrix and the k<sup>th</sup> element of the j<sup>th</sup> column of bMatrix.

> The i<sup>th</sup> and j<sup>th</sup> and k<sup>th</sup> nomenclature can be a bit much sometimes, just sit with it and re-read it a few times if it doesn't click immediately.

Now let's run our program!

```bash
$ go build linear_algebra.go
$ ./linear_algebra
```

## Outro
And that's it! We just learned a bit of Go, as well as performed some matrix operations at the same time. Later we'll build on this and calculate the determinate, perform singular value decomposition, and more all in Go!

You can find the code for this in my Github repo [here](https://github.com/jonathjd/thirty_projects).

Project 3 ✅

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/vader.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

<script type="text/javascript" src="https://cdnjs.buymeacoffee.com/1.0.0/button.prod.min.js" data-name="bmc-button" data-slug="jdickinson" data-color="#5F7FFF" data-emoji=""  data-font="Lato" data-text="Buy me a coffee" data-outline-color="#000000" data-font-color="#ffffff" data-coffee-color="#FFDD00" ></script>

<hr>
