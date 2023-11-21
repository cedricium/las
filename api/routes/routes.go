package routes

import (
	api "las_api"
	"las_api/handlers"
	"las_api/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	return &Router{
		gin.Default(),
	}
}

func (r Router) Attach(s api.Store) {
	ah := handlers.AdminHandler{Store: s}
	ph := handlers.PatronHandler{Store: s}
	// bh := BooksHandler{store: s}
	// th := TransactionHandler{store: s}
	uh := handlers.UtilHandler{Store: s}

	api := r.Group("/api/v1")
	api.GET("/healthz", uh.CheckHealth)

	auth := api.Group("/auth")
	{
		auth.POST("/login", ah.Login)
		auth.POST("/register", middleware.JWTAuthMiddleware(), ah.Register)
	}

	admins := api.Group("/admins", middleware.JWTAuthMiddleware())
	{
		admins.GET("/", ah.List)    // get list of librarians/admins
		admins.GET("/:id", ah.Show) // get specific admin
	}

	patrons := api.Group("/patrons", middleware.JWTAuthMiddleware())
	{
		patrons.GET("/", ph.List)      // get list of existing patrons
		patrons.POST("/", ph.Register) // create new patron affiliated with library
		patrons.GET("/:id", ph.Show)   // get specific patron and their recent transactions
		patrons.PUT("/:id", ph.Update) // update patron information; pay outstanding balance
	}

	inventory := api.Group("/inventory", middleware.JWTAuthMiddleware())
	{
		inventory.POST("/")                    // add items to library inventory
		inventory.GET("/")                     // get current inventory (optional: and status of each item)
		inventory.GET("/:id")                  // get inventory item and its recent transactions
		inventory.PUT("/:id")                  // update inventory item
		inventory.POST("/:id/transaction")     // checkout/issue item
		inventory.PUT("/:id/transaction/:tid") // mark item returned
	}
}
