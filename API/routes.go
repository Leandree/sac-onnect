package main
 
import (
    "github.com/julienschmidt/httprouter"
)
 
// NewRouter return all router
func NewRouter() *httprouter.Router {
    router := httprouter.New()
    router.GET("/", Index)
    router.POST("/users", SaveUsers)
    router.POST("/weight", SaveWeight)
    router.PUT("/item", UpdateItem)
    return router
}