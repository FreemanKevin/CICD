package main

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    "github.com/swaggo/swag/example/basic/docs"
    "demo/internal/handlers"
)

// @title Demo API
// @version 1.0
// @description This is a demo API server.
// @host localhost:8080
// @BasePath /api
func main() {
    r := gin.Default()

    // Swagger 文档
    docs.SwaggerInfo.Title = "Demo API"
    docs.SwaggerInfo.Description = "This is a demo API server."
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.BasePath = "/api"
    docs.SwaggerInfo.Schemes = []string{"http"}

    // 重定向根路径到 Swagger UI
    r.GET("/", func(c *gin.Context) {
        c.Redirect(301, "/swagger/index.html")
    })

    // Swagger 路由
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // API 路由组
    api := r.Group("/api")
    {
        api.GET("/items", handlers.GetItems)
        api.GET("/items/:id", handlers.GetItem)
        api.POST("/items", handlers.CreateItem)
        api.PUT("/items/:id", handlers.UpdateItem)
        api.DELETE("/items/:id", handlers.DeleteItem)
    }

    r.Run(":8080")
} 