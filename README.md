# Go Faker
A simple library to generate fake data for testing purposes in Go.

## Installation
```bash
go get github.com/rastogiji/go-faker
```

## Simple Usage
```go
import "github.com/rastogiji/go-faker"

f := gofaker.New()

f.FirstName()
f.LastName()
f.FullName()
f.Email()
f.Phone()
f.Country()
f.Profession()
```

## Struct Usage
```go
import "github.com/rastogiji/go-faker

type Person struct {
    FirstName  string `faker:"first_name"`
    LastName   string `faker:"last_name"`
    Email      string `faker:"email,32"`
    Country    string `faker:"country,123"`
    Phone      string `faker:"phone"`
    Profession string `faker:"profession"`
}

func main(){
	f := gofaker.New()
    p := Person{}
    f.Struct(&p)
    fmt.Println(p)
}
```
Feel free to open an issue or PR for any feature request or bug fix.
```