package orchestrators

import (
	"log"

	"github.com/eviccari/multithread-test-go/configs"
	"github.com/eviccari/multithread-test-go/internal/app/domain"
)

var since = 0
var loaded = 0

type HireGithubUsersOrchestrator struct {
	s domain.GithubUserService
}

func NewHireGithubUsersOrchestrator(service domain.GithubUserService) HireGithubUsersOrchestrator {
	return HireGithubUsersOrchestrator{
		s: service,
	}
}

func (o HireGithubUsersOrchestrator) Execute() (errorsList []error) {
	for {
		gusers, err := o.s.Get(getPageSize(), since)
		if err != nil {
			errorsList = append(errorsList, err)
			break
		}

		if len(gusers) == 0 {
			break
		}

		since = gusers[len(gusers)-1].ID
		log.Println(since)

		loaded += len(gusers)
		if loaded >= configs.GithubUsersQuantity {
			break
		}
	}

	return
}

func getPageSize() int {
	if loaded > 0 {
		if configs.GithubUsersQuantity-loaded < configs.GithubAPIMaxPageSize {
			return configs.GithubUsersQuantity - loaded
		} else {
			return configs.GithubAPIMaxPageSize
		}
	} else {
		if configs.GithubUsersQuantity <= configs.GithubAPIMaxPageSize {
			return configs.GithubUsersQuantity
		} else {
			return configs.GithubAPIMaxPageSize
		}
	}
}
