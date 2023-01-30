package main

import (
	"log"
	"os"

	"github.com/eviccari/multithread-test-go/internal/adapters/repositories"
	"github.com/eviccari/multithread-test-go/internal/app/domain/orchestrators"
	"github.com/eviccari/multithread-test-go/internal/app/domain/services"
)

func main() {
	r := repositories.NewGithubUserMockRepository()
	service := services.NewGithubUserServiceImpl(r)
	o := orchestrators.NewHireGithubUsersOrchestrator(service)
	el := o.Execute()

	if len(el) > 0 {
		log.Fatalf("job end with errors: %v", el)
	}

	os.Exit(0)
}
