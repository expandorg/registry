# Contributing to expandorg/registry

:+1: First off, thanks for taking the time to contribute! :+1:

We're excited to hear and learn from you. Your experiences will benefit others who read and use these guides.
We've put together the following guidelines to help you figure out where you can best be helpful. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.


## Table of Contents

0. [Types of contributions we're looking for](#types-of-contributions-were-looking-for)
0. [How Can I Contribute?](#how-can-i-contribute)
0. [Setting up your environment](#setting-up-your-environment)
0. [Pull requests](#pull-requests)
0. [License](#license)

## Types of contributions we're looking for

There are many ways you can directly contribute to the guides (in descending order of need):

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- suggesting enhancements

## How Can I Contribute?

### Report bugs using Github's [issues](https://github.com/expandorg/registry/issues)
We use GitHub issues to track public bugs. Report a bug by [opening a new issue](); it's that easy! Write bug reports with detail, background, and sample code.

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can. 
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

### Suggesting Enhancements
How Do I Submit A Enhancement Suggestion?

Enhancement suggestions are tracked as GitHub [issues](https://github.com/expandorg/registry/issues). Create an issue on this repository and provide the following information:

- Use a clear and descriptive title for the issue to identify the suggestion.
- Provide a step-by-step description of the suggested enhancement in as many details as possible.
- Describe the current behavior and explain which behavior you expected to see instead** and why.

## Setting up your environment

Clone the repository with: 

`go get -u github.com/expandorg/registry`

or create a directory `$GOPATH/src/github/expandorg` and execute: git clone git@github.com:expandorg/registry.git 

Run the project dependencies (db, etc.) with `make up`
Run the latest migration with `make migrate-latest`
Run the project with `make run`

### Dependencies

We use `dep` to manage our dependencies. To add a new vendor, use: 

`deps ensure -add DEPENDENCY`

To update vendors for built project, run:

`make update-deps`

#### Add a new migration

```make add-migration name="migration_name"```

For migration names be descriptive and start with verbs: `create_`, `drop_`, `add_`, etc. This will look at the latest migrated version (1, 2, 3) and creates 2 files with new version:

`2_migration_name.up.sql` and `2_migration_name.down.sql`

#### Migrate

You can migrate to latest:

```make migrate-latest```

OR You can migrate up and migrate down a version:

```make run-migrations action="goto" version="1"```

When you migrate up, you can see in the `schema_migrations` the last migrated version. When you migrate down, it updates the the version column in `schema_migrations`.

### Tests
```make run-tests```

## Pull Requests
We Use [Github Flow](https://guides.github.com/introduction/flow/index.html), So All Code Changes Happen Through Pull Requests
Pull requests are the best way to propose changes to the codebase (we use [Github Flow](https://guides.github.com/introduction/flow/index.html)). We actively welcome your pull requests:

1. Fork the repo and create your branch from `master`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!



## License
By contributing, you agree that your contributions will be licensed under its MLP-2 License.

Thanks! :heart:

Expand.org Team
