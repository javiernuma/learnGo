package main

import (
	"fmt"
	"net/http"
	"log"
	
)


func main() {
	
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola Mundo.")
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
	
	

	/*http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("nombre", "valor del header")
		//http.Redirect(w,r, "/dos", http.StatusMovedPermanently)
		fmt.Println("El metodo es + ", r.Method)

	/*	switch r.Method {
		case "GET":
			fmt.Fprintf(w, "hola mundo con el metodo get")
		case "POST":
			fmt.Fprintf(w, "hola mundo con el metodo post")
		case "PUT":
			fmt.Fprintf(w, "hola mundo con el metodo put")
		case "DELETE":
			fmt.Fprintf(w, "hola mundo con el metodo delete")
		default:
			http.Error(w, "Metodo no valido", http.StatusBadRequest)
		
		}*/

		// OBtener datos
		/*fmt.Println(r.URL.Query())// map
		 name := r.URL.Query().Get("name")
		if len(name) != 0{
			fmt.Println(name)
		}
		fmt.Fprintf(w, "Hola Mundo")*/
		
		// GENERAR UNA QUERY
		/*values := r.URL.Query()
		values.Del("otro")
		values.Add("name","Javier")
		values.Add("course","Go web")
		values.Add("Job","Developer")
		r.URL.RawQuery = values.Encode()
		fmt.Println(r.URL)

		fmt.Println(r.Header) // deuelve un map
		accessToken := r.Header.Get("access_token")
		if len(accessToken) != 0{
			fmt.Println(accessToken)
		}
		r.Header.Set("nombre", "valor")
		fmt.Println(r.Header)

	})
	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request){
		///http.NotFound(w,r)
		http.Error(w, "Este es un error",http.StatusConflict)
	})
	log.Fatal(http.ListenAndServe(":3000",nil))*/

	
}