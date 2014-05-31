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

---
TEST: Checkout this branch
---
# build & execute

* Go to the repo path in terminal
* Go build compile the go project
	go the workspace directory
	export GOPATH=$(pwd)
	go build hello

		├── README.md
		├── bin
		├── hello
		├── pkg
		└── src
		    └── hello
		        └── hello.go
* Crete the executable
	go install hello

		├── README.md
		├── bin
		│   └── hello
		├── hello
		├── pkg
		└── src
		    └── hello
		        └── hello.go
* Run the executable
	cd bin
	hello

	or

	export PATH=$PATH:$GOPATH/bin
	hello