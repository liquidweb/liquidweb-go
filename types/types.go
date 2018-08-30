package types

import (
	"encoding/json"
	"strconv"
)

// NumericalBoolean handles booleans in the API being returned as
type NumericalBoolean bool

// UnmarshalJSON handles the unmarshalling of the NumericalBoolean type.
func (nb *NumericalBoolean) UnmarshalJSON(d []byte) error {
	var n int64
	if err := json.Unmarshal(d, &n); err != nil {
		return err
	}
	nString := strconv.FormatInt(n, 10)
	s, err := strconv.ParseBool(nString)
	if err != nil {
		return err
	}

	*nb = NumericalBoolean(s)

	return nil
}

// FlexInt was shamelessly stolen from Chris to handle inconsistencies in the API returning numerical versus string IDs.
type FlexInt int

func (fi *FlexInt) String() string {
	return strconv.Itoa(int(*f))
}

// UnmarshalJSON handles the unmarshalling of the FlexInt type.
func (fi *FlexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = FlexInt(i)
	return nil
}
