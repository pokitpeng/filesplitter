# filesplitter
split and merge file  tool

## features
- [x] split file
- [x] merge file

## Build
```bash
go build -o fs main.go
```

## Usage
```bash
# split big file, /home/src.log --> /home/tiny.log*
./fs split -s 50M /home/src.log /home/tiny.log 

# merge files, /home/tiny.log* --> /home/dst.log
./fs merge -p /home/tiny.log /home/dst.log

# check file
md5sum /home/src.log
md5sum /home/dst.log
```