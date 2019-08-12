package types

import (
	"encoding/json"
	"net"
	"strconv"
	"time"
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
	return strconv.Itoa(int(*fi))
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

// IPAddr uses net.IP but supports Marshal/Unmarshal.
type IPAddr struct {
	net.IP
}

// String returns a string representation of the IP.
func (ip *IPAddr) String() string {
	return ip.IP.String()
}

// UnmarshalJSON unmarshals the IPAddr type.
func (ip *IPAddr) UnmarshalJSON(b []byte) error {
	var astr string
	if err := json.Unmarshal(b, &astr); err != nil {
		return err
	}

	ip.IP = net.ParseIP(astr)

	return nil
}

// MarshalJSON marshalls the IPAddr type.
func (ip *IPAddr) MarshalJSON() ([]byte, error) {
	return []byte(ip.IP.String()), nil
}

// Timestamp implements Liquid Web's custom timestamp format.
type Timestamp struct {
	time.Time
}

// LWTimestampFormat is Liquid Web's timestamp format.
const LWTimestampFormat = "2006-01-02 15:04:05"

// UnmarshalJSON parses Liquid Web's timestamp format.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := time.Parse(LWTimestampFormat, string(b))
	if err != nil {
		return err
	}
	*t = Timestamp(ts)
	return nil
}

// MarshalJSON marshalls the IPAddr type.
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts, err := time.Parse(LWTimestampFormat, t.String())
	if err != nil {
		return nil, err
	}
	return []byte(ts.String()), nil
}
