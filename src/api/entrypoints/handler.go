/*
Package entrypoints implements logic to handle requests.
*/
package entrypoints

import "github.com/gin-gonic/gin"

type Handler interface {
	Handle(c *gin.Context)
}
