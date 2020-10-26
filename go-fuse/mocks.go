package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

const pollHackInode = ^uint64(0)

type Server struct {
	mountPoint string
	fileSystem RawFileSystem
	writeMu sync.Mutex
	mountFd int
	//latencies LatencyMap
	opts *MountOptions
	//buffers bufferPool
	reqPool sync.Pool
	readPool       sync.Pool
	reqMu          sync.Mutex
	reqReaders     int
	reqInflight    []*request
	//kernelSettings InitIn
	retrieveMu   sync.Mutex
	retrieveNext uint64
	//retrieveTab  map[uint64]*retrieveCacheRequest
	singleReader bool
	canSplice    bool
	loops        sync.WaitGroup
	ready chan error
	requestProcessingMu sync.Mutex
}

type request struct {
	inflightIndex int
	cancel chan struct{}
	interrupted bool
	inputBuf []byte
	inHeader *InHeader
	inData   unsafe.Pointer
	arg      []byte
	filenames []string
	//status   Status
	flatData []byte
	//fdData   *readResultFd
	//readResult ReadResult
	startTime time.Time
	//handler *operationHandler
	bufferPoolInputBuf  []byte
	bufferPoolOutputBuf []byte
	//outBuf [outputHeaderSize]byte
	smallInputBuf [128]byte
}

type _BatchForgetIn struct {
	InHeader
	Count uint32
	Dummy uint32
}

type InHeader struct {
	Length uint32
	Opcode uint32
	Unique uint64
	NodeId uint64
	//Caller
	Padding uint32
}

type _ForgetOne struct {
	NodeId  uint64
	Nlookup uint64
}

type MountOptions struct {
	AllowOther bool
	Options []string
	MaxBackground int
	MaxWrite int
	MaxReadAhead int
	IgnoreSecurityLabels bool
	RememberInodes bool
	FsName string
	Name string
	SingleThreaded bool
	DisableXAttrs bool
	Debug bool
	EnableLocks bool
	ExplicitDataCacheControl bool
}

type RawFileSystem interface {
	//String() string
	//SetDebug(debug bool)
	//Lookup(cancel <-chan struct{}, header *InHeader, name string, out *EntryOut) (status Status)
	Forget(nodeid, nlookup uint64)
	//GetAttr(cancel <-chan struct{}, input *GetAttrIn, out *AttrOut) (code Status)
	//SetAttr(cancel <-chan struct{}, input *SetAttrIn, out *AttrOut) (code Status)
	//Mknod(cancel <-chan struct{}, input *MknodIn, name string, out *EntryOut) (code Status)
	//Mkdir(cancel <-chan struct{}, input *MkdirIn, name string, out *EntryOut) (code Status)
	//Unlink(cancel <-chan struct{}, header *InHeader, name string) (code Status)
	//Rmdir(cancel <-chan struct{}, header *InHeader, name string) (code Status)
	//Rename(cancel <-chan struct{}, input *RenameIn, oldName string, newName string) (code Status)
	//Link(cancel <-chan struct{}, input *LinkIn, filename string, out *EntryOut) (code Status)
	//Symlink(cancel <-chan struct{}, header *InHeader, pointedTo string, linkName string, out *EntryOut) (code Status)
	//Readlink(cancel <-chan struct{}, header *InHeader) (out []byte, code Status)
	//Access(cancel <-chan struct{}, input *AccessIn) (code Status)
	//GetXAttr(cancel <-chan struct{}, header *InHeader, attr string, dest []byte) (sz uint32, code Status)
	//ListXAttr(cancel <-chan struct{}, header *InHeader, dest []byte) (uint32, Status)
	//SetXAttr(cancel <-chan struct{}, input *SetXAttrIn, attr string, data []byte) Status
	//RemoveXAttr(cancel <-chan struct{}, header *InHeader, attr string) (code Status)
	//Create(cancel <-chan struct{}, input *CreateIn, name string, out *CreateOut) (code Status)
	//Open(cancel <-chan struct{}, input *OpenIn, out *OpenOut) (status Status)
	//Read(cancel <-chan struct{}, input *ReadIn, buf []byte) (ReadResult, Status)
	//Lseek(cancel <-chan struct{}, in *LseekIn, out *LseekOut) Status
	//GetLk(cancel <-chan struct{}, input *LkIn, out *LkOut) (code Status)
	//SetLk(cancel <-chan struct{}, input *LkIn) (code Status)
	//SetLkw(cancel <-chan struct{}, input *LkIn) (code Status)
	//Release(cancel <-chan struct{}, input *ReleaseIn)
	//Write(cancel <-chan struct{}, input *WriteIn, data []byte) (written uint32, code Status)
	//CopyFileRange(cancel <-chan struct{}, input *CopyFileRangeIn) (written uint32, code Status)
	//Flush(cancel <-chan struct{}, input *FlushIn) Status
	//Fsync(cancel <-chan struct{}, input *FsyncIn) (code Status)
	//Fallocate(cancel <-chan struct{}, input *FallocateIn) (code Status)
	//OpenDir(cancel <-chan struct{}, input *OpenIn, out *OpenOut) (status Status)
	//ReadDir(cancel <-chan struct{}, input *ReadIn, out *DirEntryList) Status
	//ReadDirPlus(cancel <-chan struct{}, input *ReadIn, out *DirEntryList) Status
	//ReleaseDir(input *ReleaseIn)
	//FsyncDir(cancel <-chan struct{}, input *FsyncIn) (code Status)
	//StatFs(cancel <-chan struct{}, input *InHeader, out *StatfsOut) (code Status)
	//Init(*Server)
}

type MyFileSystem struct {}

func (mfs MyFileSystem) Forget(nodeid, nlookup uint64) {
	fmt.Printf("Calling Forget with %d, %d\n", nodeid, nlookup)
}