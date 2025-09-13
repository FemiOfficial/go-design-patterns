package utils

import (
	"bytes"
	"io"
)

func GetUnProcessedMockOrder() io.Reader {
	order := `{
		"id": "1",
		"customerId": "1",
		"total": 40.04,
		"processed": false,
		"items": [
			{
				"id": 1,
				"name": "bag",
				"price": 20.02
			},
			{
				"id": 1091,
				"name": "shoe",
				"price": 20.02
			}
		]
	}`

	return bytes.NewReader([]byte(order))
}

func GetProcessedMockOrder() io.Reader {
	order := `{
		"id": "1",
		"customerId": "1",
		"total": 40.04,
		"processed": true,
		"items": [
			{
				"id": 1,
				"name": "bag",
				"price": 20.02
			},
			{
				"id": 1091,
				"name": "shoe",
				"price": 20.02
			}
		]
	}`

	return bytes.NewReader([]byte(order))
}

func GetInvalidMockOrder() io.Reader {
	order := `{
		"customerId": "1",
		"total": 40.04,
		"processed": false,
		"items": [
			{
				"id": 1,
				"name": "bag",
				"price": 20.02
			},
			{
				"id": 1091,
				"name": "shoe",
				"price": 20.02
			}
		]
	}`

	return bytes.NewReader([]byte(order))
}