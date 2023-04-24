package entity

type Relation struct {
	masterID int
	slaveID  int
}

func NewRelation(masterID int, slaveID int) *Relation {
	return &Relation{masterID: masterID, slaveID: slaveID}
}
