package bt

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Use system FZF installation
// From @jonaz GitHub
func fzf(data io.Reader) (string, error) {
	var result strings.Builder
	cmd := exec.Command("fzf")
	cmd.Stdout = &result
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(stdin, data)
	//_, err = data.WriteTo(stdin)
	if err != nil {
		return "", err
	}
	err = stdin.Close()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(result.String()), nil
}

// Use fzf to get MAC
func mac_from_fzf(names string, devices map[string]string) (string) {
	choice, err := fzf(strings.NewReader(names))
	if err != nil {
	    log.Fatal(err)
	}
	mac := devices[string(choice)]
	return mac
}

// Device map searchable fzf
func names_to_str(devices map[string]string) (string) {
	names_arr := make([]string, len(devices))
	idx := 0
	for name := range devices {
		names_arr[idx] = name
		idx++
	}
	names_str := strings.Join(names_arr, "\n")
	return names_str
}

func out_to_map(out string) (map[string]string) {
	scanner := bufio.NewScanner(strings.NewReader(out)) // Read lines
	devices := make(map[string]string)
	for scanner.Scan() {
		row := scanner.Text()

		// Process text
		cleaned := strings.SplitAfterN(string(row), " ", 2)[1]  // remove "Device"
		parts := strings.SplitAfterN(cleaned, " ", 2)  // 0: mac, 1: name
		mac := parts[0]
		mac = mac[:len(mac)-1]
		name := parts [1]
		devices[name] = mac
		}
	return devices
}

// Getting device MAC.
// Maybe should use DBUS to get BlueZ cache
func get_devices() string {
	out, err := exec.Command("bluetoothctl", "devices").Output()
	if err != nil {
	    log.Fatal(err)
	}
	return string(out)
}

// Get MAC from fzf-filtered list of known bluetoothctl devices
func get_mac() (string) {
	out := get_devices()
	devices := out_to_map(out)
	names_str := names_to_str(devices)
	mac := mac_from_fzf(names_str, devices)
	return mac
}

// Connect to device
func Connect() {
	fmt.Printf("Select device to connect to:")
	fmt.Printf("(if not in list, use `scan on` and `pair` in `bluetoothctl`)")
	mac := get_mac()
	cmd := exec.Command("bluetoothctl", "connect", mac)
	err := cmd.Run()
	if err != nil {
	    log.Fatal(err)
	}
}

// Disconnect from device
func Disconnect() {
	cmd := exec.Command("bluetoothctl", "disconnect")
	err := cmd.Run()

	if err != nil {
	    log.Fatal(err)
	}
}
