# BatteryGo

System tray app in Go to show warnings when the battery level drops below a specified level, currently set to 40%. The battery state is read once a minute.

This was written because my laptop started shutting down when the battery got to around 30% and I wanted some warning.

The system tray and warning dialog libraries are cross-platform, but I have only tested this on Windows 10.

Windows resource files (.syso) included to add an icon to the executable on Windows.

Program outline from https://dev.to/osuka42/building-a-simple-system-tray-app-with-go-899

Battery icons freeware from https://iconarchive.com/show/battery-icons-by-graphicloads.html

CGo not required.

Source repository: https://github.com/kravlost/batterygo