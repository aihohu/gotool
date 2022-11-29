package idUtil

import "github.com/bwmarrin/snowflake"

// SnowflakeId 雪花算法
// Create a new Node with a Node number of 1
var node, _ = snowflake.NewNode(1)

// SnowflakeId 简单获取SnowflakeId
// Node number 默认为1
// @return SnowflakeId
func SnowflakeId() string {
	return node.Generate().String()
}

// SnowflakeIdInt64 简单获取SnowflakeId
// Node number 默认为1
// @return SnowflakeId(int64)
func SnowflakeIdInt64() int64 {
	return node.Generate().Int64()
}
