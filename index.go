package main
 
import (
	 "github.com/hoisie/web"
)

func main() {
	web.Get("/", func() string {
		return "<h1>Hello Go!</h1>"
	})
	web.Run(":8080")
}