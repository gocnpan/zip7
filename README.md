<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-zip7.svg"/></a></p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/essentialkaos/zip7"><img src="https://pkg.go.dev/badge/github.com/essentialkaos/zip7" /></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/zip7"><img src="https://goreportcard.com/badge/github.com/essentialkaos/zip7"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-zip7-master"><img alt="codebeat badge" src="https://codebeat.co/badges/11fb655d-8da8-4694-a32b-b95ff9eed602" /></a>
  <a href="https://github.com/essentialkaos/zip7/actions"><img src="https://github.com/essentialkaos/zip7/workflows/CI/badge.svg" alt="GitHub Actions Status" /></a>
  <a href="https://github.com/essentialkaos/zip7/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/zip7/workflows/CodeQL/badge.svg" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#compatibility-and-os-support">Compatibility and OS support</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`zip7` package provides methods for working with 7z archives (`p7zip` wrapper).

### Installation

Make sure you have a working Go 1.14+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/zip7.v1
```

If you want to update `zip7` to latest stable release, do:

```
go get -u pkg.re/essentialkaos/zip7.v1
```

### Compatibility and OS support

|      Version |     1.x |
|--------------|---------|
|  `p7zip 9.x` | Partial |
| `p7zip 15.x` |    Full |
| `p7zip 16.x` |    Full |

| OS       | Support            |
|----------|--------------------|
| Linux    | :heavy_check_mark: |
| Mac OS X | :heavy_check_mark: |
| FreeBSD  | :heavy_check_mark: |
| Windows  | :x:                |

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://github.com/essentialkaos/zip7/workflows/CI/badge.svg?branch=master)](https://github.com/essentialkaos/zip7/actions) |
| `develop` | [![CI](https://github.com/essentialkaos/zip7/workflows/CI/badge.svg?branch=develop)](https://github.com/essentialkaos/zip7/actions) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
