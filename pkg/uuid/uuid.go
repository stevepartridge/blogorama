package uuid

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/stevepartridge/blogorama/pkg/utils"
)

func Generate() string {
	u, err := uuid.NewUUID()
	if err != nil {
		return generateFallback()
	}
	return u.String()
}

func generateFallback() string {
	rand := utils.RandStringWithLengthLimit(17)
	now := time.Now().UTC().UnixNano()
	raw := utils.Sha256(fmt.Sprintf("%s%d", rand, now))

	if len(raw) < 32 {
		raw = utils.RandStringWithLengthLimit(32)
	}

	return fmt.Sprintf("%s-%s-%s-%s-%s",
		raw[:8],
		raw[8:12],
		raw[12:16],
		raw[16:20],
		raw[20:32],
	)
}
