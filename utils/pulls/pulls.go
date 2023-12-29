package pulls

import (
	"fmt"
	"github.com/bjornnorgaard/argus/pkg/say"
	"github.com/bjornnorgaard/argus/types"
	"github.com/cli/go-gh/v2"
)

func Get(appliers ...OptsApply) ([]types.PR, error) {
	opts := &optsDefault
	for _, apply := range appliers {
		apply(opts)
	}

	stdout, stderr, err := gh.Exec("search", "prs",
		fmt.Sprintln(opts.state),
	)
	if err != nil {
		say.Error(stderr.String())
	}

	say.Check(stdout.String())
	return nil, nil
}
