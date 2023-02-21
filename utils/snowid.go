package utils

import (
	"github.com/bwmarrin/snowflake"
)

const nodeID int64 = 1

// GenID 生成ID时会上锁，确保不重复
func GenID() (int64, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return 0, err
	}
	return node.Generate().Int64(), nil
}
