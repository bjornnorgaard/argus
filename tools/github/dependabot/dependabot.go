package dependabot

import (
	"github.com/bjornnorgaard/argus/pkg/say"
	"github.com/bjornnorgaard/argus/utils/pulls"
)

func Run() {
	say.Eyes("Loading Dependabot pull requests...")
	prs, err := pulls.Get()

	if err != nil {
		say.Error(err.Error())
	}

	for _, p := range prs {
		say.Check(p.RepositoryWithOwner, p.Title)
	}
}
