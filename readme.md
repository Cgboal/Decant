# Decant 

Decant can be used to expand CIDR ranges into a list of IP addresses easily. For example:
```
$ echo 192.168.0.0/24 | decant
```

You can also supply a list of ranges by piping a file into decant

` $ decant < 2022-07-28T09:34:32-ipv4-ranges.txt `

## Installation 
```
go install github.com/cgboal/decant@latest
```
