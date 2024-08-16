package common

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run 优雅启停函数
func Run(r *gin.Engine, addr string, srvName string) {
	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 优雅启停
	go func() {
		log.Printf("start server in %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 处理自动关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("shutting down %s...", srvName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("%s forced to shutdown:%s", srvName, err)
	}

	select {
	case <-ctx.Done():
		log.Printf("%stimeout out", srvName)
	}

	log.Printf("%s stop successfully!", srvName)
}
