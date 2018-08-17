package command

//Result contains the result of the command execution
type Result interface {
	Val() interface{}
}

//OkResult is the success result that does not contain any payload
type OkResult struct{}

//Val returns a nil for success result
func (or OkResult) Val() interface{} {
	return nil
}

//ErrResult contains an error
type ErrResult struct {
	Value error
}

//Val returns an error
func (er ErrResult) Val() interface{} {
	return er.Value
}

//NilResult is the nil result that does not contain any payload
type NilResult struct{}

//Val returns nil
func (nr NilResult) Val() interface{} {
	return nil
}

//HelpResult contains command's help message
type HelpResult struct {
	Value string
}

//Val returns underlying command help message
func (ur HelpResult) Val() interface{} {
	return ur.Value
}

//StringResult contains a string
type StringResult struct {
	Value string
}

//Val returns string
func (sr StringResult) Val() interface{} {
	return sr.Value
}

//IntResult containt an int64 integer
type IntResult struct {
	Value int64
}

//Val returns an int64 integer
func (ir IntResult) Val() interface{} {
	return ir.Value
}

//SliceResult contains a slice of strings
type SliceResult struct {
	Value []string
}

//Val returns slice of strings
func (sr SliceResult) Val() interface{} {
	return sr.Value
}
