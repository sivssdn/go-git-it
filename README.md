# go-git-it

A command line tool to add features around git. The aim is to reduce the verbosity of common git command while making it more customizable.

### Setting up for use
For mac users, cloning the repo and setting up an alias for the build file (already included in build folder :)) works.
Setting up alias:
- Open `~/.bashrc` [or ~/.zshrc depending upon your terminal]
- Paste `alias your_preferred_program_name="/location/to/go-git-it/build/main"`

### Building for your os
After having working installation of go locally, follow the steps:
- cd into the `src` folder of the project
- Run `go build -o ../build/main`. This makes your new build as `main` inside folder build.

### Features
- **Counting Bitbucket commits into Git**: Not only bitbucket, practically everywhere where you can use `git commit`. This feature helps you counting saving your commit from bitbucket with original commit message, without original code. To set up this feature, follow the steps:
  - Make a dummy git repo.
  - `git init` inside your new dummy repo
  - Add absolute path to that repo in the .env file (.env-example attached)
Every time you commit using go-git-it, your dummy-repo is updated with the latest commit. To upload it to github, go to your dummy repo and push to remote.

- **Command Alias** : Alias for common git commands can be set with `alias.json` inside config folder. A single command can have multiple aliases.
To print a list of all `Alias` commands, command `names` is used. E.g., `ggt names`

- **Branch Name Autocomplete** : Is used in conjunction with other commands like `checkout`. Primary function is to find a branch name with a given identifier. Auto-complete is case in-sensitive.

- **Checkout** : Checks out to a branch with just a substring identifier.
 Stashes the current branch and switches to the new one. 

- **Push** : Pushes to remote taking care of setting upstream. Just simple `push` command for the user.

- **Pull** : Pull from the remote upstream of the current branch without setting up upstream or specifying current branch name. Stashes the changes on local branch before pulling.

