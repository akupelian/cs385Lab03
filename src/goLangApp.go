package main

import (
  "net/http"
   "os"
)

func messager(w http.ResponseWriter, r *http.Request) {
  message, _ := os.Hostname()
  message = message + "\n"
  w.Write([]byte(message))
}

func main() {
  http.HandleFunc("/", messager)
  if err := http.ListenAndServe(":80", nil); err != nil {
	  panic(err)
  }
}
