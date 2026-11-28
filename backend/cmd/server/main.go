package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"student-management-system/internal/config"
	hlr "student-management-system/internal/handlers"
	"student-management-system/internal/repository"
	"syscall"
	"time"

	"os/signal"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("sms project starting")

	cfg := config.Load()

	r := gin.Default()

	r.GET("/health", hlr.NewHealthHandler().Get)

	hlo := hlr.NewHelloHandler()
	demo := hlr.NewDemoHandler()
	r.GET("/api/hello", hlo.Get)

	// studentHandler /api/students:= hlr.NewStudentHandler()
	var studentHandler *hlr.StudentHandler
	studentHandler = hlr.NewStudentHandler(repository.NewStudentRepository())
	r.POST("/api/students", studentHandler.Create)
	r.GET("/api/students", studentHandler.List)
	r.GET("/api/students/:id", studentHandler.GetByID)
	r.POST("/api/students/import", studentHandler.Import)

	r.GET("/api/demo/slow", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		demo.Slow(c)
	})

	addr := ":" + cfg.Port
	srv := &http.Server{Addr: addr, Handler: r}

	go func() {
		slog.Info("starting server", "addr", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server stopped", "err", err)
			os.Exit(1)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown", "err", err)
		os.Exit(1)
	}
	slog.Info("bye")

}
