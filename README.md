# vanity

Go vanity import urls

## Usage

```go
if strings.HasPrefix(path, "/bsky") {
  util.ContentType(w, "text/html")
  fmt.Fprint(w, vanity.Git(strings.Replace(path, "/bsky", "", 1), "jordanreger.com/bsky", "git.sr.ht/~jordanreger/bsky"))
  return
}
```
