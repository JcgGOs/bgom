<!-- 
title:git cheatsheet
summary: introduce git
tag: java,python,golang
slug: git-cheatsheet
Time: 2022-05-13
-->

# git cheatsheet

## Creating Repositories
```shell
# create new repository in current directory
git init

# clone a remote repository
git clone [url]
# for example cloning the entire jquery repo locally
git clone https://github.com/jquery/jquery
```

## Branches and Tags
```shell
# List all existing branches with the latest commit comment 
git branch –av

# Switch your HEAD to branch
git checkout [branch]

# Create a new branch based on your current HEAD
git branch [new-branch]

# Create a new tracking branch based on a remote branch
git checkout --track [remote/branch]
# for example track the remote branch named feature-branch-foo
git checkout --track origin/feature-branch-foo

# Delete a local branch
git branch -d [branch]

# Tag the current commit
git tag [tag-name]
```

## Local Changes
```shell
# List all new or modified files - showing which are to staged to be commited and which are not 
git status

# View changes between staged files and unstaged changes in files
git diff

# Add all current changes to the next commit
git add [file]

# Commit all local changes in tracked  files
git commit -m "An inline  commit message"

# Commit previously staged changes
git commit
git commit -m "An inline commit message"

# Unstages the file, but preserve its contents

git reset --soft xxxxx
git reset --hard xxxxx
```

## Update and Publish
```shell
# List all remotes 
git remote -v

# Add a new remote at [url] with the given local name
git remote add [localname] [url]

# Download all changes from a remote, but don‘t integrate into them locally
git fetch

# Download all remote changes and merge them locally
git pull 

# Publish your tags to a remote
git push 
git push -f
```

## Merge & Rebase
```shell
# Merge [branch] into your current HEAD 
git merge [branch]

# Rebase your current HEAD onto [branch]
git rebase -i [branch]

# Abort a rebase 
git rebase –abort

# Continue a rebase after resolving conflicts 
git rebase –continue

```