package main

import(
  "os"
  "net/http"
)

func main() {
  logs := make(chan string, 10000)
  go RunLogging(logs)

  handler := routerHandlerFunc(router())
  handler = wrapLogging(handler, logs)
  http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
