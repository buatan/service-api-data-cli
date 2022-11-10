package main

import (
	"fmt"
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
mytens commit TBI-1234 -m ""
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
mytens finish bugfix TBI-1234
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
	fmt.Println(message)
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

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	app := &cli.App{
		Name:  "mytens",
		Usage: "Simple command to implement git flow and other command when developing MyTEnS, especially service-api-data",
		Commands: []*cli.Command{
			{
				Name: "new",
				Subcommands: []*cli.Command{
					{
						Name:     "task",
						Category: "task",
						Action: func(context *cli.Context) error {
							name := context.Args().Get(0)
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
						Name:     "bugfix",
						Category: "bugfix",
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
						Name:     "hotfix",
						Category: "hotfix",
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
						Name:     "release",
						Category: "release",
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
							return nil
						},
					},
				},
			},
			{
				Name: "push",
				Subcommands: []*cli.Command{
					{
						Name:     "task",
						Category: "task",
					},
					{
						Name:     "bugfix",
						Category: "bugfix",
					},
					{
						Name:     "hotfix",
						Category: "hotfix",
					},
					{
						Name:     "release",
						Category: "release",
					},
				},
			},
			{
				Name: "finish",
				Subcommands: []*cli.Command{
					{
						Name:     "hotfix",
						Category: "hotfix",
					},
					{
						Name:     "release",
						Category: "release",
					},
				},
			},
			{
				Name: "commit",
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
