package storm

import (
	"encoding/json"
	"testing"
)

func TestServerZoneMarshal(t *testing.T) {
	sz := ServerZone(123)

	marshalled, err := json.Marshal(sz)
	if err != nil {
		t.Error(err)
	}
	if string(marshalled) != "123" {
		t.Errorf("Marshal was incorrect, got %v", marshalled)
	}
}

func TestServerZoneUnmarshal(t *testing.T) {
	szJSON := []byte(`{ "id": 123 }`)
	var sz ServerZone

	err := json.Unmarshal(szJSON, &sz)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", sz)
	if sz != 123 {
		t.Errorf("Unmarshal was incorrect, got %v", sz)
	}
}
