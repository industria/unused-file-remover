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
