package gss

// ?
func Ox3F(...error) {}

// @
func Ox40(FuncMetaData) {}

// $
func Ox24(map[string]StructMetaData) {}

// #
func Ox23() {}

// E = p.fn(arg)  =========>  err1 := p.fn(arg)
// 							  if err1 != nil {
// 								return err1
// 							  }

var E error

type FuncMetaData struct {
	Name          string
	Requires      []string
	Hidden        bool
	DelayInmillis int
}

type StructMetaData struct {
	Name    string
	No      int
	Mutable bool
}

type IfErrorNotNilReturn = error
