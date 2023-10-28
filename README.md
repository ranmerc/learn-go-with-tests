# Learn go with tests

Following [Learn go with tests](https://github.com/quii/learn-go-with-tests)

Notes present in each individual README.mds may be verbatim copy of content from the above resource.

## Go setup

1. Run `go mod init` in sub project folder. This will create a go.mod file. Now your folder is a module.
2. Run `go work use path_to_folder` to add it to `go.work` file. Running this before converting your folder to module will not work.
