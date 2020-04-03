package ecode

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"sync/atomic"
)

var (
	_messages atomic.Value
	_codes    = map[int]struct{}{}
)

func Register(cm map[int]string) {
	_messages.Store(cm)
}

func New(e int) Code {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

func String(e string) Code {
	if e == "" {
		return OK
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}
	return Code(i)
}

// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}
func Int(i int) Code { return Code(i) }

func add(e int) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return Int(e)
}

type Codes interface {
	Error() string
	Code() int
	Message() string
}

type Code int

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e Code) Code() int { return int(e) }

func (e Code) Message() string {
	if cm, ok := _messages.Load().(map[int]string); ok {
		if msg, ok := cm[e.Code()]; ok {
			return msg
		}
	}
	return e.Error()
}
