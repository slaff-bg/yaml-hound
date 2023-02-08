# YAML Hound #
The package covering the need for non-standard approaches when working with YAML files.

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


## What is it? ##

Package working with unknown (dynamic generated) YAML structures.


## Why do we need it? ##

Sometimes it is necessary to handle frequently modified YAML structures. This is common in a rapidly growing system where the configuration file structure changes periodically. These changes lead to struct-types changes in our Go code. Therefore, this package offers an interim solution until the final finalization. of the model YAML file.


## Table of Contents ##

* [Features](#features)
* [Requirements](#requirements)
* [Usage](#usage)
* [Contributing](#contributing)
* [License](#license)


## Features ##

- [X] Retrieves a value by key or series of known keys from an unknown dynamic YAML structure.
- [ ] N-level keys ripper.
- [ ] etc.


## Requirements ##

- [Golang](https://go.dev/dl/) version go1.20
- [gopkg.in/yaml](https://gopkg.in/yaml.v3) version v3

&#x1F4CC; &nbsp; *<sub>Versions reflect the current state of the used
technologies.</sub>*


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

# Sample scenario 3
food:
  fast-food:
    chickens: ["KFC", "Chick-fil-A", "Popeyes"]
  desserts:
    chickens: "Sex in a Pan"
  herbivores: cow

# Sample scenario 4
vegetables:
  - tomato
  - cucumber

```

```go
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

    // set YAMLTracer.Footprints: the search key or strict sequence
    // of search keys
    // 
    // e.g. 1 returns: Sample scenario 1 (first level)
	yh.Footprints = []string{"version"}

    // e.g. 2 returns: Brown Leghorn
	// yh.Footprints = []string{"omnivores", "chickens"}
    // e.g. 3 returns: ["KFC", "Chick-fil-A", "Popeyes"]
	// yh.Footprints = []string{"fast-food", "chickens"}

    // e.g. 4 returns: <nil>
    // Returns <nil> because no such key is found ... OR ...
    // the key is part of a structure (with no specific value).
	// yh.Footprints = []string{"food"}


    // e.g. 5 returns: "deer" ... OR ... "cow"
    // This test case covers the scenario where we have a specified key (or
    // sequence of keys) that is contained twice in the YAML file. Since a
    // map (map[string]interface{}) is handled, the traversal sequence
    // cannot be predicted during an iteration. That is why we use the DPO
    // (Different Possible Outcomes) when checking.
	// yh.Footprints = []string{"herbivores"}

	if err := yh.FootprintSniffer(); err != nil {
		panic(err.Error())
	}

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


>__NOTE__
>
> To get the expected results, you should use V.3 of the
> [gopkg.in/yaml](https://gopkg.in/yaml.v3) package to unmarshalling the XML
> file. Do not use version V.2 or earlier of this package.
>
> Check the [tests](https://github.com/slaff-bg/yaml-hound/blob/main/yaml-hound_test.go) for more details.



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
