package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Up(c echo.Context) error {
	return c.HTML(http.StatusOK, "The <i> SKings </i> <b> Echo </b> server is <b> Up </b> ")
}

func EchoPost(c echo.Context) error {
	fp, err := c.FormParams()
	fmt.Println(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	fmt.Println(fp)
	j, err := json.Marshal(fp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSONBlob(http.StatusOK, j)
}
func main() {
	e := echo.New()
	e.GET("/", Up)
	e.POST("/", EchoPost)
	e.Logger.Fatal(e.Start(":1323"))
}
