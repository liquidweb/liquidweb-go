package types

import "strconv"

type NumericalBoolean bool

func (nb NumericalBoolean) UnmarshalJSON(d []byte) error {
	dString := string(d)
	s, err := strconv.ParseBool(dString)
	if err != nil {
		return err
	}

	nb = NumericalBoolean(s)

	return nil
}
