package main

import (
	"github.com/Rafiur/api-crud-assessment/config"
	"github.com/Rafiur/api-crud-assessment/internal/handler"
	"github.com/Rafiur/api-crud-assessment/internal/repository/repo_postgres"
	"github.com/Rafiur/api-crud-assessment/internal/router"
	"github.com/Rafiur/api-crud-assessment/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	repo := repo_postgres.NewTaskRepository(db)
	svc := usecase.NewTaskService(repo)

	h := handler.NewTaskHandler(*svc)

	r := gin.Default()
	router.SetupRoutes(r, h)

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
