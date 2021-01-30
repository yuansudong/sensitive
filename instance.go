package sensitive

import (
	"sync/atomic"
	"unsafe"
)

var _Instance *Engine

const _Ignore = '*'

// LoadWordsHandler 加载敏感词的
type LoadWordsHandler func() []string

// Init 用于初始化资源
func Init(handler LoadWordsHandler) {
	_Instance = _New().LoadFromList(handler())
}

// Get 用于获得引擎
func Get() *Engine {
	return (*Engine)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&_Instance)),
	))
}

// Update 用于更新敏感词
func Update(handler LoadWordsHandler) {
	if _Instance != nil {
		_TmpNew := _New().LoadFromList(handler())
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&_Instance)), unsafe.Pointer(_TmpNew))
	}
}

// Release 用于释放资源
func Release() {
	_Instance = nil
}
