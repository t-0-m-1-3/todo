package backend

import "github.com/gobuffalo/packr"

type Config struct {
	Address string
	Static  packr.Box
}
