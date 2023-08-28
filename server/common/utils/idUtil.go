package utils

import (
	"math/big"
	"seal/server/common/xerr"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	ADDRESS_PREFIX = "0x"
	GID_PREFIX     = "did:gid:"
)

func RandomUUID() string {
	return uuid.NewV4().String()
}

func SimpleUUID() string {
	return strings.Replace(RandomUUID(), "-", "", -1)
}

func AddrToGID(addr string) (did string, err error) {
	if addr == "" {
		return "", xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)
	}
	addr = strings.TrimPrefix(addr, ADDRESS_PREFIX)
	did += GID_PREFIX
	did += addr

	return
}

func GidToAddr(gid string) (addr string) {
	if !strings.HasPrefix(gid, GID_PREFIX) {
		return ""
	}
	addr += ADDRESS_PREFIX
	addr += strings.TrimPrefix(gid, GID_PREFIX)

	return
}

func isGidVaild(gid string) bool {
	return gid != "" &&
		strings.HasPrefix(gid, GID_PREFIX) &&
		isVaildAddress(GidToAddr(gid))
}

func isVaildAddress(addr string) bool {
	data := strings.TrimPrefix(addr, ADDRESS_PREFIX)
	_, err := new(big.Int).SetString(data, 16)
	if !err {
		return false
	}

	return len(data) == 40
}

/* ========================================================== */
