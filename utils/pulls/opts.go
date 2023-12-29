package pulls

type OptsApply func(o *optsType)

type optsType struct {
	flags map[string]string
}

var optsDefault = optsType{
	flags: map[string]string{},
}

func WithStateOpen() OptsApply {
	return func(o *optsType) {
		o.flags["state"] = "--state=open"
	}
}
