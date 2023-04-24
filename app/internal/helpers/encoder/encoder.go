package encoder

import (
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/types"
	"gopkg.in/hedzr/errors.v3"
)

// Encode encodes the password with given
// type of algo (bcrypt)
func Encode(value string) (string, error) {
	encoding := encoder.New(types.Bcrypt)

	hash, err := encoding.Encode(value)
	if err != nil {
		return "", errors.New("encoding value error:%v", err)
	}
	return hash, nil
}

// IsMatch method returns true if the hash and original
// password are match, else returns false
// The method can be used to authorize the user in the system
func IsMatch(encoded, original string) (bool, error) {
	encoding := encoder.New(types.Bcrypt)
	verify, err := encoding.Verify(encoded, original)
	if err != nil {
		return false, errors.New("error of matching password:%v", err)
	}
	return verify, nil
}
