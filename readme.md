# xogs

[![test](https://github.com/KarnerTh/xogs/actions/workflows/test.yaml/badge.svg)](https://github.com/KarnerTh/xogs/actions/workflows/test.yaml)

xogs is a terminal-based tool designed to enhance the way you work with
local logs. It streamlines log aggregation, inspection and filtering,
enabling you to handle logs more effectively and focus on what matters.

Whether you're debugging, monitoring, or just exploring log data,
xogs provides useful tools without leaving your terminal.

https://github.com/user-attachments/assets/16075b08-aae4-4d1c-9c5c-a44f3d11c4ea


Want to see what has changed? Take a look at
the [Changelog](https://github.com/KarnerTh/xogs/blob/main/changelog.md)

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

- Supports different formats through generic parsers (JSON, regex, logfmt)
- Filter logs quickly - pinpoint relevant information
- Focus on what is important - granular control of displayed data
- Stream logs from stdin or analyse a file

## How does it work

1. Setup your configuration by providing a display and pipeline config (see [Configuration](#configuration))
2. Start aggregating your logs by either providing a file or piping data (see [Example usages](#example-usages))
3. Enjoy your aggregated logs locally and start filtering and exploring ✨

## Parameters/Flags

Some configurations can be set via command line parameters/flags.
The available options can also be viewed via `xogs -h`

```
  -h, --help             help for xogs
      --profile string   the profile from the config that should be used
```

## Configuration

Xogs uses a yaml configuration file called `.xogs` that can be located
in your home and/or working directory. The configuration file in the working
directory takes precedence and can e.g. be tracked by a version control system
and be project specific.

Example configuration file can be found [here](https://github.com/KarnerTh/xogs/blob/main/.xogs)

### JSON parser

Basic configuration example assuming your logs are in JSON format like

```json
{ "level": "debug", "msg": "test" }
```

```yaml
defaultProfile: json_example # default fallback if no --profile flag is used (optional)
profiles:
  - name: json_example # custom name to be referenced by
    pipeline:
      processors:
        - parser:
            json: {} # JSON parser does not need any parameters
    displayConfig:
      columns:
        - title: level
          width: 0.1 # the fraction of width the column should use (= 10%)
          valueKey: level # the value key from the JSON log
        - title: msg
          width: 0.9
          valueKey: msg
        - title: nested
          width: 0.1
          valueKey: some.nested.data # also supports nested JSON keys
```

### Logfmt parser

Basic configuration example assuming your logs are in logfmt format like

```logfmt
level=info tag=0815 env=prod msg="test log"
```

```yaml
profiles:
  - name: logfmt_example # custom name to be referenced by
    parser:
      logfmt: {} # logfmt parser does not need any parameters
    displayConfig:
      columns:
        - title: level
          width: 0.1 # the fraction of width the column should use (= 10%)
          valueKey: level # the value key from the logfmt log
        - title: msg
          width: 0.8
          valueKey: msg
        - title: tag
          width: 0.1
          valueKey: tag
```

### Regex parser

Regex example with the `ping` output in the format

```
64 bytes from 8.8.8.8: icmp_seq=7 ttl=118 time=20.3 ms
```

```yaml
profiles:
  - name: ping_example # custom name to be referenced by
    pipeline:
      processors:
        - parser:
            regex:
              values: # regex parser specifies a list of parsed values
                - key: time # data key to reference value
                  regex: "time=(.*)" # regex to extract value
                - key: ttl
                  regex: 'ttl=(\d*)'
    displayConfig:
      columns:
        - title: time
          width: 0.1 # the fraction of width the column should use (= 10%)
          valueKey: time # the value key from the regex mapping
        - title: ttl
          width: 0.1
          valueKey: ttl
```

### Remapper

If you want to remap a value to a different key you can use a remapper.
In the following example, the JSON log uses the JSON key `msg`, but we
want to remap it to `message`:

```yaml
profiles:
  - name: json_example_with_remap
    pipeline:
      processors:
        - parser:
            json: {}
        - inputKey: msg # the source key that should be remapped
          remapper:
            targetKey: message # the target key where it should be remapped
    displayConfig:
      columns:
        - title: level
          width: 0.1
          valueKey: level
        - title: msg
          width: 0.9
          valueKey: message # we can now use the remapped key here
```

### Combine multiple parser

Sometimes log sources do not follow the same format for every log line
(even though they should) or have a combination of different formats,
so xogs allows to provide multiple processors.

Assume we have the following log format, which has a prefix followed by a
JSON log

```
[service-a] {"level": "debug", "msg": "test"}
```

```yaml
profiles:
  - name: json_with_prefix
    pipeline:
      processors:
        - parser:
            regex: # first use a regex parser to extract service name and log
              values:
                - key: service
                  regex: '\[([^\]]+)\]'
                - key: log
                  regex: '\[[^\]]+\]\s(.*)'
        - inputKey: log # use the extracted log data and parse it as JSON
          parser:
            json: {}
    displayConfig:
      columns:
        - title: service
          width: 0.1
          valueKey: "service"
        - title: level
          width: 0.1
          valueKey: level
        - title: message
          width: 0.8
          valueKey: msg
```

### Reserved value keys

There are some internal value keys that should only be used as sources, e.g.
in the display config or as input for a parser, but should never be mutated:

- `__id`: the internal id of the log
- `__raw`: the unparsed raw log line

## Key bindings

| Action                        | Key Binding       |
| ----------------------------- | ----------------- |
| Go up                         | `arrow up`, `k`   |
| Go down                       | `arrow down`, `j` |
| Go to the top                 | `g`               |
| Go to the bottom (and follow) | `G`               |
| Insert mode for filter        | `i`               |
| Exit insert                   | `esc`             |
| Open details                  | `enter`           |
| Exit details                  | `esc`, `q`        |
| Clear filter                  | `ctrl+x`          |
| Exit xogs                     | `ctrl+c`          |

## Filtering

There are two filtering methods supported currently

- string compare - search for a string somewhere in the log
- aggregated data search - limit the search to a data field
  - uses format `key:searchValue`, e.g. `level:error`

## Example usages

```bash
# pipe ping output to xogs
ping 8.8.8.8 | xogs --profile=ping

# analyze log file
xogs example.log --profile=custom
```
