package main

import (
	"log"
	"os"
	"time"

	"github.com/eviccari/multithread-test-go/internal/adapters/infra"
	"github.com/eviccari/multithread-test-go/internal/adapters/repositories"
	"github.com/eviccari/multithread-test-go/internal/app/domain/orchestrators"
	"github.com/eviccari/multithread-test-go/internal/app/domain/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Printf("JOB with process ID %d starting now: %v", os.Getpid(), time.Now().Local().UTC())
	stime := time.Now()
	db, err := infra.GetDB()
	if err != nil {
		log.Fatalf("error on connect to database: %s", err.Error())
	}

	defer infra.CloseDB(db)

	guRepo := repositories.NewGithubUserMockRepository()
	gtRepo := repositories.NewGreatestUserMySQLRepository(db)

	guService := services.NewGithubUserServiceImpl(guRepo)
	gtService := services.NewGreatestUserServiceImpl(gtRepo)

	o := orchestrators.NewHireGithubUsersMTOrchestrator(guService, gtService)

	el := o.Execute()

	if len(el) > 0 {
		log.Fatalf("job end with errors: %v", el)
	}

	log.Printf("JOB finish at: %v", time.Now().UTC())
	log.Printf("total time: %v", time.Since(stime))
	os.Exit(0)
}
