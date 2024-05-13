# vanity

Go vanity import urls

## Usage

`main.go`
```go
package main

import (
  "log"
  "net/http"

  "jordanreger.com/vanity"
)

func main() {
  mux := http.NewServeMux()

  vanity_config, err := os.ReadFile("vanity.json")
  if err != nil {
    panic(err)
  }

  vanity.Serve(mux, string(vanity_config))

  log.Fatal(http.ListenAndServe(":8080", mux))
}
```

`vanity.json`

```json
{
  "host": "example.com",
  "urls": [
    {
      "pkg": "example1",
      "url": "git.sr.ht/example/example1"
    },
    {
      "pkg": "example2",
      "url": "github.com/example/example2"
    }
  ]
}
```

How to test: `go test -v -run Example`

## Contributing

Send patches/bug reports to <~jordanreger/public-inbox@lists.sr.ht>
