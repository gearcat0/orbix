package signature

import (
	"encoding/hex"
	"github.com/gearcat0/orbix/xsig"
)

func Sign(secret string, payload map[string]interface{}) string {
	sig := xsig.Sign([]byte(secret), payload)
	dst := make([]byte, hex.EncodedLen(len(sig)))
	hex.Encode(dst, sig)
	return string(dst)
}
