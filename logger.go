package tools

import (
	"fmt"
	"runtime"
)

const (
	LogClose = iota
	LogError
	LogWarning
	LogInfo
	LogAll
)

func Compare(preset int, present int, a ...interface{}) {
	if preset > present {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		switch preset {
		case LogError:
			fmt.Printf("Error:%v %v\n", file, line)
			fmt.Println(a...)
			break
		case LogWarning:
			fmt.Printf("Warning:%v %v\n", file, line)
			fmt.Println(a...)
			break
		case LogInfo:
			fmt.Printf("Info:%v %v\n", file, line)
			fmt.Println(a...)
			break
		}
	}
}

func Println(lv int, a ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		switch lv {
		case LogError:
			fmt.Printf("Error:%v %v\n", file, line)
			fmt.Println(a...)
			break
		case LogWarning:
			fmt.Printf("Warning:%v %v\n", file, line)
			fmt.Println(a...)
			break
		case LogInfo:
			fmt.Printf("Info:%v %v\n", file, line)
			fmt.Println(a...)
			break
		}
	}
}
