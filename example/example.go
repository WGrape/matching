package main

import (
	"fmt"
	"matching/pkg/strategy"
)

func main() {
	st := strategy.UseStrategy{}
	err := st.AutoCreateStrategy("config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := strategy.User{
		UserId: "12345678",
		Score:  0.0,
		Gender: 0,
		Age:    "80",
		City:   "Yon",
		Status: 1,
	}
	implodePropertiesString, combinationList, matchedCombinationList := st.ComputeUser(user)
	fmt.Println(implodePropertiesString, combinationList, matchedCombinationList)
}
