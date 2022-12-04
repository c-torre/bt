# bt - a Bluetooth wrapper written in Go

`bt` makes it easy to connect and disconnect from Bluetooth devices.
It runs `bluetoothctl` commands to connect and disconnect from devices, using `fzf` to select from cached devices.

# Usage

1. Run `go build` to make the binary `bt`.
2. Move `bt` to a directory within `path`.
3. Run `bt --help` to see the options.

# Dependencies

- go
- fzf
