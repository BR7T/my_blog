package main

import (
	"log"
	"os"

	"github.com/BR7T/Blog/endpoints"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
      log.Println("✅ Aplicação iniciada")
    log.Println("Porta:", os.Getenv("PORT"))
    s , _ :=os.Getwd()
    log.Println("Diretório atual:", s)
  err := godotenv.Load()
  if err != nil{
	  return
  }
  router := gin.Default() 
  
  endpoints.WebEndpoint(router)
  endpoints.PostEndpoint(router)

  router.Run()
}