package helper

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ClearScreen() {
	osName := runtime.GOOS

	switch osName {
	case "linux", "darwin": // Untuk Linux dan MacOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Untuk Windows 10
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Clear screen tidak didukung pada sistem operasi ini")
	}
}

func InputString(param *string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	*param = strings.TrimSpace(input)
}

func InputInt(param *int) {
	fmt.Scan(param)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func WaitEnter() {
	fmt.Print("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Scan()
}
