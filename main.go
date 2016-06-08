package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

type CoreInfo struct {
	coreId      int
	processorId int
}

func nproc() int {
	cmd := exec.Command("nproc")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error while executing nproc")
		panic(err)
	}

	out := strings.Trim(stdout.String(), "\n")
	value, err := strconv.Atoi(out)

	if err != nil {
		fmt.Printf("Error while conv '%s' to int", out)
		panic(err)
	}

	return value
}

func main() {
	content, err := ioutil.ReadFile("/proc/cpuinfo")

	if err != nil {
		panic(err)
	}

	cpuInfo := strings.Split(string(content), "\n")

	numberOfCores := nproc()

	cores := make([]CoreInfo, numberOfCores)
	indexCounter := 0

	for _, line := range cpuInfo {
		tokens := strings.SplitN(line, ":", 2)

		if len(tokens) != 2 {
			continue
		}

		trimCutPoints := "\t \n\r"
		key := strings.Trim(tokens[0], trimCutPoints)

		if key != "processor" && key != "core id" {
			continue
		}

		if len(cores) <= indexCounter {
			cores[indexCounter] = CoreInfo{
				coreId:      0,
				processorId: 0,
			}
		}

		value, _ := strconv.Atoi(strings.Trim(tokens[1], trimCutPoints))

		if key == "processor" {
			cores[indexCounter].processorId = value
		}

		if key == "core id" {
			cores[indexCounter].coreId = value
			indexCounter = indexCounter + 1
		}
	}

	for _, coreInfo := range cores {
		if coreInfo.processorId != coreInfo.coreId {
			fmt.Printf("echo 0 > /sys/devices/system/cpu/cpu%d/online\n", coreInfo.processorId)
		}
	}
}
