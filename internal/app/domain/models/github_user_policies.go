package models

import (
	"fmt"
	"regexp"

	"github.com/eviccari/multithread-test-go/internal/utils"
)

const pm_policy_tag = "policy"

type GithubUserPoliciesManager struct {
	u GithubUser
}

func NewGithubUserPoliciesManager(u GithubUser) GithubUserPoliciesManager {
	return GithubUserPoliciesManager{
		u: u,
	}
}

func (pm GithubUserPoliciesManager) Apply() (errorsList []error) {
	for _, p := range pm.GetPolicies() {
		if err := p(); err != nil {
			errorsList = append(errorsList, err)
		}
	}
	return
}

func (pm GithubUserPoliciesManager) GetPolicies() (policies []func() error) {
	policies = append(policies, pm.LoginCannotBeEmpty)
	policies = append(policies, pm.IDMustBeGreaterThan0)
	policies = append(policies, pm.UserURLMustBeAValidURL)

	return
}

func (pm GithubUserPoliciesManager) LoginCannotBeEmpty() (err error) {
	if utils.IsEmptyString(pm.u.Login) {
		err = fmt.Errorf("%s:LoginCannotBeEmpty failed", pm_policy_tag)
	}

	return
}

func (pm GithubUserPoliciesManager) IDMustBeGreaterThan0() (err error) {
	if pm.u.ID < 1 {
		err = fmt.Errorf("%s:IDMustBeGreaterThan0 failed", pm_policy_tag)
	}

	return
}

func (pm GithubUserPoliciesManager) UserURLMustBeAValidURL() (err error) {
	pattern := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)

	if matches := pattern.MatchString(pm.u.URL); !matches {
		err = fmt.Errorf("%s:UserURLMustBeAValidURL failed", pm_policy_tag)
	}

	return
}
