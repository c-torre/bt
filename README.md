# bt - a Bluetooth wrapper written in Go

`bt` makes it easy to connect and disconnect from Bluetooth devices.
It runs `bluetoothctl` CLI (`bluez5`) commands to connect and disconnect from devices, using `fzf` to select from cached devices.

# Usage

1. Run `go build` to make the binary `bt`.
2. Use `bt --help` to see the options.

(Optional) Move `bt` to a directory within `path`.

# Dependencies

- go
- fzf
- bluez-utils
