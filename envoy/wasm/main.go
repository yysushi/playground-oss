package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

func main() {
	proxywasm.SetNewRootContext(func (contextID uint32) proxywasm.RootContext{
		return &sample{contextID: contextID}
	})
}

type sample struct {
	proxywasm.DefaultRootContext
	contextID uint32
}

// override
func (ctx *sample) OnVMStart(vmConfigurationSize int) bool {
	proxywasm.LogInfo("ʕ◔ϖ◔ʔ")
	return true
}
