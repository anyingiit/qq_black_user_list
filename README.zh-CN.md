[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:a01fab029ed44647 -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# qq_black_user_list

一个基于 Go 与 Gin 的 HTTP 接口，在 MySQL 数据表中查询某个 QQ 号，返回其记录的事件内容与核实状态，若该号码不在表中则报告未找到。

[![CI](https://github.com/anyingiit/qq_black_user_list/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/qq_black_user_list/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/qq_black_user_list)](LICENSE)

[报告问题](https://github.com/anyingiit/qq_black_user_list/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/qq_black_user_list/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

qq_black_user_list 是一个只有一个可用接口的小型 Gin HTTP 服务：
`GET /api/v1/public/even/:qq_num`（router/router.go），由
router/api/v1/blackListEven.go 中的 `GetEvenForQq` 处理。给定一个 QQ 号，它会在
MySQL 数据表中查找匹配的记录，返回记录的事件内容与核实状态；如果该号码没有记录，
则返回"未找到"错误码（models/blackListEvenModels.go）。

处理函数查询的表名是 `qqblk_even`，但 docs/even.sql 中的建表脚本创建的表名是
`even`。在这份脚本能真正为代码执行的查询提供数据之前，两者需要先对齐——例如按
`pkg/setting/setting.go` 从配置读取的表前缀来命名。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/qq_black_user_list/issues)。

## 开始使用

### 环境要求

- Go 1.16 或更高版本，即 `go.mod` 声明的下限
- 一个应用可以连接的 MySQL 服务，通过 `pkg/setting/setting.go` 读取的配置项
  （`TYPE`、`USER`、`PASSWORD`、`HOST`、`NAME`）配置
- 仓库已经附带的 Casbin RBAC 模型与策略文件，位于 `conf/casbin/`，由
  `pkg/casbin/initCasbin.go` 在启动时加载

### 安装

```sh
git clone https://github.com/anyingiit/qq_black_user_list.git
cd qq_black_user_list
go build ./...
```

创建 `conf/app/app.ini`（由 `pkg/setting/setting.go` 读取），至少包含
`[server]`、`[app]`、`[database]` 三个分节，并在 `[database]` 中填写数据库类型、
用户名、密码、主机和库名。然后创建处理函数实际查询的表——应用
`docs/even.sql` 中的建表脚本，并将表命名（或改名）为 `qqblk_even`，
以匹配 `models/blackListEvenModels.go` 真正查询的表名。

## 用法

```sh
go run main.go
curl http://127.0.0.1:8080/api/v1/public/even/<qq_num>
```

在表中查到的 QQ 号会返回其 `even_content`、`even_content_more_info` 和
`verify_status` 字段；没有对应记录的号码会返回 `ErrNoQqInList` 错误码
（`pkg/e/code_def.go`）。

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/qq_black_user_list](https://github.com/anyingiit/qq_black_user_list)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
