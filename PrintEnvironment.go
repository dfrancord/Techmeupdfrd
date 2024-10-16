package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type envionmentK8s struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	router := gin.Default()
	router.GET("/environment", getEnvironments)
	router.Run()
}

func getEnvironments(c *gin.Context) {
	var envionmentK8ss []envionmentK8s
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envionmentK8ss = append(envionmentK8ss, envionmentK8s{Name: pair[0], Value: os.Getenv(pair[0])})
	}
	c.IndentedJSON(http.StatusOK, envionmentK8ss)

}
