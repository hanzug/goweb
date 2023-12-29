package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var sf *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	sf, err = snowflake.NewNode(machineID)

	return
}

func GenID() int64 {
	return sf.Generate().Int64()
}
