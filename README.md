# libra-sdk-go
![Go Build Status](https://github.com/libra-sigs/libra-sdk-go/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/libra-sigs/libra-sdk-go)](https://goreportcard.com/report/github.com/libra-sigs/libra-sdk-go)

libra sdk for go


## Examples

### libra-client

```
% make; _output/examples/libra-client 
find . -name '*.go' | xargs gofmt -w -s
go build -o  _output/examples/libra-client  ./examples/libra-client
go test ./libra
ok  	github.com/libra-sigs/libra-sdk-go/libra	(cached) [no tests to run]
Account resource: {
    "authkey": "Xw5ayK30K7AqE8U7adbpDw5ewCDdvlwuestwqdSKhc4=",
    "balance": 209000000,
    "key_rotation_cap": false,
    "with_drawal_cap": false,
    "recv_event_count": 2,
    "recv_event_data": "AAAAAAAAAABfDlrIrfQrsCoTxTtp1ukPDl7AIN2+XC56y3Cp1IqFzg==",
    "sent_event_count": 2,
    "sent_event_data": "AQAAAAAAAABfDlrIrfQrsCoTxTtp1ukPDl7AIN2+XC56y3Cp1IqFzg==",
    "seq_no": 2,
    "event_no": 2
}
```
