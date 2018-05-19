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
	err error
}

//Val returns an error
func (er ErrResult) Val() interface{} {
	return er.err
}

//NilResult is the nil result that does not contain any payload
type NilResult struct{}

//Val returns nil
func (nr NilResult) Val() interface{} {
	return nil
}

//HelpResult contains underlying command help message
type HelpResult struct {
	cmd Command
}

//Val returns underlying command help message
func (ur HelpResult) Val() interface{} {
	return ur.cmd.Help()
}

//StringResult contains a string
type StringResult struct {
	str string
}

//Val returns string
func (sr StringResult) Val() interface{} {
	return sr.str
}

//SliceResult contains a slice of strings
type SliceResult struct {
	val []string
}

//Val returns slice of strings
func (sr SliceResult) Val() interface{} {
	return sr.val
}
