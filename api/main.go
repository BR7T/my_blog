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
    s, _ := os.Getwd()
    log.Println("Diretório atual:", s)
    
    // Tenta carregar .env, mas não falha se não existir
    _ = godotenv.Load()
    log.Println("✅ Variáveis de ambiente carregadas")
    
    router := gin.Default() 
    
    endpoints.WebEndpoint(router)
    endpoints.PostEndpoint(router)
    
    log.Println("🚀 Iniciando servidor...")
    router.Run()
}