package internal

import (
	"encoding/json"
	"testing"
)

func TestRecordUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		json string
	}{
		{
			name: "geo fields as empty strings",
			json: `{"Id":1,"Type":3,"Ttl":300,"Value":"test","Name":"_acme","IPGeoLocationInfo":"","GeolocationInfo":""}`,
		},
		{
			name: "geo fields as objects",
			json: `{"Id":1,"Type":3,"Ttl":300,"Value":"test","Name":"_acme","IPGeoLocationInfo":{"CountryCode":"US","City":"NYC"},"GeolocationInfo":{"Lat":40.7,"Lon":-74.0}}`,
		},
		{
			name: "geo fields absent",
			json: `{"Id":1,"Type":3,"Ttl":300,"Value":"test","Name":"_acme"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var r Record
			if err := json.Unmarshal([]byte(tc.json), &r); err != nil {
				t.Fatalf("failed to unmarshal Record: %v", err)
			}
			if r.Value != "test" {
				t.Errorf("expected Value=test, got %s", r.Value)
			}
		})
	}
}

func TestZoneResponseUnmarshal(t *testing.T) {
	raw := `{
		"Items": [{
			"Id": 123,
			"Domain": "example.com",
			"Records": [
				{"Id":1,"Type":3,"Ttl":300,"Value":"v","Name":"n","IPGeoLocationInfo":{"CC":"US"},"GeolocationInfo":""}
			]
		}],
		"CurrentPage": 1,
		"TotalItems": 1,
		"HasMoreItems": false
	}`

	var zr ZoneResponse
	if err := json.Unmarshal([]byte(raw), &zr); err != nil {
		t.Fatalf("failed to unmarshal ZoneResponse: %v", err)
	}
	if len(zr.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(zr.Items))
	}
	if len(zr.Items[0].Records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(zr.Items[0].Records))
	}
}
