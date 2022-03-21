# hcr
Helm chart releaser is a tool to help to host helm repository as GitHub page, where packaged chart is uploaded as
a release asset and the index file is hosted on GitHub page. This tool is very similar to
[helm chart releaser](https://github.com/helm/chart-releaser), but is simpler and works with private repos as well.

Projects are expected to have empty GitHub pages branch already created, run the following to set up the branch:
```shell
git checkout --orphan gh-pages
git rm -rf .
git commit -m "initial commit" --allow-empty
git push -u origin gh-pages`
```
After the branch has been created go to GitHub `settings -> pages -> source -> select branch -> gh-pages`.

If the GitHub pages are private, you need to use github token to add helm repo (generate read only token for this):
- add `helm repo add <repo-name> "https://<token>@raw.githubusercontent.com/<owner>/<repo>/gh-pages"`

Branch (`gh-pages`) and remote (`origin`) can be different. They can be set by `-pages-branch` and `-remote` flags when
running hcr.

## Download
- [binary](https://github.com/pete911/hcr/releases)

## Install

### Brew
- add tap `brew tap pete911/tap`
- install `brew install hcr`

### Go
- clone `git clone https://github.com/pete911/hcr.git && cd hcr`
- install `go install`

## Use

### Local
```
Usage of hcr:
  -charts-dir string
        The Helm charts location, can be specific chart (default "charts")
  -dry-run
        Whether to skip release update gh-pages index update
  -helm-key string
        Name of the key to use when signing. Used if --sign is true
  -helm-keyring string
        Location of a public keyring
  -helm-passphrase-file string
        Location of a file which contains the passphrase for the signing key
  -helm-sign
        Use a PGP private key to sign this package
  -pages-branch string
        The GitHub pages branch (default "gh-pages")
  -pre-release
        Whether the (chart) release should be marked as pre-release
  -remote string
        The Git remote for the GitHub Pages branch (default "origin")
  -tag string
        Release tag, defaults to chart version
  -token string
        GitHub Auth Token
  -version
        Print hcr version
```

### GitHub action
This is an example of how hcr can be used as a GitHub action, it is safe to run it on every commit, only commits with
changes to `Chart.yaml` `version` field will trigger release.

```yaml
name: helm-release
on:
  push:
    branches:
      - main
jobs:
  chart:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: Download chart releaser
        run: |
          curl -sSLo hcr.tar.gz "https://github.com/pete911/hcr/releases/download/v0.0.4/hcr_0.0.4_linux_amd64.tar.gz"
          tar -xzf hcr.tar.gz
          rm -f hcr.tar.gz
      - name: Package and release chart
        run: |
          git config user.email "gh-action@users.noreply.github.com"
          git config user.name "gh-action"
          ./hcr -token "${{ secrets.GITHUB_TOKEN }}"
```

Simply run `hcr` inside the project, this will:
- check if `-pages-branch` exists
- packages helm charts from `-charts-dir` to current directory as `<name>-<version>.tgz`
- adds `-pages-branch` git worktree to temp. directory
- creates release per chart in `-charts-dir` and uploads packaged chart as asset to that release
- update index file with the released charts, commit and push index file to `-pages-branch`

This makes it simpler and easier than [helm chart releaser](https://github.com/helm/chart-releaser), because we are not
reading from GitHub pages (or downloading releases) over http, so we don't face issues with restrictions on private
repositories (autogenerated pages link, authentication, ...), caching (GitHub pages are not updated immediately with
the latest changes) etc.

## Release
Releases are published when the new tag is created e.g. git tag -m "add super cool feature" v0.0.1 && git push --follow-tags
