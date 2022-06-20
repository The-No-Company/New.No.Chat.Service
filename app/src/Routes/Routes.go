package Routes

import (
	"github.com/gin-gonic/gin"
	"nochat/app/src/Controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("/contact")
	{
		group.GET("upload", Controller.Upload)
		group.GET("search", Controller.Search)
		group.GET("search-all", Controller.SearchAll)
		group.GET("count", Controller.Count)
		group.GET("last", Controller.Last)
	}

	return r
}
