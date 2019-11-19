# go-git-it

A command line tool to add features around git. The aim is to reduce the verbosity of common git command while making it more customizable.

### Features
- **Command Alias** : Alias for common git commands can be set inside `commands.go`. A single command can have multiple aliases.

- **Branch Name Autocomplete** : Is used in conjunction with other commands like `checkout`. Primary function is to find a branch name with a given identifier.

- **Checkout** : Checks out to a branch with just a substring identifier.
 Stashes the current branch and switches to the new one. 

- **Push** : Pushes to remote taking care of setting upstream. Just simple `push` command for the user.

- **Pull** : Pull from the remote upstream of the current branch without setting up upstream or specifying current branch name.