# go-control-plane Fork
This is a fork of [go-control-plane](https://github.com/envoyproxy/go-control-plane). Before introducing any changes here, please create an issue or submit a change to the upstream repository, as it might impact other users.

## How to Contribute
In this fork, we maintain two branches:

* `main`: This branch is synced daily with the upstream repository via an automated action. The only custom addition is the job responsible for synchronization.
* `release`: This branch contains all user-introduced changes. It diverges from the stable version to ensure no unreleased commits are included. Any changes you make should target this branch.

## Steps to Contribute

1. Check out the `release` branch.
2. Make your changes and create a pull request targeting the `release` branch.
3. Once your pull request is reviewed, merge it into the `release` branch.


## Releasing
For standard releases, the `Sync with Upstream` action handles synchronization and creates custom versions whenever there’s a new upstream release. We fetch upstream tags, so, for example, `v0.13.1` in this fork corresponds directly to the same version in the upstream repository.

Our job appends a custom suffix (`-kong-1`) to the upstream tag. If you merge a change and want to create a release, you should tag your commit using the following pattern:

```bash
<current-version>-kong-<incremented-number>
```

Example

The current tag is `v0.13.1-kong-1`. After merging your changes, you want to release a new version. You should tag your commit as:

`v0.13.1-kong-2`  

After tagging, push the commit to the repository to finalize the release.

## Conflict resolution

It might happen that automatic release job can hit a conflict. In this situation maintainer needs to resolve this manually:

Instruction

1. Fetch current `main` and `release` branch
```bash
git fetch origin --tags
```
2. Update current main branch with origin
```bash
git reset --hard origin/main
```
3. Change to the release branch
```bash
git checkout release
git reset --hard origin/release
```
4. Remove Untracked Files and Directories
```bash
git clean -fd
```
5. Get new tag value
```bash
./.github/scripts/get-version-tags.sh
```

Output:
```bash
Current release tag prefix: v0.13.1
Upstream tag: v0.13.2
New tag: v0.13.2-kong-1
```

if tags are equal, you are going to see message `"Tags are equal, no need to release"` and you don't need to release/rebase anything.

6. Rebase release branch between tags:
```bash
git rebase --onto  <Upstream tag> <Current release tag prefix> release
```
7. If you encounter git conflicts, resolve them, and follow git instruction
* resolve conflicts
* git add <files>
* git rebase --continue

8. Tag the commit
```bash
git tag <New tag>
```
9. Push changes to origin
```bash
git push origin release --tags --force-with-lease
```
