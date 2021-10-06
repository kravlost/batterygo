package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/distatus/battery"
	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
)

const (
	warningPercentage = 40
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon(100, warningPercentage, false))
	systray.SetTitle("BatteryGo")

	menuHdr := systray.AddMenuItem("BatteryGo", "The BatteryGo app")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Exit", "Exits BatteryGo")

	go func() {
		var lastPci int = 0

		for {
			pc, charging, err := batteryPercentage()
			pci := int(pc)

			if err != nil {
				systray.SetTooltip("Battery: Unknown")
				systray.SetIcon(getIcon(0, warningPercentage, false))
			} else {
				systray.SetTooltip(fmt.Sprintf("Battery: %d%%", pci))
				systray.SetIcon(getIcon(pci, warningPercentage, charging))
			}

			if pci < lastPci {
				switch pci {
				case warningPercentage:
					dialog.Message("Battery level: %d%%", pci).Title("BatteryGo").Info()
				case warningPercentage - 5:
					dialog.Message("Battery level: %d%%", pci).Title("BatteryGo").Info()
				case warningPercentage - 10:
					dialog.Message("Battery level: %d%%", pci).Title("BatteryGo").Error()
				}
			}

			lastPci = pci

			time.Sleep(60 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-menuHdr.ClickedCh:

			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
}

func batteryPercentage() (float64, bool, error) {
	batt, err := battery.Get(0)

	if err != nil {
		return 0, false, fmt.Errorf("could not get battery info: %w", err)
	}

	return 100 * batt.Current / batt.Full, batt.State == battery.Charging, nil
}

//go:embed assets/battery5.ico
var icon5 []byte

//go:embed assets/battery4.ico
var icon4 []byte

//go:embed assets/battery3.ico
var icon3 []byte

//go:embed assets/battery2.ico
var icon2 []byte

//go:embed assets/battery1.ico
var icon1 []byte

//go:embed assets/charging.ico
var iconCharging []byte

func getIcon(pc, warningPc int, charging bool) []byte {
	var bytes []byte

	if charging {
		return iconCharging
	}

	switch {
	case pc <= warningPc:
		bytes = icon1
	case float32(pc-warningPc)/float32(100-warningPc) < 0.25:
		bytes = icon2
	case float32(pc-warningPc)/float32(100-warningPc) < 0.50:
		bytes = icon3
	case float32(pc-warningPc)/float32(100-warningPc) < 0.75:
		bytes = icon4
	default:
		bytes = icon5
	}

	return bytes
}
