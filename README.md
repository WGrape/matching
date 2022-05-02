# matching

[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](./LICENSE) <img src="https://img.shields.io/badge/language-go-blue.svg">

A general and configurable user matching library based on Go language

## Overview
How to find the target matching users by priority ?

<img width="600" alt="image" src="https://user-images.githubusercontent.com/35942268/166154495-415c04b6-c5af-4194-a250-93d26a757ee6.png">

### 1、Core Matching Logic
Computing user and get the core data for matching.

#### (1) User properties

- gender ：male
- age ：27

#### (2) Property Combinations

- gender=male;age=27
- gender=male
- age=27

#### (3) Matched Property Combinations

- gender=female;age=27
- gender=female
- age=27

### 2、Store Users and Get Matched Users

#### (1) Store Users

<img width="600" alt="image" src="https://user-images.githubusercontent.com/35942268/166155670-7c6c9230-422a-461d-b965-3d1e2962c546.png">

#### (2) Get Matched Users

Get user ``` Matched Property Combinations``` and fetch users from ```userMap``` by key in turn.

## Usage
You can follow the steps below, or use the [example](./example/example.go).

### 1、Configuration

- app
  - ```version``` ：the version number of application
  - ```language``` ：the different agents
- strategy
  - ```rules``` ：the different rules of matching

<details>
<summary>Expand config/config.yaml file</summary>

```yaml
app:
  version: v1.0.0
  language: go1.16.10

strategy:
  rules:
    # If it is a woman, first match the male, then the female
    - gender=0:
        - gender=1
        - gender=0

    # If male, match female first, then male
    - gender=1:
        - gender=0
        - gender=1
    # ... ...
```

</details>

### 2、Create strategy object
Create a strategy object and call ```AutoCreateStrategy()```.

```go
package main
import (
    "matching/pkg/strategy"
)

func main(){
    st := strategy.UseStrategy{}
    err := st.AutoCreateStrategy("config/config.yaml")
    if err != nil {
        fmt.Println(err.Error())
        return
    }	
}
```

### 3、Compute user
You can get ```implodePropertiesString```，```combinationList``` and ```matchedCombinationList``` after calling ```ComputeUser()```.

- ```implodePropertiesString``` ：implode the properties to string
- ```combinationList``` ：get the combination list of properties
- ```matchedCombinationList``` ：get the matched combination list of properties

```go
package main
import (
  "matching/pkg/strategy"
)

func main() {
    user := strategy.User{
      UserId: "12345678",
      Score:  0.0,
      Gender: 0,
      Age:    "80",
      City:   "Yon",
      Status: 1,
    }
    implodePropertiesString, combinationList, matchedCombinationList := st.ComputeUser(user)
}
```

## Contributing
Welcome to use and contribute to this project !

## License
[LICENSE](./LICENSE)