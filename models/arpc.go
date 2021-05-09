package models

// Arg aRPC arg struct
type Arg struct {
	Name 	 string
	TypeName string
}

// Procedure aRPC procedure struct
type Procedure struct {
	Name   string
	Arg	   *Arg
	Result *string
}
