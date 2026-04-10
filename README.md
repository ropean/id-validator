# id-validator

[简体中文](README.md) | [ENGLISH](README-EN.md)

> 基于 [guanguans/id-validator](https://github.com/guanguans/id-validator) 增加了 Web 服务和可视化界面。
>
> **注意**：本项目仅负责 Web 端，Go 库本身的功能、维护和更新请以原项目为准。如需将其作为 Go 依赖库使用，请直接前往原作者仓库。

[![tests](https://github.com/guanguans/id-validator/actions/workflows/tests.yml/badge.svg)](https://github.com/guanguans/id-validator/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/guanguans/id-validator)](https://goreportcard.com/report/github.com/guanguans/id-validator)
[![GoDoc](https://godoc.org/github.com/guanguans/id-validator?status.svg)](https://godoc.org/github.com/guanguans/id-validator)
[![GitHub license](https://img.shields.io/github/license/ropean/id-validator.svg)](https://github.com/ropean/id-validator/blob/master/LICENSE)

## 功能

* 验证中国身份证号
* 获取身份证号信息
* 升级 15 位身份证号为 18 位
* 伪造符合校验的身份证号
* **Web 服务**：内嵌前端界面，支持 Docker 一键部署

## 环境要求

* Go >= 1.21

## 安装

```shell
go get -u github.com/guanguans/id-validator
```

## Web 服务

```shell
# 本地运行
make dev

# 构建
make build && make start

# Docker
make docker-build && make docker-up
```

默认监听 `:5100`，通过环境变量 `PORT` 修改端口。

## 使用

这只是一个快速介绍，请查看 [GoDoc](https://godoc.org/github.com/guanguans/id-validator) 获得详细信息。

```go
package main

import (
    idvalidator "github.com/guanguans/id-validator"
    "gopkg.in/ffmt.v1"
)

func main() {
    // 验证身份证号合法性
    ffmt.P(idvalidator.IsValid("500154199301135886", true))  // 严格模式验证大陆居民身份证18位
    ffmt.P(idvalidator.IsValid("500154199301135886", false)) // 非严格模式验证大陆居民身份证18位
    ffmt.P(idvalidator.IsValid("11010119900307803X", false)) // 大陆居民身份证末位是X18位
    ffmt.P(idvalidator.IsValid("610104620927690", false))    // 大陆居民身份证15位
    ffmt.P(idvalidator.IsValid("810000199408230021", false)) // 港澳居民居住证18位
    ffmt.P(idvalidator.IsValid("830000199201300022", false)) // 台湾居民居住证18位

    // 获取身份证号信息
    ffmt.P(idvalidator.GetInfo("500154199301135886", true))  // 严格模式获取身份证号信息
    ffmt.P(idvalidator.GetInfo("500154199301135886", false)) // 非严格模式获取身份证号信息

    // 生成可通过校验的假身份证号
    ffmt.P(idvalidator.FakeId())                                  // 随机生成
    ffmt.P(idvalidator.FakeRequireId(true, "江苏省", "200001", 1)) // 生成出生于2000年1月江苏省的男性居民身份证

    // 15位号码升级为18位
    ffmt.P(idvalidator.UpgradeId("610104620927690"))
}
```

## 测试

```shell
make test
```

## 变更日志

请参阅 [CHANGELOG](CHANGELOG.md) 获取最近有关更改的更多信息。

## 致谢

本项目基于 [guanguans/id-validator](https://github.com/guanguans/id-validator) 开发，核心验证逻辑及数据来自原项目，在此对原作者 [guanguans](https://github.com/guanguans) 及所有贡献者表示感谢。

## 相关项目

* [guanguans/id-validator](https://github.com/guanguans/id-validator)，原始项目
* [jxlwqq/id-validator](https://github.com/jxlwqq/id-validator)，jxlwqq
* [jxlwqq/id-validator.py](https://github.com/jxlwqq/id-validator.py)，jxlwqq
* [mc-zone/IDValidator](https://github.com/mc-zone/IDValidator)，mc-zone
* [renyijiu/id_validator](https://github.com/renyijiu/id_validator)，renyijiu

## 参考资料

* [中华人民共和国公民身份号码](https://zh.wikipedia.org/wiki/中华人民共和国公民身份号码)
* [中华人民共和国民政部：行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/)
* [中华人民共和国行政区划代码历史数据集](https://github.com/jxlwqq/address-code-of-china)
* [国务院办公厅关于印发《港澳台居民居住证申领发放办法》的通知](http://www.gov.cn/zhengce/content/2018-08/19/content_5314865.htm)
* [港澳台居民居住证](https://zh.wikipedia.org/wiki/港澳台居民居住证)

## 协议

MIT 许可证（MIT）。有关更多信息，请参见[协议文件](LICENSE)。
