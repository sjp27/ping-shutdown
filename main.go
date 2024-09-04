// Copyright 2024 sjp27 <https://github.com/sjp27>. All rights reserved.
// Use of this source code is governed by the MIT license that can be
// found in the LICENSE file.

// Service to ping an IP address and shutdown server if no response within given timeout e.g. power cut.

package main

import (
	"fmt"
	probing "github.com/prometheus-community/pro-bing"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const version = "v1.0"

func main() {
	if len(os.Args) < 3 {
		fmt.Println(version + " Usage: ping-shutdown <ip> <timeout(mins)>")
	} else if len(os.Args) == 3 {
		timeout, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid timeout")
			panic(err)
		}
		time.Sleep(time.Minute * 2)
		pingIp(os.Args[1], timeout)
	} else {
		fmt.Println("Too many arguments")
	}
}

// pingIp ping ip address to check power is up
func pingIp(ip string, timeout int) {

	var failCount int = 0

	fmt.Println("Start pinging IP address:" + ip + " timeout:" + strconv.FormatInt(int64(timeout), 10) + "mins")

	for range time.Tick(time.Minute * 1) {

		pinger, err := probing.NewPinger(ip)

		if err != nil {
			panic(err)
		}

		pinger.Count = 1
		pinger.Timeout = time.Second * 2

		err = pinger.Run()

		if err != nil {
			panic(err)
		}

		stats := pinger.Statistics()

		if stats.PacketsRecv == 0 {
			failCount++
			fmt.Println("Ping failure:" + strconv.FormatInt(int64(failCount), 10))
			if failCount >= timeout {
				fmt.Println("Ping failure shutdown")
				cmd := exec.Command("shutdown", "-h", "now")
				_, err := cmd.Output()

				if err != nil {
					fmt.Println("Unable to shutdown:" + err.Error())
					panic(err)
				}
			}
		} else {
			failCount = 0
		}
	}
}
