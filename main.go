package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
mytens new task TBI-1234
	git checkout develop
	git checkout -b task/TBI-1234
*/
/*
mytens commit TBI-1234 -m "[Module] Commit message"
	git add .
	git commit -m "$(m) #TBI-1234"
*/
/*
mytens push task TBI-1234
	git checkout task/TBI-1234
	git push origin task/TBI-1234
	git checkout develop
	git pull upstream develop
	git merge task/TBI-1234
	git push origin develop
	https://gitlab.playcourt.id/haqiramadhani/service-api-data/-/merge_requests/new?merge_request%5Bsource_branch%5D=develop&merge_request%5Btarget_branch%5D=develop
*/
/*
mytens new bugfix TBI-1234
	git checkout develop
	git checkout -b bugfix/TBI-1234
*/
/*
mytens push bugfix TBI-1234
	git checkout task/TBI-1234
	git merge bugfix/TBI-1234
	git push origin task/TBI-1234
	git checkout develop
	git pull upstream develop
	git merge task/TBI-1234
	git push origin develop
	git branch -D bugfix/TBI-1234
	https://gitlab.playcourt.id/haqiramadhani/service-api-data/-/merge_requests/new?merge_request%5Bsource_branch%5D=develop&merge_request%5Btarget_branch%5D=develop
*/
/*
mytens new hotfix v1.0.1
	git checkout master
	git checkout -b hotfix/v1.0.1
*/
/*
mytens push hotfix v1.0.1
	git checkout release
	git pull upstream release
	git merge hotfix/v1.0.1
	git push origin release
	git checkout hotfix/v1.0.1
	https://gitlab.playcourt.id/haqiramadhani/service-api-data/-/merge_requests/new?merge_request%5Bsource_branch%5D=release&merge_request%5Btarget_branch%5D=release
*/
/*
mytens finish hotfix v1.0.1
	npm i -g auto-changelog
	git checkout hotfix/v1.0.1
	git tag v1.0.1
	auto-changelog
	git add .
	git commit -m "Update changelog v1.0.1"
	git tag -d v1.0.1
	git tah v1.0.1
	git checkout master
	git pull upstream master
	git merge v1.0.1
	git push origin v1.0.1
	git push upstream v1.0.1
	git push origin master
	git branch -D hotfix/v1.0.1
	git checkout develop
	https://gitlab.playcourt.id/haqiramadhani/service-api-data/-/merge_requests/new?merge_request%5Bsource_branch%5D=master&merge_request%5Btarget_branch%5D=master
*/
/*
mytens new release v1.1.0 -b TBI-1234,TBI-1235,TBI-1236
	git checkout master
	git pull upstream master
	git checkout -b release/v1.1.0
		git merge develop (without -b flag)
	git merge TBI-1234 (with -b flag separated by comma)
	git merge TBI-1235 (with -b flag separated by comma)
	git merge TBI-1236 (with -b flag separated by comma)
	git checkout release
	git pull upstream release
	git merge release/v1.1.0
	git push origin release
	git branch -D TBI-1234 (with -b flag separated by comma)
	git branch -D TBI-1235 (with -b flag separated by comma)
	git branch -D TBI-1236 (with -b flag separated by comma)
	git checkout release/v1.1.0
	https://gitlab.playcourt.id/haqiramadhani/service-api-data/-/merge_requests/new?merge_request%5Bsource_branch%5D=release&merge_request%5Btarget_branch%5D=release
*/
/*
mytens push release v1.1.0
	git checkout release
	git pull upstream release
	git merge release/v1.1.0
	git push origin release
	git checkout release/v1.1.0
	https://gitlab.playcourt.id/haqiramadhani/service-api-data/-/merge_requests/new?merge_request%5Bsource_branch%5D=release&merge_request%5Btarget_branch%5D=release
*/
/*
mytens finish release v1.1.0
	npm i -g auto-changelog
	git checkout release/v1.1.0
	git tag v1.1.0
	auto-changelog
	git add .
	git commit -m "Update changelog v1.1.0"
	git tag -d v1.1.0
	git tah v1.1.0
	git checkout master
	git pull upstream master
	git merge v1.1.0
	git push origin v1.1.0
	git push upstream v1.1.0
	git push origin master
	git branch -D hotfix/v1.1.0
	git checkout develop
	https://gitlab.playcourt.id/haqiramadhani/service-api-data/-/merge_requests/new?merge_request%5Bsource_branch%5D=master&merge_request%5Btarget_branch%5D=master
*/

func execGit(command, path, message string) error {
	fmt.Println(message)
	cmd := exec.Command("git", strings.Split(command, " ")...)
	cmd.Dir = path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func execGitFormatting(command, path, message, override string) error {
	if strings.Contains(command, "#") {
		fmt.Println(message)
		return execCmd(command, path)
	} else if strings.Contains(message, "%") {
		fmt.Println(fmt.Sprintf(message, override))
	} else {
		fmt.Println(message)
	}
	var args []string
	for _, s := range strings.Split(command, " ") {
		if strings.Contains(s, "%") {
			args = append(args, fmt.Sprintf(s, override))
		} else {
			args = append(args, s)
		}
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = path
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func execCmd(command, path string) error {
	command = strings.Replace(command, "#", "", 1)
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Path = path
	return cmd.Run()
}

func genMRUrl(branch string) string {
	repoUrl := "https://gitlab.playcourt.id/haqiramadhani/service-api-data"
	return fmt.Sprintf("%s/-/merge_requests/new?merge_request[source_branch]=%s&merge_request[target_branch]=%s", repoUrl, branch, branch)
}

func requiredParam(param, message string) bool {
	if param == "" {
		log.Println(color.Colorize(color.Red, message))
		return true
	}
	return false
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	app := &cli.App{
		Name:        "mytens",
		Usage:       "Simple command to implement git flow and other command when developing MyTEnS, especially service-api-data",
		UsageText:   "mytens new task TBI-1234\nmytens commit TBI-1234 -m \"[Module] Commit message\"\nmytens push task TBI-1234\nmytens new bugfix TBI-1234\nmytens push bugfix TBI-1234\nmytens new hotfix v1.0.1\nmytens push hotfix v1.0.1\nmytens finish hotfix v1.0.1\nmytens new release v1.1.0 -b TBI-1234,TBI-1235,TBI-1236\nmytens push release v1.1.0\nmytens finish release v1.1.0",
		Description: fmt.Sprintf("You can use %s (without -b flag) for release from branch develop", color.Colorize(color.Blue, "mytens new release v1.1.0")),
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"g", "n"},
				Usage:   "Starting new action",
				Subcommands: []*cli.Command{
					{
						Name:        "task",
						Usage:       "Creating new task",
						Description: "- TBI-1234 is a task name, you can use backlog code.\n\nAction:\n\tgit checkout develop\n\tgit checkout -b task/TBI-1234",
						UsageText:   "mytens new task TBI-1234",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							if requiredParam(name, "Task name is required") {
								return nil
							}
							fmt.Printf("Creating task: %s\n", name)
							commands := []string{
								"checkout develop",
								"checkout -b task/" + name,
							}
							for _, s := range commands {
								err := execGit(s, path, "git "+s)
								if err != nil {
									return err
								}
							}
							return nil
						},
					},
					{
						Name:        "bugfix",
						Description: "",
						Category:    "bugfix",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							fmt.Printf("Preparing bugfix: %s\n", name)
							commands := []string{
								"checkout develop",
								"checkout -b bugfix/" + name,
							}
							for _, s := range commands {
								err := execGit(s, path, "git "+s)
								if err != nil {
									return err
								}
							}
							return nil
						},
					},
					{
						Name:        "hotfix",
						Description: "",
						Category:    "hotfix",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							fmt.Printf("Preparing hotfix: %s\n", name)
							commands := []string{
								"checkout develop",
								"checkout -b hotfix/" + name,
							}
							for _, s := range commands {
								err := execGit(s, path, "git "+s)
								if err != nil {
									return err
								}
							}
							return nil
						},
					},
					{
						Name:        "release",
						Description: "",
						Category:    "release",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "branches", Aliases: []string{"b"}},
						},
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							branchesValue := context.Value("branches")
							if branchesValue == "" {
								branchesValue = "develop"
							}
							branches := strings.Split(fmt.Sprintf("%v", branchesValue), ",")
							fmt.Printf("Publishing %v to release: %s\n", branches, name)
							commands := []string{
								"checkout master",
								"pull upstream master",
								"checkout -b release/" + name,
							}
							for _, branch := range branches {
								commands = append(commands, "merge "+branch)
							}
							commands = append(commands, "checkout release")
							commands = append(commands, "pull upstream release")
							commands = append(commands, "merge release/"+name)
							commands = append(commands, "push origin release")
							commands = append(commands, "checkout release/"+name)
							for _, branch := range branches {
								commands = append(commands, "branch -D "+branch)
							}
							for _, s := range commands {
								err := execGit(s, path, "git "+s)
								if err != nil {
									return err
								}
							}
							fmt.Println(genMRUrl("release"))
							return nil
						},
					},
				},
			},
			{
				Name:    "push",
				Aliases: []string{"p"},
				Usage:   "Publish to remote branch",
				Subcommands: []*cli.Command{
					{
						Name:        "task",
						Description: "",
						Category:    "task",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							fmt.Printf("Preparing hotfix: %s\n", name)
							commands := []string{
								"checkout task/" + name,
								"push origin task/" + name,
								"checkout develop",
								"pull upstream develop",
								"merge task/" + name,
								"push origin develop",
							}
							for _, s := range commands {
								err := execGit(s, path, "git "+s)
								if err != nil {
									return err
								}
							}
							fmt.Println(genMRUrl("develop"))
							return nil
						},
					},
					{
						Name:        "bugfix",
						Description: "",
						Category:    "bugfix",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							fmt.Printf("Preparing hotfix: %s\n", name)
							commands := []string{
								"checkout task/" + name,
								"merge bugfix/" + name,
								"push origin task/" + name,
								"checkout develop",
								"pull upstream develop",
								"merge task/" + name,
								"push origin develop",
								"branch -D bugfix/" + name,
							}
							for _, s := range commands {
								err := execGit(s, path, "git "+s)
								if err != nil {
									return err
								}
							}
							fmt.Println(genMRUrl("develop"))
							return nil
						},
					},
					{
						Name:        "hotfix",
						Description: "",
						Category:    "hotfix",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							fmt.Printf("Preparing hotfix: %s\n", name)
							commands := []string{
								"checkout release",
								"pull upstream release",
								"merge hotfix/" + name,
								"push origin release",
								"checkout hotfix/" + name,
							}
							for _, s := range commands {
								err := execGit(s, path, "git "+s)
								if err != nil {
									return err
								}
							}
							fmt.Println(genMRUrl("release"))
							return nil
						},
					},
					{
						Name:        "release",
						Description: "",
						Category:    "release",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							fmt.Printf("Preparing hotfix: %s\n", name)
							commands := []string{
								"checkout release",
								"pull upstream release",
								"merge release/" + name,
								"push origin release",
								"checkout release/" + name,
							}
							for _, s := range commands {
								err := execGit(s, path, "git "+s)
								if err != nil {
									return err
								}
							}
							fmt.Println(genMRUrl("release"))
							return nil
						},
					},
				},
			},
			{
				Name:    "finish",
				Aliases: []string{"f"},
				Usage:   "Start build for production",
				Subcommands: []*cli.Command{
					{
						Name:        "hotfix",
						Description: "",
						Category:    "hotfix",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							message := "Update changelog " + name
							fmt.Printf("Preparing hotfix: %s\n", name)
							commands := []string{
								"#npm i -g auto-changelog",
								"checkout release/" + name,
								"tag " + name,
								"#auto-changelog",
								"add -A",
								"commit -m %q",
								"tag -d " + name,
								"tag " + name,
								"checkout master",
								"pull upstream master",
								"merge " + name,
								"push origin " + name,
								"push upstream " + name,
								"push origin master",
								"branch -D hotfix/" + name,
								"checkout develop",
							}
							for _, s := range commands {
								err := execGitFormatting(s, path, "git "+s, message)
								if err != nil {
									return err
								}
							}
							fmt.Println(genMRUrl("master"))
							return nil
						},
					},
					{
						Name:        "release",
						Description: "",
						Category:    "release",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
							message := "Update changelog " + name
							fmt.Printf("Preparing hotfix: %s\n", name)
							commands := []string{
								"#npm i -g auto-changelog",
								"checkout release/" + name,
								"tag " + name,
								"#auto-changelog",
								"add .",
								"commit -m %q",
								"tag -d " + name,
								"tah " + name,
								"checkout master",
								"pull upstream master",
								"merge " + name,
								"push origin " + name,
								"push upstream " + name,
								"push origin master",
								"branch -D hotfix/" + name,
								"checkout develop",
							}
							for _, s := range commands {
								err := execGitFormatting(s, path, "git "+s, message)
								if err != nil {
									return err
								}
							}
							fmt.Println(genMRUrl("master"))
							return nil
						},
					},
				},
			},
			{
				Name:    "commit",
				Aliases: []string{"c"},
				Usage:   "Commit changes",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
				},
				Action: func(context *cli.Context) error {
					message := fmt.Sprintf("%v", context.Value("message"))
					fmt.Printf("Commiting: %q\n", message)
					commands := []string{
						"add -A",
						"commit -m %q",
					}
					for _, s := range commands {
						err := execGitFormatting(s, path, "git "+s, message)
						if err != nil {
							return err
						}
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
