package idx

import "github.com/rs/xid"

func New() string {
	return xid.New().String()
}
