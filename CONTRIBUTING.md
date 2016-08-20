
# Contributing to wuxia

We welcome contributions to wuxia of any kind including documentation, themes,
organization, tutorials, blog posts, bug reports, issues, feature requests,
feature implementations, pull requests, answering questions on the forum,
helping to manage issues, etc.


## Table of Contents

* [Asking Support Questions](#asking-support-questions)
* [Reporting Issues](#reporting-issues)
* [Submitting Patches](#submitting-patches)
  * [Code Contribution Guidelines](#code-contribution-guidelines)
  * [Git Commit Message Guidelines](#git-commit-message-guidelines)
  * [Using Git Remotes](#using-git-remotes)
  * [Build wuxia with Your Changes](#build-wuxia-with-your-changes)
  * [Add Compile Information to wuxia](#add-compile-information-to-wuxia)

## Asking Support Questions

We have an active [discussion forum]() where users and developers can ask questions.
Please don't use the Github issue tracker to ask questions.

## Reporting Issues

If you believe you have found a defect in wuxia or its documentation, use
the Github [issue tracker](https://github.com/gernest/wuxia/issues) to report the problem to the wuxia maintainers.
If you're not sure if it's a bug or not, start by asking in the [discussion forum](http://discuss.gowuxia.io).
When reporting the issue, please provide the version of wuxia in use (`wuxia version`) and your operating system.

## Submitting Patches

The wuxia project welcomes all contributors and contributions regardless of skill or experience level.
If you are interested in helping with the project, we will help you with your contribution.
wuxia is a very active project with many contributions happening daily.
Because we want to create the best possible product for our users and the best contribution experience for our developers,
we have a set of guidelines which ensure that all contributions are acceptable.
The guidelines are not intended as a filter or barrier to participation.
If you are unfamiliar with the contribution process, the wuxia team will help you and teach you how to bring your contribution in accordance with the guidelines.

### Code Contribution Guidelines

To make the contribution process as seamless as possible, we ask for the following:

* Go ahead and fork the project and make your changes.  We encourage pull requests to allow for review and discussion of code changes.
* When you’re ready to create a pull request, be sure to:
    * Sign the [CLA](CLA.md).
    * Have test cases for the new code. If you have questions about how to do this, please ask in your pull request.
    * Run `make check`.
    * Add documentation if you are adding new features or changing functionality.  The docs site lives in `/docs`.
    * Squash your commits into a single commit. `git rebase -i`. It’s okay to force update your pull request with `git push -f`.
    * Make sure `make test.` passes, and `make build` completes. [Travis CI](https://travis-ci.org/gernets/wuxia) (Linux and OS&nbsp;X) 
    * Follow the **Git Commit Message Guidelines** below.

### Git Commit Message Guidelines

This [blog article](http://chris.beams.io/posts/git-commit/) is a good resource for learning how to write good commit messages,
the most important part being that each commit message should have a title/subject in imperative mood starting with a capital letter and no trailing period:
*"Return error on wrong use of the Paginator"*, **NOT** *"returning some error."*
Also, if your commit references one or more GitHub issues, always end your commit message body with *See #1234* or *Fixes #1234*.
Replace *1234* with the GitHub issue ID. The last example will close the issue when the commit is merged into *master*.
Sometimes it makes sense to prefix the commit message with the packagename (or docs folder) all lowercased ending with a colon.
That is fine, but the rest of the rules above apply.
So it is "tpl: Add emojify template func", not "tpl: add emojify template func.", and "docs: Document emoji", not "doc: document emoji."

An example:

```text
TODO
```

### Using Git Remotes

Due to the way Go handles package imports, the best approach for working on a
wuxia fork is to use Git Remotes.  Here's a simple walk-through for getting
started:

1. Get the latest wuxia sources:

    ```
    go get -u -t github.com/gernest/wuxia/...
    ```

1. Change to the wuxia source directory:

    ```
    cd $GOPATH/src/github.com/gernest/wuxia
    ```

1. Create a new branch for your changes (the branch name is arbitrary):

    ```
    git checkout -b iss1234
    ```

1. After making your changes, commit them to your new branch:

    ```
    git commit -a -v
    ```

1. Fork wuxia in Github.

1. Add your fork as a new remote (the remote name, "fork" in this example, is arbitrary):

    ```
    git remote add fork git://github.com/USERNAME/wuxia.git
    ```

1. Push the changes to your new remote:

    ```
    git push --set-upstream fork iss1234
    ```

1. You're now ready to submit a PR based upon the new branch in your forked repository.

### Build wuxia with Your Changes

```bash
cd $GOPATH/src/github.com/gernest/wuxia
make build
make install
```

### Add Compile Information to wuxia

To add compile information to wuxia, replace the `go build` command with the following 

```bash
make build-git-release
```

This will result in `wuxia version` output that looks similar to:

    wuxia Static Site Generator v0.13-DEV-8042E77 buildDate: 2014-12-25T03:25:57-07:00

Alternatively, just run `make` &mdash; all the “magic” above is already in the `Makefile`.  :wink:

