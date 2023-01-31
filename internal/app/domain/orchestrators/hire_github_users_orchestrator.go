package orchestrators

import (
	"log"
	"time"

	"github.com/eviccari/multithread-test-go/configs"
	"github.com/eviccari/multithread-test-go/internal/app/domain"
)

var since = 0
var loaded = 0
var stime time.Time

type HireGithubUsersOrchestrator struct {
	guService domain.GithubUserService
	gtService domain.GreatestUserService
}

func NewHireGithubUsersOrchestrator(guService domain.GithubUserService, gtService domain.GreatestUserService) HireGithubUsersOrchestrator {
	return HireGithubUsersOrchestrator{
		guService: guService,
		gtService: gtService,
	}
}

func (o HireGithubUsersOrchestrator) Execute() (errorsList []error) {
	stime = time.Now()
	if err := o.gtService.SetEmpty(); err != nil {
		errorsList = append(errorsList, err)
		return
	}

	log.Printf("truncate table -> total time: %v", time.Since(stime))

	for {
		//		stime = time.Now()
		gusers, err := o.guService.Get(getPageSize(), since)

		if err != nil {
			errorsList = append(errorsList, err)
			break
		}

		if len(gusers) == 0 {
			break
		}

		for _, guDTO := range gusers {
			gtDTO := domain.BuildGreatestUserDTOFromGithubUserDTO(guDTO)
			errorsList = o.gtService.Create(gtDTO)
			if len(errorsList) > 0 {
				return
			}
		}

		since = gusers[len(gusers)-1].ID

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
