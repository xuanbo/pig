package entity

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func init() {
	var err error
	// Create a new Node with a Node number of 1
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

// GenerateID Generate a snowflake ID.
func GenerateID() string {
	id := node.Generate()
	return id.String()
}
