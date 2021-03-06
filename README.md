# hello-world-go

--source : http://golang.org/doc/code.html

A simple hello world program in go. The idea behind this small git repo is to serve as a boiler plate for any further experimentation on go programming. A few things this repo has incorporated:

 	* Definition of a workspace in go
 	* The GOPATH environment variable
 	* A note on go env
 	* Basic example of an executable
 	* Basic example of a library 
 	* Baisc example of a test
 	* Packages
 		* Using standard packages
 		* Creating packages
 		* Using custom packages
 		* Fetching remote packages
 		* Using remote packages

## Workspace
A workspace is a directory where you keep all your go code. In order to make most use of go tool, your should have 3 directories in it.
	1. src
	2. pkg
	3. bin

## GOPATH
The location of the workspace.

## go env
print Go environment information


Example 1 : hello world!
---
# checkout, build & execute

* git checkout example-1-executable
* Go build compile the go project
	* go the workspace directory
	* export GOPATH=$(pwd)
	* go build hello
	```
	├── README.md
	├── bin
	├── hello
	├── pkg
	└── src
	    └── hello
	        └── hello.go
	```

* Create the executable
	* go install hello
	```
	├── README.md
	├── bin
	│   └── hello
	├── hello
	├── pkg
	└── src
	    └── hello
	        └── hello.go
	```

* Run the executable
	* cd bin
	* hello

	or

	* export PATH=$PATH:$GOPATH/bin
	* hello

Example 2 : Create a library
---
# checkout, build & execute

* git checkout example-2-library
* Go build compile the go project
	* go the workspace directory
	* export GOPATH=$(pwd)
	* go build newmath
	```
	.
	├── README.md
		├── bin
		│   └── hello
		├── pkg
		└── src
		    ├── hello
		    │   └── hello.go
		    └── newmath
		        └── sqrt.go
	```

* Install the package
	* go install newmath
	```
	.
	├── README.md
	├── bin
	│   └── hello
	├── pkg
	│   └── darwin_amd64
	│       └── newmath.a
	└── src
	    ├── hello
	    │   └── hello.go
	    └── newmath
	        └── sqrt.go
	```

Example 3 : use the custom library method from the package we created in example2 
---
# checkout, build & execute

* git checkout example-3-use-custom-package
* Go build compile the go project
	* go the workspace directory
	* export GOPATH=$(pwd)
	* go build hello
	```
	.
	├── README.md
	├── bin
	│   └── hello
	├── hello
	├── pkg
	│   └── darwin_amd64
	│       └── newmath.a
	└── src
	    ├── hello
	    │   └── hello.go
	    └── newmath
	        └── sqrt.go
	```

* Recreate the executable
	* go install newmath
	```
	.
	├── README.md
	├── bin
	│   └── hello
	├── hello
	├── pkg
	│   └── darwin_amd64
	│       └── newmath.a
	└── src
	    ├── hello
	    │   └── hello.go
	    └── newmath
	        └── sqrt.go
	```
* Run the executable
	* bin/hello
	```
	Hello world!
	Sqrt(2) = 1.414213562373095
	```

Example 4 : Testing packages
---
# checkout, build & execute

* git checkout example-4-testing-packages
* Build the newmath package
	* go the workspace directory
	* go build newmath
	```
	.
	├── README.md
	├── bin
	│   └── hello
	├── pkg
	│   └── darwin_amd64
	│       └── newmath.a
	└── src
	    ├── hello
	    │   └── hello.go
	    └── newmath
	        ├── sqrt.go
	        └── sqrt_test.go
	```

* Test the package
	* go test newmath
	```
	ok  	newmath	0.032s
	```

Example 5 : Using remote packages
---
# checkout, build & execute

* git checkout example-5-remote-packages
* Build the newmath package
	* go the workspace directory
	* go build hello
	```
	src/hello/hello.go:8:2: cannot find package "code.google.com/p/go.example/newmath" in any of:
	/usr/local/go/src/pkg/code.google.com/p/go.example/newmath (from $GOROOT)
	/Users/mnegi/git/my-git-examples/hello-world-go/src/code.google.com/p/go.example/newmath (from $GOPATH)
	```

	```
	.
	├── README.md
	├── bin
	│   └── hello
	├── pkg
	│   └── darwin_amd64
	│       └── newmath.a
	└── src
	    ├── hello
	    │   └── hello.go
	    └── newmath
	        ├── sqrt.go
	        └── sqrt_test.go
	```
* get the package
	* go get code.google.com/p/go.example/newmath
	```
	.
	├── README.md
	├── bin
	│   └── hello
	├── pkg
	│   └── darwin_amd64
	│       ├── code.google.com
	│       │   └── p
	│       │       └── go.example
	│       │           └── newmath.a
	│       └── newmath.a
	└── src
	    ├── code.google.com
	    │   └── p
	    │       └── go.example
	    │           ├── LICENSE
	    │           ├── PATENTS
	    │           ├── codereview.cfg
	    │           ├── hello
	    │           │   └── hello.go
	    │           └── newmath
	    │               ├── sqrt.go
	    │               └── sqrt_test.go
	    ├── hello
	    │   └── hello.go
	    └── newmath
	        ├── sqrt.go
	        └── sqrt_test.go
	```	

* build and install 
	* go build hello
	* go install hello
	* bin/hello
	```
	Hello world!
	Sqrt(2) = 1.414213562373095
	```

Example 6 : using packages with same name
---
# checkout, build & execute

* git checkout example-6-packages-with-same-name
* Go build compile the go project
	* go the workspace directory
	* export GOPATH=$(pwd)
	* go build hello
* Recreate the executable
	* go install hello
* Run the executable
	* bin/hello
	```
	Hello world!
	Sqrt(2) = 1.414213562373095
	Sqrt(2) = 1.414213562373095
	```
