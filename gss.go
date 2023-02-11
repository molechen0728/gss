package gss

// ?
func Ox3F(...error) {}

// @
func Ox40() {}

// $
func Ox24() {}

// #
func Ox23() {}

// E = p.fn(arg)  =========>  err1 := p.fn(arg)
// 							  if err1 != nil {
// 								return err1
// 							  }

var E error
