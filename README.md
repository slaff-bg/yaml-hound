# YAML Hound #
The package covers uncommon needs when working with YAML files.

```go
yh := yamlhound.YAMLTracer{}
if err := yamlReader("path/to/your/test.yaml", &yh.UnmYAML); err != nil {
    panic(err.Error())
}
yh.Footprints = []string{"version"}

fmt.Println("That's what you're looking for:", yh.Caught)
```


[![MIT licensed][shield-license]](#)
[![Top Language][top-language]](#)
[![Go Version][go-version]](#)
[![Go Reference](https://pkg.go.dev/badge/github.com/slaff-bg/yaml-hound.svg)](https://pkg.go.dev/github.com/slaff-bg/yaml-hound)


## What is it? ##

Package working with unknown (dynamic generated) YAML structures.


## Why do we need it? ##

Sometimes it is necessary to handle frequently modified YAML structures. This is common in a rapidly growing system where the configuration YAML file structure changes periodically. These changes lead to struct-types changes in our Go code. Therefore, this package offers an interim solution until the structure finalization of the YAML file.


## Table of Contents ##

* [Features](#features)
* [Requirements](#requirements)
* [Usage](#usage)
* [Contributing](#contributing)
* [License](#license)


## Features ##

- [X] Retrieves a value by key or series of known keys from an unknown dynamic YAML structure.
- [x] Retrieve the first-level keys.
- [x] Retrieve the next-level subkeys of a given key.
- [x] Retrieve the next-level subkeys of a given subset of keys in exact order.
- [ ] etc.


## Requirements/Dependencies ##

- [Golang](https://go.dev/dl/) version go1.20
- [gopkg.in/yaml](https://gopkg.in/yaml.v3) version v3


>__NOTE__
>
> The package should also work with version go1.19, but it was developed under version go1.20, and we still have no tests with the previous one.

>__NOTE__
>
> To get the expected results, you should use V.3 of the
> [gopkg.in/yaml](https://gopkg.in/yaml.v3) package to unmarshalling the XML
> file. Do not use version V.2 or earlier of this package.
>
> Check the [tests](https://github.com/slaff-bg/yaml-hound/blob/main/yaml-hound_test.go) for more technical details.


## Usage ##

```yaml
# file.yaml

# Sample scenario 1
version: "Sample scenario 1 (first level)"

# Sample scenario 2
config:
  herbivores: deer
  predators:
    amphibians: "Saltwater crocodile"
    terrestrial:
      - lion
      - tiger
      - "polar bear"
      - wolf
    flying:
      eagles: "Bald eagle" # :)
      vultures: "Bearded vulture" # :)
      owls: ["Bare-legged owl", "Barn owl", "Snowy owl"]
  omnivores:
    chickens: "Brown Leghorn"
    pigs:
      European: [4, 7]
      Asians: [7, 47]
      American: [47, 84]

# Sample scenario 3
food:
  fast-food:
    chickens: ["KFC", "Chick-fil-A", "Popeyes"]
  desserts:
    chickens: "Sex in a Pan"
  herbivores: cow
  omnivores:
    pigs:
      grilled: n/a
      in-oven: n/a

# Sample scenario 4
vegetables:
  - tomato
  - cucumber

```

```go
// base example
package main

import (
	"fmt"
	"os"

	yamlhound "github.com/slaff-bg/yaml-hound"
	yaml "gopkg.in/yaml.v3"
)

func main() {
  // instance of YAMLTracer{}
  yh := yamlhound.YAMLTracer{}

  // YAMLTracer.UnmYAML: unmarshalled YAML file
  if err := yamlReader("path/to/your/file.yaml", &yh.UnmYAML); err != nil {
    panic(err.Error())
  }

  // Set YAMLTracer.Footprints: the search key or
  // strict sequence of search keys.
  // 
  // e.g. 1.1.
  // 
  // In this case, it looks for a first-level YAML key.
  yh.Footprints = []string{"version"}

  // Performs the function of finding the value by the
  // defined key (or sequence of keys).
  if err := yh.FootprintSniffer(); err != nil {
    panic(err.Error())
  }

  // Lets us know if a value was found or not.
  if yh.Found {
    fmt.Println("That's what you're looking for:", yh.Caught)
  }
}

// yamlReader read and parse YAML file.
func yamlReader(f string, c *map[string]interface{}) error {
	yamlFile, err := os.ReadFile(f)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal([]byte(yamlFile), &c); err != nil {
		return err
	}

	return nil
}
```


```go
  // e.g. 1.2.
  // 
  // In this case, it looks for a third-level YAML key.
  // The keys sequence provided is part of the YAML tree 
  // nd indicates that a key belonging to a specific part
  // of the tree sequence is being sought.
  yh.Footprints = []string{"omnivores", "chickens"}
  _ := yh.FootprintSniffer()
  fmt.Println(yh.Caught) // prints: Brown Leghorn
```

```go
  // e.g. 1.3.
  // 
  // Similar to the previous example, except the key
  // is part of a different sequence.
  yh.Footprints = []string{"fast-food", "chickens"}
  _ := yh.FootprintSniffer()
  fmt.Println(yh.Caught) // prints: ["KFC", "Chick-fil-A", "Popeyes"]
```

```go
  // e.g. 1.4.
  // 
  // Prints <nil> because no such key is found
  // ... OR ...
  // the key is part of a structure (with no specific value).
  yh.Footprints = []string{"food"}
  _ := yh.FootprintSniffer()
  fmt.Println(yh.Caught) // prints: <nil>
```

```go
  // e.g. 1.5.
  // 
  // This test case covers the scenario where we have a
  // specified key (or sequence of keys) that is contained
  // twice in the YAML file.
  // Since a map (map[string]interface{}) is handled,
  // the traversal sequence cannot be predicted during
  // an iteration.
  yh.Footprints = []string{"herbivores"}
  _ := yh.FootprintSniffer()
  fmt.Println(yh.Caught) // prints: "deer" ... OR ... "cow"
```

```go
  // KeysOfKey retrieves the first-level keys of the first
  // found match against the provided key. If more than
  // one key is provided, the function looks for an exact match
  // of the sequence in the YAML tree.
  //
  // e.g. 2.1.
  //
  // If no YAMLTracer.Footprints are provided will collect
  // the first-level keys of the YAML structure.
  _ = yt.KeysOfKey()
  fmt.Println(yh.Caught) // prints: ["vegetables", "version", "config", "food"]
```

```go
  // e.g. 2.2.
  //
  // First-level key WITH sub-level.
  yt.Footprints = []string{"food"}
  _ = yt.KeysOfKey()
  fmt.Println(yh.Caught) // prints: ["fast-food", "desserts", "herbivores", "omnivores"]
```

```go
  // e.g. 2.3.
  //
  // Second-level key WITH sub-level.
  yt.Footprints = []string{"predators"}
  _ = yt.KeysOfKey()
  fmt.Println(yh.Caught) // prints: ["amphibians", "terrestrial", "flying"]
```

```go
  // e.g. 2.4.
  //
  // Second-level key WITH sub-level by using a strict key sequence.
  yt.Footprints = []string{"config", "predators"}
  _ = yt.KeysOfKey()
  fmt.Println(yh.Caught) // prints: ["amphibians", "terrestrial", "flying"]
```

```go
  // e.g. 2.5.
  //
  // Third-level key WITH sub-level by using a strict key sequence.
  yt.Footprints = []string{"predators", "flying"}
  _ = yt.KeysOfKey()
  fmt.Println(yh.Caught) // prints: ["eagles", "vultures", "owls"]
```

```go
  // e.g. 2.6.
  //
  // This test case covers the scenario where we have a specified key (or
  // sequence of keys) that is contained twice in the YAML file. Since a
  // map (map[string]interface{}) is handled, the traversal sequence
  // cannot be predicted during an iteration.
  yt.Footprints = []string{"pigs"}
  _ = yt.KeysOfKey()
  fmt.Println(yh.Caught)
  // prints: ["grilled", "in-oven"]
  // ... OR ...
  // ["European", "Asians", "American"]
```



## Contributing ##

To contribute to YAML Hound, clone this repo locally and commit your code on a separate branch. Please write unit tests for your code, and run the linter before opening a pull-request:


## License ##

[License rights and limitations (MIT).](https://github.com/slaff-bg/yaml-hound/blob/main/LICENSE)


[shield-license]: https://img.shields.io/github/license/slaff-bg/yaml-hound?style=flat&logo=github
[top-language]: https://img.shields.io/github/languages/top/slaff-bg/yaml-hound?style=flat&logo=github
[go-version]: https://img.shields.io/github/go-mod/go-version/slaff-bg/yaml-hound?style=flat&logo=github
[tag-latest]: https://img.shields.io/github/v/tag/slaff-bg/yaml-hound?style=flat&logo=github
[repo-syze]: https://img.shields.io/github/repo-size/slaff-bg/yaml-hound?style=flat&logo=github
[shield-coverage]: https://img.shields.io/badge/coverage-95.0%25-brightgreen.svg?style=flat&logo=github
