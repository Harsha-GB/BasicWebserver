package main
import(
	"fmt"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path !="/Hello"{
		http.Error(w, " you are not supposed to be here",http.StatusNotFound)
		return
	
	}
	if r.Method != "GET"{
		http.Error(w, " METHOD IS NOT SUPPORTED",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"Hello")

}
func formHandler(w http.ResponseWriter,r *http.Request){
	if err:=r.ParseForm(); err !=nil{
		fmt.Fprintf(w,"Parseform() err %v",err)
		return
	}
	fmt.Fprintf(w,"Post request successful")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	age:=r.FormValue("age")
	fmt.Fprintf(w, "Name = %s\n",name)
	fmt.Fprintf(w, "Address=%s\n",address)
	fmt.Fprintf(w, "Age=%v\n",age)

}
func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)  //route router
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/Hello",helloHandler)

	fmt.Printf("Starting server at port 8080 \n")
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}

}