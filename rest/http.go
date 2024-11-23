package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/halovina/golang-tutorial/infra/postgres"
	"github.com/halovina/golang-tutorial/model"
	"github.com/halovina/golang-tutorial/router"
)

func Start() error {
	// init database connection
	db := postgres.New()
	// migrate database
	err := postgres.Migrate(
		db.GormDB,
		model.UserData{},
	)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:8500",
		Handler: router.RegisterRouter(),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	errCh := make(chan error)

	log.Println(fmt.Sprintf("starting %s server at %s", "samplegolang", srv.Addr))

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			switch err {
			case http.ErrServerClosed:
				return
			default:
				errCh <- err
			}
		}
	}()

	select {
	case <-stop:
		log.Println("app shutdwon")
		return srv.Shutdown(context.Background())
	case err := <-errCh:
		log.Println(err)
		return err
	}
}
