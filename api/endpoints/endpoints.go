package endpoints

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/BR7T/Blog/functions"
	"github.com/BR7T/Blog/structs"
	"github.com/gin-gonic/gin"
)

func CorsMid()gin.HandlerFunc{
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
        c.Header("Access-Control-Allow-Methods", "GET")
        c.Header("Access-Control-Allow-Headers", "Content-Type, X-ADMIN-SECRET")

        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

func Bucket(bucket *structs.TokenBucket) gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Printf("Bucket:\nTokens Restantes:%v\n" , bucket.Qtd)
		if !bucket.Consume(){
			c.JSON(429 , gin.H{
				"error" : "Aguarde e faça a requisição novamente",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func PostEndpoint(router *gin.Engine){

	bAdd := &structs.TokenBucket{Qtd:10 , Max: 10}
	go bAdd.InsertToken()

	bGet := &structs.TokenBucket{Qtd: 10 , Max: 10}
	go bGet.InsertToken()

	router.Use(CorsMid())
	post := router.Group("/post")

	

	router.OPTIONS("/*path", func(c *gin.Context) {
    c.AbortWithStatus(http.StatusNoContent)
})
	
	{
		post.GET("/get" , Bucket(bGet) , func(c *gin.Context) {
			Get(c)
		})
	}
}

func Get(c *gin.Context){
	page := c.Request.URL.Query().Get("page")
	size := c.Request.URL.Query().Get("size")

	idPost := c.Request.URL.Query().Get("id")

	if strings.TrimSpace(idPost) != ""{
		v , err := strconv.Atoi(idPost)
		if err != nil{
			c.JSON(400 , gin.H{
				"error":"Valor de id precisa ser um número inteiro",
			})
			return
		}

		post , err := functions.GetPostIDDatabase(v)
		if err != nil {
			c.JSON(400 , gin.H{
				"error":"Erro ao buscar ID no banco de dados",
			})
			fmt.Println(err)
			return
		}
		c.JSON(200 , gin.H{
			"message":post,
		})
		return
		

	}

	if strings.Trim(page , " ") == " "{
		c.JSON(400 , gin.H{
			"error":"Não foi encontrado valor de página ou tamanho de página",
		})
		return
	}
	if strings.Trim(size , " ") == ""{
		size = "20"
	}
	
	pageInt , err1 := strconv.Atoi(page)
	sizeInt , err2 := strconv.Atoi(size)

	if err1 != nil || err2 != nil || pageInt < 1 || sizeInt < 1 {
		c.JSON(400 , gin.H{
			"error":"Valor de page ou size precisam ser número inteiros e válidos",
		})
		return
	}

	list  , err := functions.GetPostDatabase(pageInt , sizeInt)
	if err != nil{
		fmt.Println(err)

		c.JSON(400 , gin.H{
			"error":err,
		})
		return
	}

	c.JSON(200 , gin.H{
		"message":list,
	})
}