package json

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net"
	"net/url"
)

// Standard library types with json.Marshaler/Unmarshaler or
// encoding.TextMarshaler/TextUnmarshaler implementations

type IPNet struct {
	net.IPNet
}

func (i *IPNet) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *IPNet) UnmarshalText(text []byte) error {
	_, parsed, err := net.ParseCIDR(string(text))
	if err != nil {
		return err
	}
	i.IP = parsed.IP
	i.Mask = parsed.Mask
	return nil
}

type URL struct {
	url.URL
}

func (u *URL) MarshalJSON() ([]byte, error) {
	return u.MarshalBinary()
}

func (u *URL) UnmarshalJSON(text []byte) error {
	return u.UnmarshalBinary(text)
}
