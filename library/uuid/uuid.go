package uuid

import "github.com/bwmarrin/snowflake"

const nodeBits = 3
const stepBits = 3 // maximum 8 unique IDs to be generated every millisecond
const initNode = 1

func init() {
	snowflake.NodeBits = nodeBits
	snowflake.StepBits = stepBits
}

func UUID() (int64, error) {
	// Create a new Node with a Node number of initNode
	node, err := snowflake.NewNode(initNode)
	if err != nil {
		return 0, err
	}
	return node.Generate().Int64(), nil
}
