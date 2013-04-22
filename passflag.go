/*
passflag is a micropackage that provides a Password type flag.
If the Password flag is used, a silence, no echo'ed prompt will ask for the
password.
*/
package passflag

import (
	"flag"
	"github.com/gcmurphy/getpass"
)

// Password sets a password flag.
// It uses the same API format as the `flag` built-in package.
func Password(f string, def bool, des string) *bool {
	b := flag.Bool(f, def, des)
	flag.Parse()
	return b
}

// Passflag will prompt a password input and returns the password
// introduced and an error if exists.
func Get() (string, error) {
	return getpass.GetPass()
}
