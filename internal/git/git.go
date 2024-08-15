package git

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	git_config "github.com/go-git/go-git/v5/config"
)

func GitUser(proj_path string) string {
	// TODO try checking the local project directory .git and get a user from that
	var gitCfg *git_config.Config

	gitRepo, err := git.PlainOpen(proj_path)
	if err != nil {
		color.Red("git is not inited for this is the project")
		return "User undefined : Email undefined"
	}

	gitCfg, err = gitRepo.ConfigScoped(git_config.LocalScope)
	if err != nil {
		color.Red("git scope error " + err.Error())
	}

	if gitCfg.User.Name != "" {
		return fmt.Sprintf("%s : %s", gitCfg.User.Name, gitCfg.User.Email)
	}
	color.Red("git user is not found in the local scope. Try checking global scope")

	gitCfg, err = gitRepo.ConfigScoped(git_config.GlobalScope)
	if err != nil {
		color.Red("git user is not found in the global scope")
		return "User undefined : Email undefined"
	}

	return fmt.Sprintf("%s : %s", gitCfg.User.Name, gitCfg.User.Email)

}
