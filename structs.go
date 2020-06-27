package main

// Control typedef for control
type Control struct {
	Location string
	Key      string
	Value    string
}

// Def typedef for vulnerability definition
type Def struct {
	Os          string
	Version     uint8
	Description string
	Control     Control
}
