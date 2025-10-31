# Notification Center

This is a notification microservice.

## Features
- support apns and mipush notifications
- REST API to manage notification and user device tokens

## Usage

### Build
```shell
git clone https://github.com/OpenTreeHole/notification.git
cd notification
# install swag and generate docs
go install github.com/swaggo/swag/cmd/swag@latest
swag init --parseInternal --parseDependency --parseDepth 1 # to generate the latest docs, this should be run before compiling
# build and run
go build -o notification.exe
./notification.exe
```

### Run
Before running, export `MODE` and `BASE_PATH`:

```bash
export BASE_PATH=$PWD
# Available modes: production / dev / test / perf
export MODE=dev
go run main.go
```

### Test
Please export `MODE=test` and `BASE_PATH=$PWD`
to avoid relative path errors in unit tests.

Device tokens must be set to test push notifications,
export `${service}_DEVICE_TOKEN` for each push service,
e.g. `APNS_DEVICE_TOKEN=1234567`

### API Docs
Please visit http://localhost:8000/docs after running app

## Badge

[//]: # ([![build]&#40;https://github.com/OpenTreeHole/notification/actions/workflows/master.yaml/badge.svg&#41;]&#40;https://github.com/OpenTreeHole/notification/actions/workflows/master.yaml&#41;)
[//]: # ([![dev build]&#40;https://github.com/OpenTreeHole/notification/actions/workflows/dev.yaml/badge.svg&#41;]&#40;https://github.com/OpenTreeHole/notification/actions/workflows/dev.yaml&#41;)

[![stars](https://img.shields.io/github/stars/OpenTreeHole/notification)](https://github.com/OpenTreeHole/notification/stargazers)
[![issues](https://img.shields.io/github/issues/OpenTreeHole/notification)](https://github.com/OpenTreeHole/notification/issues)
[![pull requests](https://img.shields.io/github/issues-pr/OpenTreeHole/notification)](https://github.com/OpenTreeHole/notification/pulls)

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

### Powered by

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/OpenTreeHole/notification/issues/new) or [Submit PRs](https://github.com/OpenTreeHole/notification/compare).

We are now in rapid development, any contribution would be of great help.
For the developing roadmap, please visit [this issue](https://github.com/OpenTreeHole/notification/issues/1).

### Contributors

This project exists thanks to all the people who contribute.

<a href="https://github.com/OpenTreeHole/notification/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=OpenTreeHole/notification"  alt="contributors"/>
</a>

## Licence

[![license](https://img.shields.io/github/license/OpenTreeHole/notification)](https://github.com/OpenTreeHole/notification/blob/master/LICENSE)
© OpenTreeHole
