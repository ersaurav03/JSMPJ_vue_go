package main
import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/",handler).Methods("GET")
	r.HandleFunc("/about",aboutHandler).Methods("GET")
	r.HandleFunc("/contact",contactHandler).Methods("GET")
	r.HandleFunc("/login",loginHandler).Methods("GET")

	http.Handle("/",r)
	http.ListenAndServe(":8080",nil)
}
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Home")
}
func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"About")
}
func contactHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Contact")
}
func loginHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Login")
}

