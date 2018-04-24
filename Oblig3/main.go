package main
import (
  "net/http"
  "strings"
  "time"
)

func sayHello(w http.ResponseWriter, r *http.Request)
{
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello client" + message
  w.Write([]byte(message))
}


{
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
