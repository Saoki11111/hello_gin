package main
import (
  "log"
  "github.com/gin-gonic/gin"
  "net/http"
)
func main() {
  router := gin.Default()

  router.StaticFS("/static", http.Dir("static"))

  router.GET("/", func(ctx *gin.Context) {
    ctx.Redirect(302, "/static/index.html")
  })

  router.GET("/hello", func(ctx *gin.Context) {
    name := ctx.Query("name")
    ctx.Header("Context-Type", "text/html; charset=UTF-8")
    ctx.String(200, "<h1>Hello, " + name + "</h1>")
  })

  err := router.Run("127.0.0.1:8888")
  if err != nil {
    log.Fatal("fatal server!", err)
  }
}
