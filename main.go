package main 
import "net/http"

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello to my go webserver"))
	})
	http.ListenAndServe(":9090", nil)
}