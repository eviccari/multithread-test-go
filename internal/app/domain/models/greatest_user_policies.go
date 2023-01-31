package models

import (
	"fmt"

	"github.com/eviccari/multithread-test-go/internal/utils"
)

const greatest_user_policy_tag = "policy"

type GreatestUserPoliciesManager struct {
	gu GreatestUser
}

func NewGreatestUserPoliciesManager(greatestUser GreatestUser) GreatestUserPoliciesManager {
	return GreatestUserPoliciesManager{
		gu: greatestUser,
	}
}

func (gupm GreatestUserPoliciesManager) GetPolicies() (polices []func() error) {
	polices = append(polices, gupm.IDCannotBeEmpty)
	polices = append(polices, gupm.NewEmailCannotBeEmpty)
	polices = append(polices, gupm.LegacyLoginCannotBeEmpty)

	return
}

func (gupm GreatestUserPoliciesManager) Apply() (errorsList []error) {
	for _, p := range gupm.GetPolicies() {
		if err := p(); err != nil {
			errorsList = append(errorsList, err)
		}
	}
	return
}

func (gupm GreatestUserPoliciesManager) IDCannotBeEmpty() (err error) {
	if utils.IsEmptyString(gupm.gu.ID) {
		err = fmt.Errorf("%s:IDCannotBeEmpty failed", greatest_user_policy_tag)
	}
	return
}

func (gupm GreatestUserPoliciesManager) NewEmailCannotBeEmpty() (err error) {
	if utils.IsEmptyString(gupm.gu.NewEmail) {
		err = fmt.Errorf("%s:NewEmailCannotBeEmpty failed", greatest_user_policy_tag)
	}
	return
}

func (gupm GreatestUserPoliciesManager) LegacyLoginCannotBeEmpty() (err error) {
	if utils.IsEmptyString(gupm.gu.NewEmail) {
		err = fmt.Errorf("%s:LegacyLoginCannotBeEmpty failed", greatest_user_policy_tag)
	}
	return
}
