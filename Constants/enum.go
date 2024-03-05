package Constants

import (
	"errors"
)

type CountryCode string

const (
	PL CountryCode = "pl"
	UA CountryCode = "ua"
	US CountryCode = "us"
	IN CountryCode = "in" //India
	AF CountryCode = "af" //Afghanistan
	AS CountryCode = "as" //American Samoa
	AU CountryCode = "au" //Australia
	BD CountryCode = "bd" //Bangladesh
	BE CountryCode = "be" //Belgium
	BT CountryCode = "bt" //Bhutan
	BR CountryCode = "br" //Brazil
	KH CountryCode = "kh" //Cambodia
	CA CountryCode = "ca" //Canada
	CO CountryCode = "co" //Colombia
	DE CountryCode = "de" //Germany
	GR CountryCode = "gr" //Greece
	GL CountryCode = "gl" //Greenland
)

func (c CountryCode) IsValidConutryCode() bool {
	return c == PL ||
		c == UA ||
		c == US ||
		c == AF ||
		c == AS ||
		c == IN ||
		c == AU ||
		c == BD ||
		c == BE ||
		c == BT ||
		c == BR ||
		c == KH ||
		c == CA ||
		c == CO ||
		c == DE ||
		c == GR ||
		c == GL
}

var (
	ErrInvalidCountryCodeFound = errors.New("ErrInvalidCountryCodeFound")
)
