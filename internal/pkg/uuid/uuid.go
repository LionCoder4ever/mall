package uuid

import (
	"github.com/bwmarrin/snowflake"
)

const nodeBits = 3
const stepBits = 3 // maximum 8 unique IDs to be generated every millisecond
const initNode = 1

var node *snowflake.Node

func NewUUID() (err error) {
	snowflake.NodeBits = nodeBits
	snowflake.StepBits = stepBits
	// Create a new Node with a Node number of initNode
	node, err = snowflake.NewNode(initNode)
	return err
}

func UUID() int64 {
	return node.Generate().Int64()
}
