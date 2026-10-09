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

	"github.com/ArteShow/DoThat/internal/api"
	"github.com/ArteShow/DoThat/internal/database"
	"github.com/ArteShow/DoThat/internal/learning"
	"github.com/ArteShow/DoThat/internal/planner"
	"github.com/ArteShow/DoThat/internal/tasks"
)

func main() {
	// DB init
	db, err := database.Init("migrations", "data/dothat.db")
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	if err = db.DB.Ping(); err != nil {
		panic(err)
	}

	// Manager init
	taskManager, err := tasks.NewTaskManager(db.DB)
	if err != nil {
		panic(err)
	}

	plannerManager := planner.NewPlannerManager(db.DB)
	learningManager := learning.NewLearningManager(db.DB)

	// HTTP init
	router := api.NewRouter(taskManager, plannerManager, learningManager)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           router.Router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverErrors := make(chan error, 1)

	go func() {
		log.Println("Successfully listening on port 8080")
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err = <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}

	case <-ctx.Done():
		log.Println("Shutting down...")
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err = server.Shutdown(shutdownCtx); err != nil {
		panic(err)
	}
}
