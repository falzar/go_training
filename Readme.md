# Go learning

This repository is a compilation of Go exercises in the road of learning

### Tech

* [Go] -  Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. Version 1.9. **Go is not an Object Oriented Language**. Is a statically typed language. 
* [Visual Studio Code] - isual Studio Code is a lightweight but powerful source code editor which runs on your desktop and is available for Windows, macOS and Linux. Version 1.16.

### Installation

[Download Go]. The version used was Apple macOS _Darwing_ version 1.9

+ Installation folder is `/usr/local/go` and `/usr/local/go/bin` is on _PATH_ environment variable
+ You can also add the `bin` folder on .bach_profile
    + open .bach_profile
    + add to the last line `export PATH="/usr/local/go/bin:$PATH"`
    + save
    + Reload with `. ~/.bash_profile`
+ Create ***Workspace*** directory `$HOME/go` with `mkdir $HOME/go` to change the directory, see [Setting GOPATH]

for development on ***VS Code*** over macOS X Sierra:

+ Download insiders version of [Visual Studio Code] so we have newest features.
+ Unzip the folder and then move to Applications folder
+ You can create a symbolic link of the workspace to any folder you need work with `ln -s $HOME/go`

Set up the editor

+ Install Go Extension
+ install ***gocode*** with `go get -v github.com/nsf/gocode`
+ Open a `.go` file with editor, you can see on the bottom side the words `Analysis Tools Missing`. Click over this text and install missing tools.

### Compile, run and test

+ To compile a _.go_ file use `go build` on file's folder
+ To run compiled file `./<name>`
+ To just run use `go run <name>.go` over ***executable packages***
+ To run tests, go to `test` file and run `go test` 

### Install Packages

[Packages Documentation]


### List of exercises

+ **cards/main.go**: Shows Slice definition, subslice, function definition, "instance" calls. Shows how to receive two values return from a function.
+ **cards/deck.go**: Shows an aproximation of Object Oriented on Go. Define a type and functions with a receiver. Contains functions that returns multiple values. Contains for loop use. I/O operations. Type conversion. Use of _strings_ package. Error handling. _OS_ package use. _rand_ package use. One line swap for slice. _time_ package use.
+ **cards/deck_test.go**: Testing exercises with go.
+ **even_detector/main.go**: Test of slice and for loop.
+ **structs/main.go**: Creation and example of structs. Information about receiver functions. Pointers.
+ **interfaces_chatbots/main.go**: Interfaces basics.
+ **interfaces\_http/main.go**: Use of _net_, _http_ packages use and implementation/use of the _Writer_ and _Reader_ interfaces.
+ **interfaces\_printfile/main.go**: Use of _os_ and _io_ packages use and use of the _Writer_ and _Reader_ interfaces.
+ **web_status_checker/main.go**: Principles of paralell programming, "Routines", "Chanels", Sleep function, function literals.
+ **go_web/templates/main.go**: Shows how to read and send messages and data structures to templates. Templates functions.
+ **go\_web/templates/date\_formatting/main.go**: Show how to use properly date values on templates and pass functions from code to templates. Package _time_ use.
+ **go\_web/templates/pipeline/main.go**: Show how to use pipeline to chain functions on templates.
+ **go\_web/templates/nested/main.go**: Show how to use templates inside other templates.
+ **go\_web/templates/methods/main.go**: Show how to use struct methods inside templates.

### TODO

### Useful links

+ Udemy course -> [Go: The Complete Developer's Guide (Golang)]
+ interactive go -> [The Go Playground]
+ Todd McLeod go training -> [GoesToEleven]
+ Todd McLeod go web training -> [GoesToEleven web]
+ The template package documentation. Also important to see predefined global function built in  -> [Template Package Documentation]


License
----

MIT License

Copyright (c) 2017 John Tangarife

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [Go]: <https://golang.org/>
   [Download Go]: <https://golang.org/dl/>
   [Visual Studio Code]: <https://code.visualstudio.com/>
   [Setting GOPATH]: <https://github.com/golang/go/wiki/Setting-GOPATH>
   [Packages Documentation]: <https://golang.org/pkg/>
   [Go: The Complete Developer's Guide (Golang)]: <https://www.udemy.com/go-the-complete-developers-guide/learn/v4/overview>
   [The Go Playground]: <https://play.golang.org>
   [GoesToEleven]:<https://github.com/GoesToEleven/GolangTraining>
   [GoesToEleven web]:<https://github.com/GoesToEleven/golang-web-dev>
   [Template Package Documentation]: <https://golang.org/pkg/text/template/>