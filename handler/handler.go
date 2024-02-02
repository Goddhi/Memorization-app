package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// Handler struct  holds reuired function for handler to function....
type Handler struct{}


// Config will hold services that will eventually be injected into this
//handler layer on layer initialization
type Config struct {
	R *gin.Engine
}

// Newhandler initializes the handler with required services along with http routes
// Does not return as it deals directly witth the reference to the gin engine
func Newhandler(c *Config) {
	// Create an handler (which will lalter have injejcted services )
	h := &Handler{} // currently has no properties

	//create an  account group endpoint called api
	g := c.R.Group("/api/account")

	g.GET("/", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "space persons",
		})
		
	})
}

// Me handler calls services for getting
// a user's details
func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "it's signup",
	})
}

// Signup handler
func (h *Handler) Signup(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Hello": "It's id a signup",
	})
}

/// Signin Handler 
func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello" : "it's signup",
	})
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "Its tokens",
	})

// 
} 