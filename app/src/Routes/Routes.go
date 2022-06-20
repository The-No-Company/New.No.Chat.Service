package Routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nochat/app/src/Controller"
	"nochat/app/src/Static"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	route.LoadHTMLGlob("public/**")

	for _, value := range Static.GetStyles() {
		route.StaticFile(value[0], value[1])
	}

	group := route.Group("/contact")
	{
		group.GET("upload", Controller.Upload)
		group.GET("search", Controller.Search)
		group.GET("search-all", Controller.SearchAll)
		group.GET("count", Controller.Count)
		group.GET("last", Controller.Last)
	}

	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusAccepted, "index.html", nil)
	})

	return route
}
