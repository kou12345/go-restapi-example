package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kou12345/go-restapi-example/api/internal/controller/system"
	"github.com/kou12345/go-restapi-example/api/internal/controller/user"

	_ "github.com/kou12345/go-restapi-example/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const LATEST = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(LATEST)

	// 死活監視用
	// 可読性のためにブロックを分けている
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	// ユーザー管理機能
	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUsers)
		v1.GET("/:id", userHandler.GetUserById)
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}
