## Media files sorter

This was made to fix the mess on my backups hard drives.

### Build
    go build -o hd-sorter *.go

The program as two modes, snapshot and sort.

### Snapshot
Will iterate recursively over the current directory and print the file types based on their extensions and their occurence.

    ./hd-sorter snapshot

### Sort
Will iterate recursively over the current directory and move the files based on their extensions in the right directory.

    ./hd-sorter sort