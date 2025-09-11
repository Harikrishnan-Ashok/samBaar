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
		store.BatteryValue.Value = "Err"
	} else {
		store.BgColor = theme.DarkTheme.Colors.WarningColor
		part := strings.Split(string(out), ",")
		store.BatteryValue.Value = strings.TrimSpace(part[1])
		if strings.TrimSpace(part[1]) == "100%" {
			store.RemainingTime.Value = "Nil"
		} else {
			store.RemainingTime.Value = strings.Fields(strings.TrimSpace(part[2]))[0]
		}
		split := strings.Split(part[0], ":")
		if len(split) >= 2 {
			store.AdapterStatus.Value = strings.TrimSpace(split[1])
		}

	}
}

// GetTimeStatus : to query the Time
func GetTimeStatus(store *state.UIState) {
	out, err := exec.Command("date", "+%H:%M").Output()
	if err != nil {
		store.TimeStatus.Value = "Err querying time info"
	} else {
		store.TimeStatus.Value = strings.TrimSpace(string(out))
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
						status = "enabled"
						store.BluetoothStatus.BgColor = theme.DarkTheme.Colors.EnabledColor
					} else if strings.ToLower(parts[1]) == "no" {
						status = "disabled"
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
		store.DateStatus.Value = "Err getting date"
	} else {
		store.DateStatus.Value = string(out)
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
			if ifaceState != "connected" {
				store.WifiStatus.BgColor = theme.DarkTheme.Colors.DisabledColor
				store.WifiStatus.Value = "disabled"
			} else {
				store.WifiStatus.BgColor = theme.DarkTheme.Colors.EnabledColor
				store.WifiStatus.Value = "enabled"
			}
		case "ethernet":
			if ifaceState != "connected" {
				store.EthernetStatus.BgColor = theme.DarkTheme.Colors.DisabledColor
				store.EthernetStatus.Value = "enabled"
			} else {
				store.EthernetStatus.BgColor = theme.DarkTheme.Colors.EnabledColor
				store.EthernetStatus.Value = "disabled"
			}

		}
	}
}

func GetBrightnessStatus(store *state.UIState) {
	out, err := exec.Command("brightnessctl", "-m").Output()
	if err != nil {
		fmt.Println(err)
	} else {
		currBrightness := strings.Split(string(out), ",")[3]
		store.Brightness.Value = currBrightness
	}
}
func GetVolumeStatus(store *state.UIState) {
	out, err := exec.Command("pamixer", "--get-volume").Output()
	if err != nil {
		fmt.Println(err)
	} else {
		store.VolStatus.Value = string(out)
	}
}
