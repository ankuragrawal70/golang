package main

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
)

type Student struct{
	Name string `json:"name"`
	Age int32	`json:"age"`
}

//type students []Student


func index(w http.ResponseWriter, r *http.Request){
	//fmt.Fprint(w, "success")
	//x := Student{"a", 24}
	//io.WriteString(w, json.Marshal(x))

	//_, err := json.Marshal(x)
	//if err != nil{
	//	panic(err)
	//}
	students := make([]Student, 2)
	students[0] = Student{"an", 24}
	students[1] = Student{"ab", 25}
	y, err := json.Marshal(students[0])
	if err != nil{
		panic(err)
	}


	//w.Header().Set("content-type", "text/json")
	//w.WriteHeader(http.StatusOK)
	w.Write(y)
	//json.NewEncoder(w).Encode(students)
}



func main(){
	port := 8000
	fmt.Println(port)
	http.HandleFunc("/", index)
	http.ListenAndServe(":"+ strconv.Itoa(port), nil)
	fmt.Println("successfully started server")

}