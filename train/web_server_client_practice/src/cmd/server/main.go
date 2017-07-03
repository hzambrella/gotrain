package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"engine/datastore"
)

var data datastore.Data

func main() {

	d, err := datastore.ParseDataFromFile("file/test")
	if err != nil {
		log.Println(err)
		d = make(datastore.Data)
	}
	data = d

	http.HandleFunc("/test/view", View)
	http.HandleFunc("/test/modify", Modify)
	
	
	addr := ":8080"
	log.Println("Start at " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func View(w http.ResponseWriter, r *http.Request) {
	var result []byte
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()

	key := r.FormValue("key")
	if len(key) == 0 {
		result = data.ToJson()
		return
	}
	val := data.Get(key)
	if len(val) == 0 {
		result = []byte("key not found")
		return
	}
	result = []byte(val)

}

//TODO
func Add(w http.ResponseWriter, r *http.Request) {
	
}


func Modify(w http.ResponseWriter, r *http.Request) {
debug, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(debug))
	}
	var result []byte
	defer func() {
		if _, err := w.Write(result); err != nil {
			log.Fatal(err)
		}
	}()

	/*
	b,err:=httputil.DumpRequest(r,true)
	if err!=nil{
		log.Fatal(err.Error())
	}
	fmt.Println(string(b))
	*/
	r.ParseForm()
	key := r.FormValue("key")
	val := data.Get(key)
	fmt.Println("key:",key)
	if len(val) == 0 {
		result = []byte("key not found")
		return
	}
	value := r.FormValue("data")
	if len(value) == 0 {
		result = []byte("need data value")
		return
	}
	data.Set(key, value)
	if err := datastore.SaveDataToFile("test", data); err != nil {
		result = []byte(err.Error())
		return
	}
}

//TODO
func Delete(w http.ResponseWriter, r *http.Request) {
	
}
