package main 

import "github.com/gin-gonic/gin"

var router *gin.Engine

client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))

if err != nil {
	panic(err)
}

func main() {
	router := gin.Default()
	router.Run()	
}