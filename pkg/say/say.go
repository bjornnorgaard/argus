package say

import (
	"fmt"
)

func Check(msg string, a ...any) {
	if 0 < len(a) {
		fmt.Println("✅ ", msg, a)
	} else {
		fmt.Println("✅ ", msg)
	}
}

func Error(msg string, a ...any) {
	if 0 < len(a) {
		fmt.Println("❌ ", msg, a)
	} else {
		fmt.Println("❌ ", msg)
	}
}

func Eyes(msg string, a ...any) {
	if 0 < len(a) {
		fmt.Println("👀 ", msg, a)
	} else {
		fmt.Println("👀 ", msg)
	}
}
