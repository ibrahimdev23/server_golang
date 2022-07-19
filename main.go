package main

import (
	//"go_server/router"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func TimeNow(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()["zone"]
	location, _:= time.LoadLocation(string(params[1]))

	fmt.Fprint(w, "time is", time.Now().In(location))
	

}


func handler(w http.ResponseWriter, _*http.Request) {
	fmt.Fprint(w, "Hellow World Thanks for connecting")
}


func handler1(w http.ResponseWriter, _*http.Request ) {
	fmt.Fprint(w, "Handler 1")
}

func Now(w http.ResponseWriter,  _*http.Request){
	now:=time.Now()
	p:=make(map[string]string)
	p["now"]=now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)


}




func main() {

	// fmt.Println("8000")
	r:=mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/please", handler).Methods("GET")

	r.HandleFunc("/us", handler1).Methods("GET")
	r.HandleFunc("/now", Now).Methods("GET")

	r.HandleFunc("/localtime", TimeNow).Methods("GET")
		http.Handle("/", r)

	http.ListenAndServe(":8000", nil)





}