package main

import "fmt"

func main() {
	fmt.Println("Map is pass by reference in Go")
	m := make(map[string]int)
	m["abc"] = 1
	m["ijk"] = 2
	m["xyz"] = 3
	fun1(m)
	fmt.Println("main", m)
}

func fun1(m map[string]int) {
	fmt.Println("Inside fun1")
	fmt.Println("fun1 before", m)
	fun2(m)
	fmt.Println("fun1 after", m)
}

func fun2(m map[string]int) {
	m["abc"] = 100
	eventList := map[string]interface{}{
		"_lm.resourceId": []interface{}{
			map[string]interface{}{
				"system.deviceId":   "1",
				"system.deviceName": "2",
			},
		},
	}
	eventList2 := map[string][]map[string]any{
		"_lm.resourceId": {
			{
				"system.deviceId":   "1",
				"system.deviceName": "2",
			},
		},
	}
	_, _ = eventList, eventList2
}
