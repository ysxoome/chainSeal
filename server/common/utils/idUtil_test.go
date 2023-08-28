package utils

import (
	"fmt"
	"testing"
)

func TestRandomUUID(t *testing.T) {
	fmt.Println(RandomUUID())
}

func TestSimpleUUID(t *testing.T) {
	fmt.Println(SimpleUUID())
}

func TestAddrToGid(t *testing.T) {
	fmt.Println(AddrToGID("0x0af384061c5e257daf5fb0ea45e2840c7e106395"))
	// did:gid:0af384061c5e257daf5fb0ea45e2840c7e106395
}

func TestGidToAddr(t *testing.T) {
	addr := (GidToAddr("did:gid:0af384061c5e257daf5fb0ea45e2840c7e106395"))
	fmt.Println(isVaildAddress(addr))
}