package command

//Reply contains the result of the command execution.
type Reply interface {
	Val() interface{}
}

//OkReply is the success reply that does not contain any payload.
type OkReply struct{}

//Val returns a nil for success result.
func (or OkReply) Val() interface{} {
	return nil
}

//ErrReply contains an error.
type ErrReply struct {
	Value error
}

//Val returns an error.
func (er ErrReply) Val() interface{} {
	return er.Value
}

//NilReply is the nil result that does not contain any payload.
type NilReply struct{}

//Val returns nil.
func (nr NilReply) Val() interface{} {
	return nil
}

//RawStringReply contain a string that should not be formatted.
type RawStringReply struct {
	Value string
}

//Val returns underlying command help message.
func (ur RawStringReply) Val() interface{} {
	return ur.Value
}

//StringReply contains a string.
type StringReply struct {
	Value string
}

//Val returns string.
func (sr StringReply) Val() interface{} {
	return sr.Value
}

//IntReply contains an int64 integer.
type IntReply struct {
	Value int64
}

//Val returns an int64 integer.
func (ir IntReply) Val() interface{} {
	return ir.Value
}

//SliceReply contains a slice of strings.
type SliceReply struct {
	Value []string
}

//Val returns slice of strings.
func (sr SliceReply) Val() interface{} {
	return sr.Value
}
