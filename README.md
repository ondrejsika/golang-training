[Ondrej Sika (sika.io)](https://sika.io) | <ondrej@sika.io>

# Golang Training

## Course

## Install Go

- https://go.dev/doc/install

### Mac

```
brew install golang
```

## Setup Go Project

```
mkdir hello-world
```

```
cd hello-world
```

### editorconfig

Create `.editorconfig` file

```
root = true
[*]
indent_style = space
indent_size = 2
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true
end_of_line = lf
max_line_length = off
[Makefile]
indent_style = tab
[*.go]
indent_style = tab
```

Create `.gitignore` file using **slu**

```
slu ft ec --go
```

### go mod init

```
go mod init hello-world
```

or

```
go mod init github.com/ondrejsika/hello-world
```

### main.go

Create `main.go` file

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello World")
}
```

## Thank you! & Questions?

That's it. Do you have any questions? **Let's go for a beer!**

### Ondrej Sika

- email: <ondrej@sika.io>
- web: <https://sika.io>
- twitter: [@ondrejsika](https://twitter.com/ondrejsika)
- linkedin: [/in/ondrejsika/](https://linkedin.com/in/ondrejsika/)
- Newsletter, Slack, Facebook & Linkedin Groups: <https://join.sika.io>

_Do you like the course? Write me recommendation on Twitter (with handle `@ondrejsika`) and LinkedIn (add me [/in/ondrejsika](https://www.linkedin.com/in/ondrejsika/) and I'll send you request for recommendation). **Thanks**._

Wanna to go for a beer or do some work together? Just [book me](https://book-me.sika.io) :)
