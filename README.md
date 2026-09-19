<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# qq_black_user_list

A Go and Gin HTTP API that looks up a QQ account number in a MySQL table and returns its recorded incident and verification status, or reports the number as not listed.

**English** · [简体中文](README.zh-CN.md)

[![CI](https://github.com/anyingiit/qq_black_user_list/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/qq_black_user_list/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/qq_black_user_list)](LICENSE)

[Report a bug](https://github.com/anyingiit/qq_black_user_list/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/qq_black_user_list/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

qq_black_user_list is a small Gin-based HTTP API with one working endpoint:
`GET /api/v1/public/even/:qq_num` (router/router.go), handled by `GetEvenForQq`
in router/api/v1/blackListEven.go. Given a QQ account number, it looks up a
matching row in a MySQL table and returns the recorded incident text and a
verification status, or a "not found" error code if the number has no record
(models/blackListEvenModels.go).

The handler queries a table named `qqblk_even`, but the schema shipped in
docs/even.sql creates a table named `even`. The two need to be reconciled
before that schema will actually serve the query the code runs — for example
by applying the table prefix `pkg/setting/setting.go` reads from
configuration.

See the [open issues](https://github.com/anyingiit/qq_black_user_list/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- Go 1.16 or newer, the floor declared in `go.mod`
- A MySQL server the app can reach, configured through the keys
  `pkg/setting/setting.go` loads (`TYPE`, `USER`, `PASSWORD`, `HOST`, `NAME`)
- The Casbin RBAC model and policy already shipped under `conf/casbin/`, which
  `pkg/casbin/initCasbin.go` loads at startup

### Installation

```sh
git clone https://github.com/anyingiit/qq_black_user_list.git
cd qq_black_user_list
go build ./...
```

Create `conf/app/app.ini` (read by `pkg/setting/setting.go`) with `[server]`,
`[app]` and `[database]` sections carrying at least the database type, user,
password, host and name. Then create the table the handler queries — apply
the schema in `docs/even.sql`, naming (or renaming) the table `qqblk_even` to
match what `models/blackListEvenModels.go` actually selects from.

## Usage

```sh
go run main.go
curl http://127.0.0.1:8080/api/v1/public/even/<qq_num>
```

A QQ number found in the table comes back with its `even_content`,
`even_content_more_info` and `verify_status` fields; a number with no row
comes back with the `ErrNoQqInList` error code (`pkg/e/code_def.go`).

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/qq_black_user_list](https://github.com/anyingiit/qq_black_user_list)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
