package orchestrators

import (
	"log"
	"sync"
	"time"

	"github.com/eviccari/multithread-test-go/configs"
	"github.com/eviccari/multithread-test-go/internal/app/domain"
	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
)

var mt_since = 0
var mt_loaded = 0
var mt_stime time.Time

type HireGithubUsersMTOrchestrator struct {
	guService domain.GithubUserService
	gtService domain.GreatestUserService
}

func NewHireGithubUsersMTOrchestrator(guService domain.GithubUserService, gtService domain.GreatestUserService) HireGithubUsersMTOrchestrator {
	return HireGithubUsersMTOrchestrator{
		guService: guService,
		gtService: gtService,
	}
}

func (o HireGithubUsersMTOrchestrator) Execute() (errorsList []error) {
	mt_stime = time.Now()
	if err := o.gtService.SetEmpty(); err != nil {
		errorsList = append(errorsList, err)
		return
	}

	log.Printf("truncate table -> total time: %v", time.Since(mt_stime))

	for {
		gusers, err := o.guService.Get(mtGetPageSize(), mt_since)

		if err != nil {
			errorsList = append(errorsList, err)
			break
		}

		if len(gusers) == 0 {
			break
		}

		mt_since = gusers[len(gusers)-1].ID

		for len(gusers) > 0 {
			chunk := getMultiThreadChunk(gusers)

			wg := sync.WaitGroup{}
			c := make(chan []error, len(chunk))

			mu := sync.Mutex{}

			for _, gUser := range chunk {
				wg.Add(1)
				go func(guser dtos.GithubUserDTO) {
					defer wg.Done()

					el := o.gtService.Create(domain.BuildGreatestUserDTOFromGithubUserDTO(guser))
					if len(el) == 0 {
						mu.Lock()
						mt_loaded++
						mu.Unlock()
					}
					c <- el
				}(gUser)
			}

			wg.Wait()
			close(c)

			for el := range c {
				if len(el) > 0 {
					errorsList = el
					return
				}
			}

			if mt_loaded >= configs.GithubUsersQuantity {
				return
			}

			gusers = removeChunkedUsersFromList(len(chunk), gusers)
		}
	}
	return
}

func mtGetPageSize() int {
	if mt_loaded > 0 {
		if configs.GithubUsersQuantity-mt_loaded < configs.GithubAPIMaxPageSize {
			return configs.GithubUsersQuantity - mt_loaded
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

func getMultiThreadChunk(gUsersList []dtos.GithubUserDTO) (chunk []dtos.GithubUserDTO) {
	limit := 0
	if configs.MultiThreadSize > len(gUsersList) {
		limit = len(gUsersList)
	} else {
		limit = configs.MultiThreadSize
	}

	chunk = []dtos.GithubUserDTO{}
	chunk = append(chunk, gUsersList[:limit]...)

	return
}

func removeChunkedUsersFromList(chunkSize int, gusers []dtos.GithubUserDTO) (updatedList []dtos.GithubUserDTO) {
	updatedList = gusers[chunkSize:]
	return
}
