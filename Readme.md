# Go learning

This repository is a compilation of Go exercises in the road of learning

### Tech

* [Go] -  Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. Version 1.9.
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
+ You can create s symbolic link of the workspace to any folder you need work with `ln -s $HOME/go`

Set up the editor

+ Install Go Extension
+ install ***gocode*** with `go get -v github.com/nsf/gocode`
+ Open a `.go` file with editor, you can see on the bottom side the words `Analysis Tools Missing`. Click over this text and install missing tools.

### Compile

+ To compile a _.go_ file use `go build` on file's folder
+ To run compiled file `./<name>`
+ To just run use `go run <name>.go` over ***executable packages***

### Install Packages

[Packages Documentation]


### List of exercises

### TODO

### Useful links


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