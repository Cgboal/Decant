# Decant 
Decant can be used to expand CIDR ranges into a list of IP addresses easily. For example:
```
$ echo 192.168.0.0/24 | decant
```

You can also supply a list of ranges by cating a file into decant. 

## Installation 
```
go get github.com/cgboal/decant
```
