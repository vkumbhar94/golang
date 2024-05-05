package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// exp := regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)
	// exp := regexp.MustCompile(`(?P<first>\d+)\.(?:\d+).(?P<third>\d+)`)
	// exp := regexp.MustCompile(`((?:\w+)+),?`)
	// exp := regexp.MustCompile(`(((?:\d+)+)\.?)+`)
	exp := regexp.MustCompile(`(?P<first>(?:\d+)+)(?:\.?(?P<num>(?:\d+)+))*`)
	match := exp.FindStringSubmatch("1234.5678.9.7")

	fmt.Println(match)
	result := make(map[string]string)
	fmt.Println(strings.Join(exp.SubexpNames(), ","))
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	fmt.Printf("by name: %s %s %s\n", result["first"], result["second"], result["num"])
	ParseLogUsingRegex()
}

func ParseLogUsingRegex() {
	fmt.Println("parsing santaba logline")
	logLine := `[2024-04-15 20:34:39.059 PDT] [MSG] [INFO] [HighPriorityMessageTaskManager-1:::] [HTTPDeliveryResultTask.processMessage:121] HTTP Delivery result task received, CONTEXT=company=di, alertId=5152942, alertInstanceId=LMD43826, integrationId=4, alertStatus=update`
	exp := regexp.MustCompile(`\[(?P<timestamp>[^\]]+)\] \[(?P<msgType>[^\]]+)\] \[(?P<logLevel>[^\]]+)\] \[(?P<thread>[^\]]+)\] \[(?P<method>[^\]]+)\] (?P<message>.*), CONTEXT=(?P<context>.*)`)
	// exp := regexp.MustCompile("\\[(?P<timestamp>[^\\]]+)\\] \\[(?P<msgType>[^\\]]+)\\] \\[(?P<logLevel>[^\\]]+)\\] \\[(?P<thread>[^\\]]+)\\] \\[(?P<method>[^\\]]+)\\] (?P<message>.*)")
	match := exp.FindStringSubmatch(logLine)
	fmt.Println(match)
	result := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	fmt.Println(result)

}
