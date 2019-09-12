package main

// Graph with edges and nodes
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// Edge with pointer from and to
type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Node with identitifier and type
type Node struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
