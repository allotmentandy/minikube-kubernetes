package hello
import (
   "fmt"
   "log"
   "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintln(w, "It works!")
}
func main() {
   http.HandleFunc("/", handler)
   log.Fatal(http.ListenAndServe(":8080", nil))
}