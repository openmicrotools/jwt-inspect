# Contributing

## Asking for a feature or reporting a bug
Please add an to our issue to our existing issues for the repository.

For bugs include:
- The issue (why is this a bug)
- What the expected behavior should be
- What the actual behavior is
- How to reproduce the bug
- Steps to fix/mitigate (optional!)
- Expected level of effort (LOE) (optional!)

For feature requests include:
- The feature being requested with a detailed description
- A description of why the feature is needed (user stories)
- Any ideas at a possible solution (optional!)
- LOE (optional!)
- Any additional information to help understand the ask (optional!)

## Working on the code base

We use a fork and merge method for committing changes.
1. Fork the project to your own github account
1. Clone the repo to your local computer
1. Set the upstream for main
    - ```remote add upstream <upstream repository>```
    - confirm it was added with ```git remote -v```. You should see both origin and upstream
1. Create a new branch on fork to do work
1. Push to branch on own fork once work is committed/ready
1. Open a PR for the change to merge with main (or a branch) on the upstream repo
1. PR's must pass tests and be approved by a maintainer.
1. Get feedback, and once approved the branch from your fork can be merged into main in the upstream repository.

When your fork becomes out of sync with the upstream
1. From main of your fork, run ```git fetch upstream```
1. Run ```git rebase --ff upstream/main``` to pull in the changes
    - If your branch needs to be brought up to sync with changes in the upstream
        1. Complete rebasing on your fork to have main at parity with upstream/main
        1. ```git checkout <branch>```
        1. ```git merge --ff main```
        1. Resolve any merge conflicts
        1. ```git push``` (optional)
