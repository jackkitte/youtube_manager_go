package routes

import (
	"github.com/jackkitte/youtube_manager_go/middlewares"
	"github.com/jackkitte/youtube_manager_go/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}
	fg := g.Group("/favorite", middlewares.FirebaseGuard())
	{
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
	}
}
