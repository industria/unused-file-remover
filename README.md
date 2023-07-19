# unused-file-remover

Very simple tool to remove files that have not been accesses (atime) for x amount of time.

It will simply ignore directories so you will possibly have empty directories left after running this tool.


## usage

```sh
Usage: unused-file-remover arguments
  -cachePath string
        Path where binaries are cached on disk
  -maxage string
        Maximum age, as Go duration format, file access time must have before it is deleted (default: 168h 7 days) (default "168h")
  -verbose
        Verbose output
```

## using the docker image

There is a docker image containing the unused-file-remover program which can be used as follows:

```sh
docker run jlindstorff/unused-file-remover:latest -cachePath=/tmp/ -maxage=168h -verbose
```
