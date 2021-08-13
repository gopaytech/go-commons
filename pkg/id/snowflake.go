package id

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sony/sonyflake"
)

var snowflake *sonyflake.Sonyflake
var snowflakeSetting sonyflake.Settings

func init() {
	rand.Seed(time.Now().UnixNano())
	snowflakeSetting.MachineID = func() (machineId uint16, e error) {
		return uint16(rand.Uint64()), nil
	}
}

func NextID() (uint64, error) {
	snowflake = sonyflake.NewSonyflake(snowflakeSetting)
	return snowflake.NextID()
}

func NextHexID() (id string, err error) {
	snowflake = sonyflake.NewSonyflake(snowflakeSetting)
	sId, err := snowflake.NextID()
	id = fmt.Sprintf("%X", sId)
	return
}

func Extract(id uint64) map[string]uint64 {
	return sonyflake.Decompose(id)
}
