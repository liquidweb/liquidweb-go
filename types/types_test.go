package types

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	ts, err := NewTimestamp("2009-09-09 11:12:13")
	if err != nil {
		t.Error(err)
	}
	marshalled, err := json.Marshal(ts)
	if err != nil {
		t.Error(err)
	}
	if string(marshalled) != `"2009-09-09 11:12:13"` {
		t.Errorf("Marshal was incorrect, got %v", marshalled)
	}
}

func TestUnMarshal(t *testing.T) {
	tsJSON := []byte(`"2009-09-09 11:12:13"`)
	var ts Timestamp

	err := json.Unmarshal(tsJSON, &ts)
	if err != nil {
		t.Error(err)
	}

	if ts.String() != `2009-09-09 11:12:13` {
		t.Errorf("Unmarshal was incorrect, got %v", ts)
	}
}
