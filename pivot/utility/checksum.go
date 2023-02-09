package utility

import (
	"crypto/sha256"

	"github.com/pivot-g/pivot/pivot/log"
)

func GenCheckSum(files *map[string][]byte) *map[string][32]byte {
	out := map[string][32]byte{}
	for file, data := range *files {
		sha256 := sha256.Sum256(data)
		log.Debug("Check sum generated")
		out[file] = sha256
	}
	return &out
}
