package main
 
import (
	"github.com/hoisie/web"
	"github.com/hoisie/mustache"
)

const (
	address string = "127.0.0.1:3000"
)

func indexRoute(val string) string {
	return mustache.RenderFile("templates/index.html",map[string]string{"phrase" : val}) 
}

func main() {
	web.Get('/(.*)', indexRoute)
    web.Run(address)
}