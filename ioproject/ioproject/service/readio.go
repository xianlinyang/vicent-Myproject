package service

import (
	"github.com/robfig/cron/v3"
	"ioproject/config"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func ReadIOstat(cf config.Config, ch chan int, m chan map[string]float64) error {
	iomap := make(map[string]float64)
	var count int

	ticker := cron.New()
	_, err := ticker.AddFunc("@every 1s", func() {
		count = count + 1
		cmd := exec.Command("iostat", "-d", "-x", cf.Disk)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
			return
		}

		singleMap, err := paramMap(string(out), cf.Disk)
		if err != nil {
			return
		}
		for param, paramValue := range singleMap {
			if value, ok := iomap[param]; ok {
				iomap[param] = value + paramValue
			} else {
				iomap[param] = paramValue
			}
		}

	})
	if err != nil {
		return err
	}

	ticker.Start()

	select {
	case a := <-ch:
		if a == -1 {
			m1 := make(map[string]float64)
			for key, value := range iomap {
				m1[key] = value / float64(count)
			}
			ticker.Stop()
			m <- m1
			return err
		}
	}
	return nil
}

func paramMap(param, diskName string) (map[string]float64, error) {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	m3 := make(map[string]float64)
	commandParamStr := GetBetweenStr(param, "Device", diskName)
	commandStrArr := strings.Fields(commandParamStr)
	for i := 0; i < len(commandStrArr); i++ {
		if "Device" != commandStrArr[i] {
			m1[commandStrArr[i]] = i
		}
	}

	commandStr := strings.Fields(param)

	var paramArr []string
	var isTrue bool
	for i := 0; i < len(commandStr); i++ {
		if isTrue {
			paramArr = append(paramArr, commandStr[i])
		}

		if commandStr[i] == diskName {
			isTrue = true
		}
	}

	for i := 0; i < len(paramArr); i++ {
		m2[paramArr[i]] = i + 1
	}

	for key, i := range m1 {
		for value, j := range m2 {
			if i == j && key != "Device" {
				paramValue, err := strconv.ParseFloat(value, 2)
				if err != nil {
					return nil, err
				}
				m3[key] = paramValue
			}
		}
	}
	return m3, nil
}

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
