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
	location, _:= time.LoadLocation(string(params[0]))

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

func getAddress(w http.ResponseWriter, r *http.Request)



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





// Assignment 2 )  Write  a go program to get the client IP and reflect/echo
//back to client page ,after this add feature to server so that server could block to a particular client with a message

// : func ReadUserIP(r *http.Request) string {
//     IPAddress := r.Header.Get("X-Real-Ip")
//     if IPAddress == "" {
//         IPAddress = r.Header.Get("X-Forwarded-For")
//     }
//     if IPAddress == "" {
//         IPAddress = r.RemoteAddr
//     }
//     return IPAddress
// }
// X-Real-Ip - fetches first true IP (if the requests sits behind multiple NAT sources/load balancer)

// X-Forwarded-For - if for some reason X-Real-Ip is blank and does not return response, get from X-Forwarded-For

// Remote Address - last resort (usually won't be reliable as this might be the last ip or if it is a naked http request to server ie no load balancer)

// add feature to store all past access like date and time to access to the server and display a message that  clientIP:xxxxxxxxxxxx 
//  your last access to server is  Date  :  xxxx/xx/xxxx time :xxxxxxx