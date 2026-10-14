package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 优雅退出（graceful shutdown）：生产服务的必修课。
//
// 标准三步：
//  1. 监听 SIGINT / SIGTERM（k8s 滚动发布、Ctrl+C 都是发这两种信号）
//  2. 收到信号后调用 http.Server.Shutdown：停止接收新请求，等存量请求处理完
//  3. 给一个超时上限，超时后强制退出，避免坏请求卡死整个发布流程
//
// 体验方式：
//
//	go run ./patterns/gracefulshutdown
//	# 另开终端：curl localhost:8080/slow   （3 秒才返回）
//	# 回到第一个终端按 Ctrl+C，观察日志：收到信号 -> 等待存量请求 -> slow 正常返回 -> 进程退出
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("fast response\n"))
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		// 模拟一个耗时 3 秒的业务处理，用来观察"存量请求被等完"
		time.Sleep(3 * time.Second)
		w.Write([]byte("slow response done\n"))
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		// 生产环境还应配置：ReadTimeout / WriteTimeout / IdleTimeout
	}

	// ListenAndServe 会一直阻塞，所以放到单独的 goroutine 里
	go func() {
		log.Println("server listening on :8080 （curl localhost:8080/fast 或 /slow，Ctrl+C 退出）")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// Shutdown 后 ListenAndServe 返回 ErrServerClosed，属于正常退出，不算错误
			log.Fatalf("listen: %v", err)
		}
	}()

	// 阻塞等待退出信号。channel 缓冲为 1：signal.Notify 不会阻塞在发送上
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("收到退出信号，开始优雅退出...")

	// 给存量请求最多 5 秒完成；生产环境按最慢接口的 P99 设置
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("等待超时，强制退出: %v", err)
	}
	log.Println("server exited gracefully")
}
