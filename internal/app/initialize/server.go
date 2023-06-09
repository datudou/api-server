package initialize

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/retail-ai-test/config"
	"github.com/retail-ai-test/internal/api"
	db "github.com/retail-ai-test/internal/pkg/database"
	"github.com/retail-ai-test/internal/pkg/repo"
	"github.com/retail-ai-test/internal/pkg/service"
	"go.uber.org/zap"
)

func InitServer() {
	// init repo
	conn := db.GetConn()
	repos := repo.NewRepositories(conn)
	// init service
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	// init router
	r := gin.Default()
	api.NewHandler(&api.Config{
		R:        r,
		Services: services,
	})

	srv := &http.Server{
		Addr:    config.ServerConf.Server.Addr,
		Handler: r,
	}

	go func() {
		// service connections
		zap.S().Infof("Server listen on: %s", config.ServerConf.Server.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Shutdown Server ...")

	timeout := config.ServerConf.Server.ShutdownTimeout

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of ${timeout} seconds.
	<-ctx.Done()
	zap.S().Infof("timeout of %d seconds.", timeout)
}
