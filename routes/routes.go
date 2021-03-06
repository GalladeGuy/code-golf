package routes

import (
	"database/sql"
	"math/rand"
	"net/http"

	"github.com/JRaspass/code-golf/middleware"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var db *sql.DB
var Router = httprouter.New()

func init() {
	var err error

	if db, err = sql.Open("postgres", ""); err != nil {
		panic(err)
	}

	Router.GET("/", middleware.Gzip(home))
	Router.GET("/about", middleware.Gzip(about))
	Router.GET("/assets/:0", asset)
	Router.GET("/callback", callback)
	Router.GET("/favicon.ico", asset)
	Router.GET("/logout", middleware.Gzip(logout))
	Router.GET("/random", random)
	Router.GET("/robots.txt", robots)
	Router.GET("/scores", middleware.Gzip(scores))
	Router.GET("/scores/:0", middleware.Gzip(scores))
	Router.GET("/scores/:0/:1", middleware.Gzip(scores))
	Router.GET("/scores/:0/:1/all", middleware.Gzip(scores))
	Router.GET("/scores/:0/:1/mini", middleware.Gzip(scoresMini))
	Router.GET("/stats", middleware.Gzip(stats))
	Router.GET("/users/:0", middleware.Gzip(user))

	for _, h := range holes {
		Router.GET("/"+h.ID, middleware.Gzip(hole))
	}

	Router.POST("/solution", middleware.Gzip(solution))

	Router.NotFound = http.HandlerFunc(print404)
}

func random(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, holes[rand.Intn(len(holes))].ID, 302)
}

func robots(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
}
