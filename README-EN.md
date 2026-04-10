# id-validator

[简体中文](README.md) | [ENGLISH](README-EN.md)

> Based on [guanguans/id-validator](https://github.com/guanguans/id-validator), extended with a built-in Web server and UI.
>
> **Note**: This fork only covers the Web layer. For the Go library itself — features, maintenance, and updates — please refer to the original project. If you need it as a Go dependency, go there directly.

[![tests](https://github.com/ropean/id-validator/actions/workflows/tests.yml/badge.svg)](https://github.com/ropean/id-validator/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/guanguans/id-validator)](https://goreportcard.com/report/github.com/guanguans/id-validator)
[![GoDoc](https://godoc.org/github.com/guanguans/id-validator?status.svg)](https://godoc.org/github.com/guanguans/id-validator)
[![GitHub license](https://img.shields.io/github/license/ropean/id-validator.svg)](https://github.com/ropean/id-validator/blob/master/LICENSE)

## Features

* Verify China ID number
* Get ID number information
* Upgrade 15-digit ID number to 18 digits
* Generate ID numbers that pass validation
* **Web server**: embedded UI, supports one-click Docker deployment

## Requirement

* Go >= 1.21

## Installation

```shell
go get -u github.com/guanguans/id-validator
```

## Web Server

```shell
# Local development
make dev

# Build & run
make build && make start

# Docker
make docker-build && make docker-up
```

Default port is `:8080`, configurable via the `PORT` environment variable.

## Usage

This is just a quick introduction, view the [GoDoc](https://godoc.org/github.com/guanguans/id-validator) for details.

```go
package main

import (
    idvalidator "github.com/guanguans/id-validator"
    "gopkg.in/ffmt.v1"
)

func main() {
    // Validate ID number
    ffmt.P(idvalidator.IsValid("500154199301135886", true))  // strict mode, 18-digit mainland
    ffmt.P(idvalidator.IsValid("500154199301135886", false)) // non-strict mode, 18-digit mainland
    ffmt.P(idvalidator.IsValid("11010119900307803X", false)) // 18-digit, ending with X
    ffmt.P(idvalidator.IsValid("610104620927690", false))    // 15-digit mainland
    ffmt.P(idvalidator.IsValid("810000199408230021", false)) // HK/Macao residence permit
    ffmt.P(idvalidator.IsValid("830000199201300022", false)) // Taiwan residence permit

    // Get ID info
    ffmt.P(idvalidator.GetInfo("500154199301135886", true))
    ffmt.P(idvalidator.GetInfo("500154199301135886", false))

    // Generate fake ID numbers
    ffmt.P(idvalidator.FakeId())
    ffmt.P(idvalidator.FakeRequireId(true, "江苏省", "200001", 1))

    // Upgrade 15-digit to 18-digit
    ffmt.P(idvalidator.UpgradeId("610104620927690"))
}
```

## Testing

```shell
make test
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Credits

This project is built on top of [guanguans/id-validator](https://github.com/guanguans/id-validator). The core validation logic and data originate from the original project. Many thanks to [guanguans](https://github.com/guanguans) and all contributors.

## Related projects

* [guanguans/id-validator](https://github.com/guanguans/id-validator), original project
* [jxlwqq/id-validator](https://github.com/jxlwqq/id-validator), by jxlwqq
* [jxlwqq/id-validator.py](https://github.com/jxlwqq/id-validator.py), by jxlwqq
* [mc-zone/IDValidator](https://github.com/mc-zone/IDValidator), by mc-zone
* [renyijiu/id_validator](https://github.com/renyijiu/id_validator), by renyijiu

## Reference material

* [People's Republic of China citizenship number](https://zh.wikipedia.org/wiki/中华人民共和国公民身份号码)
* [Ministry of Civil Affairs of the People's Republic of China: Administrative division code](http://www.mca.gov.cn/article/sj/xzqh/)
* [Historical data set of administrative division codes of the People's Republic of China](https://github.com/jxlwqq/address-code-of-china)
* [Residence Permit for Hong Kong, Macao and Taiwan Residents](https://zh.wikipedia.org/wiki/港澳台居民居住证)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
