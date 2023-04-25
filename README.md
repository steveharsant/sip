# SIP: Simple IP

A super simple Go app to quickly get IPv4 addresses for different interfaces.

## Why

* `ifconfig` is old and too verbose
* `Get-NetIPConfiguration` is slow and even more verbose
* `ifconfig` is deprecated
* `ip` prints text that is too close together for a quick glance.

It's also the first thing I have ever written in Go and wanted a more interesting and useful project than another 'Hello World' 😉

## Available Flags

| Flag   | Behaviour                                                                              |
|--------|----------------------------------------------------------------------------------------|
| `None` | Prints the hostname and queries the local resolver for the hosts registered IP address |
| `-a`   | Prints hostname, external IP address and addresses for all ainterfaces                 |
| `-e`   | Prints only the external IP address by querying api.ipify.org                          |
| `-f`   | To be used with `-i`. Regex pattern filtering for all interfaces                       |
| `-h`   | Prints only the hostname                                                               |
| `-i`   | Prints all up interfaces                                                               |
| `-v`   | Prints the version                                                                     |

## Roadmap

* Add further info like: dns, default route, IPv6, CIDR, etc

> [Inspired by one of the first things I ever wrote in bash, initially about 10 years ago](https://github.com/steveharsant/bash_scripts/blob/master/sip.sh)
