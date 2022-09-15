# goladok3

[![Go Reference](https://pkg.go.dev/badge/github.com/masv3971/goladok3.svg)](https://pkg.go.dev/github.com/masv3971/goladok3)

## Installation 
```
go get github.com/masv3971/eduid_amapi
 ```

 ## Example
 ```go
 package main

import (
    "github.com/masv3971/eduid_amapi"
)

func main() {
    amapi, err := eduid_amapi.New()
    // handle error
    
    amapi.Users.Update(
    
}

    