# YAML Hound #
Functionalities covering the need for non-standard approaches when working with
YAML files.


[![Code coverage][shield-coverage]](#)
[![MIT licensed][shield-license]](#)
![badge](https://gist.githubusercontent.com/slaff-bg/0fdaa350d0428e1801a1a3f9d9e9ca98/raw/3fac59b83e1eef9cf7936c441451584b0e464c8b/test.json)


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

- [ ] Retrieves a value by key or series of known keys from an unknown dynamic YAML structure.
- [ ] N-level keys ripper.
- [ ] etc.


## Requirements ##

- [Golang](https://go.dev/dl/) version go1.20

&#x1F4CC; &nbsp; *<sub>Versions reflect the current state of the used
technologies.</sub>*

## Usage ##



## Contributing ##

To contribute to Paddington, clone this repo locally and commit your code on a separate branch. Please write unit tests for your code, and run the linter before opening a pull-request:


## License ##

[License rights and limitations (MIT).](https://github.com/slaff-bg/yaml-hound/blob/main/LICENSE)


[shield-coverage]: https://img.shields.io/badge/coverage-0%25-brightgreen.svg
[shield-license]: https://img.shields.io/badge/license-MIT-blue.svg
