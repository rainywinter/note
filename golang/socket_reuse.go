package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

var server *http.Server
var addr = ":8080"

func main() {
	lnCfg := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) (err error) {
			err = c.Control(func(fd uintptr) {
				err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEPORT, 1)
				if err != nil {
					return
				}
				err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
				if err != nil {
					return
				}
			})
			return
		},
		KeepAlive: 0,
	}
	ln, err := lnCfg.Listen(context.Background(), "tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		var r = mux.NewRouter()
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.RemoteAddr)
			_, _ = w.Write([]byte("hello."))
		})

		log.Println("server http ", addr, os.Getpid())
		server = &http.Server{Handler: r}
		log.Fatal(server.Serve(ln))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	log.Println("shutdown: ", server.Shutdown(ctx))
	log.Println("exit ", os.Getpid())
	os.Exit(0)
}
