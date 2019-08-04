package main

import ("fmt";"net/http")
func index(w  http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"<h1>Hai Guys</h1>");
}

func main() {
	http.HandleFunc("/",index);
	http.ListenAndServe(":8000",nil);
	fmt.Printf("hello, world\n");
}
