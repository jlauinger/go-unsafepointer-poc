package main

import (
	"log"
	"reflect"
	"unsafe"
)

// doBatchForget - forget a list of NodeIds
func doBatchForget(server *Server, req *request) {
	in := (*_BatchForgetIn)(req.inData)
	wantBytes := uintptr(in.Count) * unsafe.Sizeof(_ForgetOne{})
	if uintptr(len(req.arg)) < wantBytes {
		// We have no return value to complain, so log an error.
		log.Printf("Too few bytes for batch forget. Got %d bytes, want %d (%d entries)",
			len(req.arg), wantBytes, in.Count)
	}

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&req.arg[0])),
		Len:  int(in.Count),
		Cap:  int(in.Count),
	}

	forgets := *(*[]_ForgetOne)(unsafe.Pointer(h))
	for i, f := range forgets {
		if server.opts.Debug {
			log.Printf("doBatchForget: rx %d %d/%d: FORGET i%d {Nlookup=%d}",
				req.inHeader.Unique, i+1, len(forgets), f.NodeId, f.Nlookup)
		}
		if f.NodeId == pollHackInode {
			continue
		}
		server.fileSystem.Forget(f.NodeId, f.Nlookup)
	}
}
