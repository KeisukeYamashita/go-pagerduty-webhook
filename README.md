# go-pagerduty-webhook

[![License][license-badge]][license]
[![CircleCI][circleci-badge]][circleci]
[![codecov](https://codecov.io/gh/KeisukeYamashita/go-pagerduty-webhook/branch/master/graph/badge.svg)](https://codecov.io/gh/KeisukeYamashita/go-pagerduty-webhook)
[![GolangCI][golangci-badge]][golangci]

`go-pagerduty-webhook` is a package to use for [Pagerduty Webhook V2](https://v2.developer.pagerduty.com/docs/webhooks-v2-overview).

## What

The goal is to provide all types for parcing payload from the Webhook.

## Usage

### Use in your Webhook server

```
func (ctr Controller) IncidentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var payload webhook.Payload
		err = decoder.Decode(&payload)
    }
}
```

## Author

* [KeisukeYamashita<19yamashita15@gmail.com>](https://github.com/KeisukeYamashita)

<!-- badge links -->

[license]: LICENSE
[circleci]: https://circleci.com/gh/KeisukeYamashita/workflows/go-pagerduty-recorder
[godoc]: https://godoc.org/github.com/KeisukeYamashita/go-pagerduty-recorder
[go-report-card]: https://goreportcard.com/report/github.com/KeisukeYamashita/go-pagerduty-recorder
[golangci]: https://golangci.com/r/github.com/KeisukeYamashita/go-pagerduty-recorder

[license-badge]: https://img.shields.io/badge/license-Apache%202.0-%23E93424
[circleci-badge]: https://img.shields.io/circleci/project/github/KeisukeYamashita/go-pagerduty-webhook?label=circleci&logo=circleci
[godoc-badge]: https://img.shields.io/badge/godoc.org-reference-blue.svg
[go-report-card-badge]: https://goreportcard.com/badge/github.com/KeisukeYamashita/go-pagerduty-recorder
[golangci-badge]: https://golangci.com/badges/github.com/KeisukeYamashita/go-pagerduty-webhook.svg