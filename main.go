package main 
import(
	"fmt"
	"log"
	"net/http"
)
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm failed: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successfull.")
	name := r.FormValue("name")
	aaddress := r.FormValue("address")
	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v\n", aaddress)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not support.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}