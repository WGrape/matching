package strategy

import (
	"fmt"
	"strings"
	"testing"
)

func TestUseStrategy_AutoCreateStrategy(t *testing.T) {
	strategy := UseStrategy{}
	err := strategy.AutoCreateStrategy("../../config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}

	fmt.Println(strategy.Table)
}

func TestUseStrategy_ComputeUser(t *testing.T) {
	strategy := UseStrategy{}
	err := strategy.AutoCreateStrategy("../../config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
		return
	}

	user := User{
		UserId: "100",
		Score:  1.0,
		Gender: 1,
		Age:    "90",
		City:   "ANU",
		Status: 1,
	}
	implodePropertiesString, combinationList, matchedCombinationList := strategy.ComputeUser(user)
	if implodePropertiesString != "gender=1;age=90;city=ANU" {
		t.Fail()
		return
	}

	if strings.Join(combinationList, "/") != "gender=1;age=90;city=ANU/gender=1;age=90/gender=1;city=ANU/age=90;city=ANU/gender=1/age=90/city=ANU" {
		t.Fail()
		return
	}

	if strings.Join(matchedCombinationList, "/") != "gender=0;age=90;city=ANU/gender=1;age=90;city=ANU/gender=0;city=ANU/gender=1;city=ANU/age=90;city=ANU/gender=0;age=90/gender=1;age=90/gender=0/gender=1/age=90" {
		t.Fail()
		return
	}
}
