package ecode

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync/atomic"
	"time"
)

type ecodes struct {
	codes atomic.Value
}

type data struct {
	Code    map[int]string
	Version int64
}

var (
	defualtEcodes = &ecodes{}
)

// Init init ecode from ecode service
func Init() {
	defualtEcodes.codes.Store(make(map[int]string))
	ver, _ := defualtEcodes.update(0)
	go defualtEcodes.updateproc(ver)
}

/**
  get new mapping
*/
func (e *ecodes) updateproc(lastVer int64) {
	var (
		ver int64
		err error
	)
	for {
		if ver, err = e.update(lastVer); err != nil {
			log.Printf("error occur %s ", err.Error())
		}
		lastVer = ver
		time.Sleep(time.Minute * 15)
	}
}

func (e *ecodes) update(ver int64) (lver int64, err error) {
	var (
		d     = &data{}
		bytes []byte
	)
	// mock json
	// TODO read error code mapping from backend
	if bytes, err = ioutil.ReadFile("./c.json"); err != nil {
		return
	}
	if err = json.Unmarshal(bytes, d); err != nil {
		return
	}
	oCodes, ok := e.codes.Load().(map[int]string)
	if !ok {
		return
	}
	nCodes := copy(oCodes)
	// merge new code mapping
	for k, v := range d.Code {
		nCodes[k] = v
	}
	// update ecode mapping
	Register(nCodes)
	// save the mapping in local
	e.codes.Store(nCodes)
	return d.Version, nil
}

/**
copy atomic.Value to new map
*/
func copy(src map[int]string) (dst map[int]string) {
	dst = make(map[int]string)
	for k1, v1 := range src {
		dst[k1] = v1
	}
	return
}
