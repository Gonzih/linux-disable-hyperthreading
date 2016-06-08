# Disable hyperthreading on linux

## Installation

```bash
go get github.com/Gonzih/linux-disable-hyperthreading
```

## Usage

```bash
$ linux-disable-hyperthreading
echo 0 > /sys/devices/system/cpu/cpu4/online
echo 0 > /sys/devices/system/cpu/cpu5/online
echo 0 > /sys/devices/system/cpu/cpu6/online
echo 0 > /sys/devices/system/cpu/cpu7/online

$ linux-disable-hypearthreading | sudo bash
```
