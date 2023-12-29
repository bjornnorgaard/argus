package github

import (
	"github.com/bjornnorgaard/argus/pkg/say"
	"github.com/cli/go-gh/v2"
)

func Run() {
	exec, b, err := gh.Exec("auth", "status")
	if err != nil {
		say.Error(b.String())
	}

	say.Check(exec.String())
}
