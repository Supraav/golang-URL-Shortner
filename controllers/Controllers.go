package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"treeleaf/models"

	"github.com/alextanhongpin/base62"
	"github.com/gin-gonic/gin"
)

func GetURL(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")
	longURL,ok:=models.URLs[shortCode]

	if !ok{
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Short code not found"})
	}
	ctx.Redirect(http.StatusMovedPermanently, longURL)
}

func CreateShortURL(ctx *gin.Context){
	var input models.InputSchema
	
	ctx.BindJSON(&input)


	if input.URL==""{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	err:=isValidURL(input.URL)
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	
	// to base62
	encoded:=base62.Decode(input.URL)
	shortCode:=strconv.Itoa(int(encoded))

	models.URLs[shortCode] = input.URL
	fmt.Println(models.URLs)

	response := map[string]string{
		"shortURL": ctx.Request.Host+"/"+shortCode, 
	}

	ctx.JSON(http.StatusCreated, response)
}

func isValidURL(inputURL string) error {
    if _,err := url.Parse(inputURL); err != nil {
        return errors.New("invalid URL")
    }

	response, err := http.Get(inputURL)
	if err != nil{
		return errors.New("URL doesnot exist")
	}

	if response.StatusCode!=200{
		return errors.New("URL doesnot exist")
	}
	
	if !(strings.HasPrefix(inputURL, "http://") || strings.HasPrefix(inputURL, "https://")){
		return errors.New("invalid URL input. E.g: https://www.example.com")
	}
	return nil
}
