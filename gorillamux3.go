package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"

)

type Article struct {
	id int `json : id`
	title string `json : title`
	desc string `json : desc`
	content string `json : content`
}

var Articles []Article

func main(){
	Articles = []Article {
		Article{id :"1", title: "Hello", desc: "Article Description", content: "Article Content"},
        Article{id : "2",title: "Hello 2", desc: "Article Description", content: "Article Content"},}

	}
	handleRequests()

}

func  handleRequest(){
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    // add our new DELETE endpoint here
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Printf(w, "welcome to the Homepage!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}



