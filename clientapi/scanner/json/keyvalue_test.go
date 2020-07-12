package json_test

import (
	"encoding/json"
	"strings"
	"testing"

	. "github.com/walkline/shippingpg/clientapi/scanner/json"
)

func TestScanForKeyValueScannerWithPortsData(t *testing.T) {
	r := strings.NewReader(`{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }
}`)
	scanner := NewKeyValueScanner(r, 20)
	result := map[string]json.RawMessage{}
	for scanner.Scan() {
		result[scanner.Key()] = scanner.Value()
	}

	if string(result["AEAJM"]) != `{
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }` {
		t.Error("bad value for key 'AEAJM'", string(result["AEAJM"]))
	}

	if string(result["AEAUH"]) != `{
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }` {
		t.Error("bad value for key 'AEAUH'", string(result["AEAUH"]))
	}
}
