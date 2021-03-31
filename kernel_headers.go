package qti_kernel_headers

import (
	"android/soong/android"
	"android/soong/cc"
	"fmt"
)

func init() {
	// Register our own module type
	android.RegisterModuleType("qti_kernel_headers_defaults", qtiKernelHeadersDefaultsFactory)
}

func qtiKernelHeadersDefaultsFactory() android.Module {
	module := cc.DefaultsFactory()

	// When the module is loaded, execute qtiKernelHeadersDefaults
	android.AddLoadHook(module, qtiKernelHeadersDefaults)

	return module
}

func qtiKernelHeadersDefaults(ctx android.LoadHookContext) {
	version := ctx.Config().VendorConfig("qti_kernel_headers").String("version")
	qtimodule := fmt.Sprintf("qti_kernel_headers_%s", version)

	p := struct {
		Export_header_lib_headers []string
		Header_libs []string
	}{
		[]string{qtimodule},
		[]string{qtimodule},
    }

	ctx.AppendProperties(&p)
}
