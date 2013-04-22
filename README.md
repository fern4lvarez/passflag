
# passflag

Password flag for Go command line apps. [Documentation online](http://godoc.org/github.com/fern4lvarez/passflag)

**passflag** provides a silence, no echo'ed prompt that will ask for a
password.

## Install

```
go get github.com/fern4lvarez/passflag
```

##Usage

This is an example on how to use `passflag` on a command line app.

```
// GOPATH/src/pf_example/pf_example.go
package main

import (
  "flag"
  "fmt"
  "os"
  "github.com/fern4lvarez/passflag"
)

func main() {
  name := flag.String("name", "", "User name.")
  passflag.Password("p", true, "User Password.")
  flag.Parse()

  pw, _ := passflag.Get()
  fmt.Println("Name = ", *name)
  fmt.Println("Password = ", pw)
}
```

`go get` your example:

```
go get pf_example/pf_example.go
```

And try it on your terminal:

```
$ pf_example -name Foo -p
Password: (no echo)
Name =  Foo
Password = Bar
```


##License
passflag is MIT licensed, see [here](https://github.com/fern4lvarez/passflag/blob/master/LICENSE)
