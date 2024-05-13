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

  vanity.Serve(mux)

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
