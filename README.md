

# Notes on Gin
- I think the listener and routing should occur on main.go
- the REST methods should be located in a separate file
- compiling the module to become a daemon.


# Notes on Packages
- a package is the same as a module in JS
- Go will automatically export methods from a package (so long as the method is capitalized)
- when running `go mod init <package_name>`, do not include a root directory unless publishing. The package name should be the root directory name.

# Notes on dotenv
- use this package "github.com/joho/godotenv"
- go has built in support for environment variables, but not for dotenv. Hence the package.
- create a helper function which reads env var and returns a string.

# RFIs
- How do I use headers?
- How do I use cookies?
- How do we access a database next?
- How can we multithread the listener?

# Genaric Notes
Functions do not need to declare void.

Arrays are not objects. Arrays are values and each array is independent.
Use Slices instead of arrays. Slices are mutable, dynamic, and reffrenced; therefore, lower memory usage.

Functions are stored as values. We can create "arrow funcitons" because of this property.
```
func main() {
	myFunction := func(x, y int) int {
		return math.Sqrt(x*x + y*y);
	};
	
	fmt.Println(myFunction(1, 3));
}
```

You can return a function from a function. This is called a closure.
```
func add() func(int) int {
	return func(x int) int {
		sum += x;
		return sum;	
	}
}

func main() {

}

```

Go does not have classes. Instead places methods on type structs

```

```


Methods of array
append - add data to array
make - creates a slice as if it were an array
copy - duplicates the slice