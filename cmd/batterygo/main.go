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

//go:embed assets/battery.ico
var iconBytes []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(iconBytes)
	systray.SetTitle("BatteryGo")

	menuHdr := systray.AddMenuItem("BatteryGo", "The BatteryGo app")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Exit", "Exits BatteryGo")

	go func() {
		var lastPci int = 0

		for {
			pc, err := batteryPercentage()
			pci := int(pc)

			if err != nil {
				systray.SetTooltip("Battery: Unknown")
			} else {
				systray.SetTooltip(fmt.Sprintf("Battery: %d%%", pci))
			}

			if pci < lastPci && pci == warningPercentage {
				dialog.Message("Battery level: %d%%", pci).Title("Low battery level!").Info()
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

func batteryPercentage() (float64, error) {
	battery, err := battery.Get(0)

	if err != nil {
		return 0, fmt.Errorf("could not get battery info: %w", err)
	}

	return 100 * battery.Current / battery.Full, nil
}
