<!-- omit in toc -->
# Contributing to simple-proxy

First off, thanks for taking the time to contribute! ❤️

All types of contributions are encouraged and valued. See the [Table of Contents](#table-of-contents) for different ways
to help and details about how this project handles them. Please make sure to read the relevant section before making
your contribution. It will make it a lot easier for us maintainers and smooth out the experience for all involved.
The community looks forward to your contributions. 🎉

> And if you like the project, but just don't have time to contribute, that's fine. There are other easy ways to support
> the project and show your appreciation, which we would also be very happy about:
> - Star the project
> - Tweet about it
> - Refer this project in your project's readme
> - Mention the project at local meetups and tell your friends/colleagues

<!-- omit in toc -->
## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [I Have a Question](#i-have-a-question)
- [I Want To Contribute](#i-want-to-contribute)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Enhancements](#suggesting-enhancements)

## Code of Conduct

This project and everyone participating in it is governed by the
[simple-proxy Code of Conduct](https://github.com/jthomperoo/simple-proxy/blob/main/CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code. Please report unacceptable behavior
to j.thomperoo@hotmail.com.

## I Have a Question

> If you want to ask a question, we assume that you have read the available
> [Documentation](https://github.com/jthomperoo/simple-proxy/blob/main/README.md).

Before you ask a question, it is best to search for existing
[Issues](https://github.com/jthomperoo/simple-proxy/issues) that might help you. In case you have found a
suitable issue and still need clarification, you can write your question in this issue. It is also advisable to
search the internet for answers first.

If you then still feel the need to ask a question and need clarification, we recommend the following:

- Open an [Issue](https://github.com/jthomperoo/simple-proxy/issues/new).
- Provide as much context as you can about what you're running into.
- Provide project and platform versions, depending on what seems relevant.

We will then take care of the issue as soon as possible.

## I Want To Contribute

> ### Legal Notice <!-- omit in toc -->
> When contributing to this project, you must agree that you have authored 100% of the content, that you have the
> necessary rights to the content and that the content you contribute may be provided under the project license.

### Reporting Bugs

<!-- omit in toc -->
#### Before Submitting a Bug Report

A good bug report shouldn't leave others needing to chase you up for more information. Therefore, we ask you to
investigate carefully, collect information and describe the issue in detail in your report. Please complete the
following steps in advance to help us fix any potential bug as fast as possible.

- Determine if your bug is really a bug and not an error on your side e.g. using incompatible environment
components/versions (Make sure that you have read the
[documentation](https://github.com/jthomperoo/simple-proxy/blob/main/README.md). If you are looking for support, you
might want to check [this section](#i-have-a-question)).
- To see if other users have experienced (and potentially already solved) the same issue you are having, check if there
is not already a bug report existing for your bug or error in the
[bug tracker](https://github.com/jthomperoo/simple-proxy/issues).
- Also make sure to search the internet (including Stack Overflow) to see if users outside of the GitHub community have
discussed the issue.
- Collect information about the bug:
- Kubernetes version.
- Any libraries/tooling that you are using that may affect it.
- Can you reliably reproduce the issue? And can you also reproduce it with older versions?

<!-- omit in toc -->
#### How Do I Submit a Good Bug Report?

We use GitHub issues to track bugs and errors. If you run into an issue with the project:

- Open an [Issue](https://github.com/jthomperoo/simple-proxy/issues/new). (Since we can't be sure at this point
whether it is a bug or not, we ask you not to talk about a bug yet and not to label the issue.)
- Explain the behavior you would expect and the actual behavior.
- Please provide as much context as possible and describe the *reproduction steps* that someone else can follow to
recreate the issue on their own. This usually includes your code. For good bug reports you should isolate the problem
and create a reduced test case.
- Provide the information you collected in the previous section.

Once it's filed:

- Someone will try to reproduce the issue with your provided steps.
- If someone is able to reproduce the issue the issue will be left to be [implemented by
someone](#your-first-code-contribution).

### Suggesting Enhancements

This section guides you through submitting an enhancement suggestion for simple-proxy,
**including completely new features and minor improvements to existing functionality**. Following these guidelines will
help maintainers and the community to understand your suggestion and find related suggestions.

<!-- omit in toc -->
#### Before Submitting an Enhancement

- Make sure that you are using the latest version.
- Read the [documentation](https://github.com/jthomperoo/simple-proxy/blob/main/README.md) carefully and find out if
the functionality is already covered, maybe by an individual configuration.
- Perform a [search](https://github.com/jthomperoo/simple-proxy/issues) to see if the enhancement has already
been suggested. If it has, add a comment to the existing issue instead of opening a new one.
- Find out whether your idea fits with the scope and aims of the project. It's up to you to make a strong case to
convince the project's developers of the merits of this feature. Keep in mind that we want features that will be useful
to the majority of our users and not just a small subset. If you're just targeting a minority of users, consider
writing an add-on/plugin library.

<!-- omit in toc -->
#### How Do I Submit a Good Enhancement Suggestion?

Enhancement suggestions are tracked as [GitHub issues](https://github.com/jthomperoo/simple-proxy/issues).

- Use a **clear and descriptive title** for the issue to identify the suggestion.
- Provide a **step-by-step description of the suggested enhancement** in as many details as possible.
- **Describe the current behavior** and **explain which behavior you expected to see instead** and why. At this point
you can also tell which alternatives do not work for you.
- **Explain why this enhancement would be useful** to most simple-proxy users. You may also want to point out
the other projects that solved it better and which could serve as inspiration.

## Developing

### Environment

Developing this project requires these dependencies:

* [Go](https://golang.org/doc/install) >= `1.24`
* [Static Check](https://github.com/dominikh/go-tools) == `2023.1 (v0.4.0)`

### Commands

* `make` - builds the binaries.
* `make test` - runs the unit tests.
* `make lint` - lints the code.
* `make format` - formats the code, must be run to pass the CI.
* `make zip` - zips the code for distribution along with the license.

## Styleguides

### Commit messages

Commit messages should follow the ['How to Write a Git Commit Message'](https://chris.beams.io/posts/git-commit/) guide.

### Documentation

Documentation should be in plain english, with 120 character max line width.

### Code

Project code should pass the linter and all tests should pass.

## Attribution

This guide is based on the **contributing-gen**. [Make your own](https://github.com/bttger/contributing-gen)!
