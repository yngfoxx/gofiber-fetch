# gofiber-fetch

A simplified goFiber http-fast fetch package, inspired by node-fetch / JS fetch.

## Usage example

### GET Request
```
package main

import fetch "github.com/yngfoxx/gofiber-fetch"

func main() {
  
  url := "https://483882fa-9a8f-49bb-a8fb-a9cfe121f4a9.mock.pstmn.io/ping"
  
  res := fetch.Method("GET").FiberFetch(url)
  if res.Error != nil {
    panic(res.Error)
  }
  
  log.Printf("Status: %d", res.Status) // Status: 200
  log.Printf("Data: %s", string(res.Data.([]byte))) // Data: pong
}
```

### POST Request
```
package main

import fetch "github.com/yngfoxx/gofiber-fetch"

func main() {
  // use a POST endpoint, do not use sample url
  url := "https://483882fa-9a8f-49bb-a8fb-a9cfe121f4a9.mock.pstmn.io/ping"
  
  qs := fetch.Method("POST")
  // or with Authorization header
  // qs := fetch.Method("POST").SetAuthorization("Bearer ...")
  
  qs.Header = []string{"Content-Type=application/json"}
  qs.Body = map[string]interface{}{
    "foo": "bar",
  }
  
  res := qs.FiberFetch(url)
  log.Printf("Status: %d", res.Status)
  log.Printf("Data: %s", string(res.Data.([]byte)))
}
```

---
Note: all other possible http method follows the same pattern as GET and POST
