package strategy

import (
	"fmt"
	"github.com/WGrape/golib/permutation"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

// Config the struct of config
type Config struct {
	Strategy map[string]interface{} `yaml:"strategy"`
	App      map[string]string      `yaml:"app"`
}

// UseStrategy the struct of strategy.
type UseStrategy struct {
	Rules map[string][]string
}

// User the struct of user.
type User struct {
	UserId string  // the id of user
	Score  float32 // the score of heap element

	// the user basic properties.
	Gender uint8
	Age    string
	City   string
	Status uint8 // user status, 0:not exists, 1:exists

	// the features for matching.
	ImplodePropertiesString string
	KeyList                 []string
	MatchKeyList            []string
}

// DefaultUserKey the default key
const DefaultUserKey = "default"

// AutoCreateStrategy create the strategy automatically
func (strategy *UseStrategy) AutoCreateStrategy(configFile string) error {
	// read config file
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	// parse the config file
	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return err
	}

	strategy.Rules = make(map[string][]string)
	for _, item1 := range config.Strategy["rules"].([]interface{}) {
		for key, item2 := range item1.(map[interface{}]interface{}) {
			for _, item3 := range item2.([]interface{}) {
				strategy.Rules[key.(string)] = append(strategy.Rules[key.(string)], item3.(string))
			}
		}
	}
	strategy.Rules["default"] = []string{"default"}
	return nil
}

// ComputeUser compute the user strategy data through properties
// return string: implodePropertiesString, []string: combinationList, []string: matchedCombinationList
func (strategy *UseStrategy) ComputeUser(user User) (string, []string, []string) {
	// implode the property to string
	propertyList := strategy.getPropertyList(user)
	implodePropertiesString := strings.Join(propertyList, ";")

	// get the combination list of properties
	combinationList := strategy.getCombinationList(propertyList)

	// get the matched combination list of properties
	matchedCombinationList := strategy.getMatchedCombinationList(combinationList, user)

	return implodePropertiesString, combinationList, matchedCombinationList
}

// getPropertyList get the property list of user
func (strategy *UseStrategy) getPropertyList(user User) []string {
	var propertyList []string

	// property of gender
	genderProperty := fmt.Sprintf("gender=%d", user.Gender)
	propertyList = append(propertyList, genderProperty)

	// property of age
	ageProperty := ""
	if user.Age != "" {
		ageProperty = fmt.Sprintf("age=%s", user.Age)
	}
	if ageProperty != "" {
		propertyList = append(propertyList, ageProperty)
	}

	// property of city
	cityProperty := ""
	if user.City != "" {
		cityProperty = fmt.Sprintf("city=%s", user.City)
	}
	if cityProperty != "" {
		propertyList = append(propertyList, cityProperty)
	}

	return propertyList
}

// getCombinationList get the combination list of properties
func (strategy *UseStrategy) getCombinationList(propertyList []string) []string {
	return permutation.GetCombinationsWithImplode(propertyList, ";")
}

// getMatchedCombinationList get the matched combination list of properties.
func (strategy *UseStrategy) getMatchedCombinationList(combinationList []string, user User) []string {
	var matchedCombinationList []string
	for _, key := range combinationList {
		if strings.Contains(key, "city=") {
			continue
		}
		if _, ok := strategy.Rules[key]; !ok {
			continue
		}
		matchedCombinationList = append(matchedCombinationList, strategy.Rules[key]...)
	}

	var newMatchedCombinationList []string
	for _, matchedKey := range matchedCombinationList {
		if user.City != "" {
			if matchedKey != "default" {
				newMatchedCombinationList = append(newMatchedCombinationList, fmt.Sprintf("%s;city=%s", matchedKey, user.City))
			}
		}
	}
	newMatchedCombinationList = append(newMatchedCombinationList, matchedCombinationList...)
	return newMatchedCombinationList
}
