package main

import (
	"fmt"
	"time"
  "net/http"
  "echo"
)
var typeoftime = "2021-08-23 00:00:00"
func hello(c echo.Context) error {
	g,err := Hi("大牙", "zh_CN")
	t1 := time.Now()
	t := t1.Local()
	ts := t.Format(typeoftime)
	g += ", "
	g += ts
	fmt.Printf("%s, %s\n", g, err)
  return c.String(http.StatusOK, g)
}

func main() {
  e := echo.New()

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}