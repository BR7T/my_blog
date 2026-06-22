package main

import (
	"github.com/BR7T/Blog/endpoints"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil{
	  return
  }
  router := gin.Default() 
  
  endpoints.WebEndpoint(router)
  endpoints.PostEndpoint(router)

  router.Run()
}