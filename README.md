# SIP: Simple IP

A super simple Go app to get all IPv4 addresses for interfaces in that are up.

## Why

* `ifconfig` is old and too verbose
* `Get-NetIPConfiguration` is slow and even more verbose
* `ifconfig` is deprecated
* `ip` prints text that is too close together for a quick glance.

It's also the first thing I have ever written in Go and wanted a more interesting and useful project than another 'Hello World' 😉

## Ideas for improvements

* Get external IP address
* Switches to do things like:
  * Only return IP addresses
  * Only return addresses for physical interfaces
  * Sort by alphabetical, network, etc
  * Add further info like: dns, default route, IPv6, CIDR, etc

> [Inspired by one of the first things I ever wrote in bash, initially about 10 years ago](https://github.com/steveharsant/bash_scripts/blob/master/sip.sh)
