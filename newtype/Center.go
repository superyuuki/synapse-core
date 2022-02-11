package newtype

//Center is where Runs come from.
type Center interface {
}

//Holder is where runs are stored.
type Holder interface {
}

//Run is a ticking threadpool that can cancel a holder
type Run interface {
}
