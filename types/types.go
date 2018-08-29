package types

import (
	"encoding/json"
	"strconv"
)

// NumericalBoolean handles booleans in the API being returned as
type NumericalBoolean bool

// UnmarshalJSON handles the unmarshalling of the NumericalBoolean type.
func (nb NumericalBoolean) UnmarshalJSON(d []byte) error {
	dString := string(d)
	s, err := strconv.ParseBool(dString)
	if err != nil {
		return err
	}

	nb = NumericalBoolean(s)

	return nil
}

// FlexInt was shamelessly stolen from Chris to handle inconsistencies in the API returning numerical versus string IDs.
type FlexInt int

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
