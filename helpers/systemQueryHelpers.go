// Package helpers ..
package helpers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Harikrishnan-Ashok/samBaar/state"
	"github.com/Harikrishnan-Ashok/samBaar/theme"
)

// GetBatteryStatus : to query the battery status
func GetBatteryStatus(store *state.UIState) {
	out, err := exec.Command("acpi", "-b").Output()
	if err != nil {
		store.BatteryValue = "Err"
	} else {
		store.BgColor = theme.DarkTheme.Colors.WarningColor
		part := strings.Split(string(out), ",")
		store.BatteryValue = strings.TrimSpace(part[1])
		store.RemainingTime = strings.Fields(strings.TrimSpace(part[2]))[0]
		split := strings.Split(part[0], ":")
		if len(split) >= 2 {
			store.AdapterStatus = strings.TrimSpace(split[1])
		}

	}
}

// GetTimeStatus : to query the Time
func GetTimeStatus(store *state.UIState) {
	out, err := exec.Command("date", "+%H:%M").Output()
	if err != nil {
		store.TimeStatus = "Err querying time info"
	} else {
		store.TimeStatus = strings.TrimSpace(string(out))
	}

}

// GetBluetoothStatus : to query BLuetooth Status
func GetBluetoothStatus(store *state.UIState) {
	out, err := exec.Command("bluetoothctl", "show").Output()
	if err != nil {
		store.BluetoothStatus.Value = "Err"
	} else {
		lines := strings.Split(string(out), "\n")
		status := "Unknown"
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "Powered:") {
				parts := strings.Fields(line)
				if len(parts) == 2 {
					if strings.ToLower(parts[1]) == "yes" {
						status = "Blue"
						store.BluetoothStatus.BgColor = theme.DarkTheme.Colors.EnabledColor
					} else if strings.ToLower(parts[1]) == "no" {
						status = "Blue"
						store.BluetoothStatus.BgColor = theme.DarkTheme.Colors.DisabledColor
					}
				}
				break
			}
		}
		store.BluetoothStatus.Value = status
	}
	out, err = exec.Command("date", "+ %d - %b - %Y  %A").Output()
	if err != nil {
		store.DateStatus = "Err getting date"
	} else {
		store.DateStatus = string(out)
	}
}

// GetNetworkStatus : to query the Time
func GetNetworkStatus(store *state.UIState) {
	out, err := exec.Command("nmcli", "-t", "-f", "TYPE,STATE", "device").Output()
	if err != nil {
		fmt.Println("Error running nmcli:", err)
		return
	}
	for _, line := range strings.Split(string(out), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}
		ifaceType := parts[0]
		ifaceState := parts[1]
		switch ifaceType {
		case "wifi":
			store.WifiStatus.Value = "WiFi"
			if ifaceState != "conected" {
				store.WifiStatus.BgColor = theme.DarkTheme.Colors.EnabledColor
			} else {
				store.WifiStatus.BgColor = theme.DarkTheme.Colors.DisabledColor
			}
		case "ethernet":
			store.WifiStatus.Value = "WiFi"
			if ifaceState != "conected" {
				store.EthernetStatus.BgColor = theme.DarkTheme.Colors.EnabledColor
			} else {
				store.EthernetStatus.BgColor = theme.DarkTheme.Colors.DisabledColor
			}

		}
	}
}
