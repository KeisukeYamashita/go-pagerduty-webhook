# go-pagerduty-webhook

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](license)
[![CircleCI](https://circleci.com/gh/KeisukeYamashita/go-pagerduty-webhook.svg?style=svg)](https://circleci.com/gh/KeisukeYamashita/go-pagerduty-webhook)
[![codecov](https://codecov.io/gh/KeisukeYamashita/go-pagerduty-webhook/branch/master/graph/badge.svg)](https://codecov.io/gh/KeisukeYamashita/go-pagerduty-webhook)

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