# xogs
Your local logs - just better

Want to see what has changed? Take a look at
the [Changelog](https://github.com/KarnerTh/xogs/blob/main/changelog.md)

# Contents
<ul>
  <li><a href="#installation">Installation</a></li>
  <li><a href="#features">Features</a></li>
  <li><a href="#how-does-it-work">How does it work</a></li>
  <li><a href="#parametersflags">Parameters/Flags</a></li>
  <li><a href="#configuration-file">Configuration file</a></li>
  <li><a href="#example-usages">Example usages</a></li>
</ul>

## Installation
Through the go toolchain (requires go version >=1.21)
```sh
go install github.com/KarnerTh/xogs@latest
```

or

just head over to the [Releases](https://github.com/KarnerTh/xogs/releases) 
page and download the right executable for your operating system. To be able 
to use it globally on your system, add the executable to your path.

## Features
* Supports differnt formats through generic parsers (json, regex, logfmt)
* Filter your logs
* Focus on what is important - granular control of dislayed data

## How does it work
TBD

## Parameters/Flags
Some configurations can be set via command line parameters/flags. 
The available options can also be viewed via `xogs -h`

```
  -h, --help             help for xogs
      --profile string   the profile from the config that should be used
```

## Configuration file

Example:

https://github.com/KarnerTh/xogs/blob/main/.xogs

## Example usages
```bash
# pipe ping output to xogs
ping 8.8.8.8 | xogs --profile=ping
```
