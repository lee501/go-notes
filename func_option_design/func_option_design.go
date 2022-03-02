package func_option_design

type generatorOpt struct {
	useAllExportedFields bool
	throwErrorOnCycle    bool
	schemaCustomizer     SchemaCustomizerFn
}

type Option func(opt *generatorOpt)

//callback func
type SchemaCustomizerFn func(name string) error

func UseAllExportedFields() Option {
	return func(opt *generatorOpt) {
		opt.useAllExportedFields = false
	}
}

func ThrowErrorOnCycle() Option {
	return func(opt *generatorOpt) {
		opt.throwErrorOnCycle = true
	}
}

func schemaCustomizer(sc SchemaCustomizerFn) Option {
	return func(opt *generatorOpt) {
		opt.schemaCustomizer = sc
	}
}

//构造struct generatorOpt
func GenerateOpt(ops ...Option) *generatorOpt {
	gopt := &generatorOpt{}
	for _, fn := range ops {
		fn(gopt)
	}
	return gopt
}
