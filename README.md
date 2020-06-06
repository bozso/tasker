# tasker

## Introduction

A simple wrapper around the [task](https://github.com/go-task/task)
golang command line program, that searches for `Taskfile.yml` in parent
directories and executes the task command with the found file.

Since it is a relly simple program, it might be merged into my
[gotoolbox](https://github.com/bozso/gotoolbox) repository.

## TODO

- At the moment it will append the parent directory string (`..`)
to the current filename (e.g. `Taskfile.yml` at the start) until it
finds it. It will not stop even if it has reached the root directory.
- Not cross platform. Root directory name should be queried based on
operating system.
