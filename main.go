package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/go-vgo/robotgo"
)

// openBrowser tries to open the URL in a browser,
// and returns whether it succeed in doing so.
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func openSource() {
	openBrowser("https://google.com")

	robotgo.Sleep(3)

	robotgo.ActiveName("chrome")

	// find the process id by the process name
	fpid, err := robotgo.FindIds("chrome")
	if err == nil {
		if len(fpid) > 0 {
			robotgo.ActivePID(fpid[0])

			tl := robotgo.GetTitle(fpid[0])
			fmt.Println("pid[0] title is: ", tl)

			x, y, w, h := robotgo.GetBounds(fpid[0])
			fmt.Println("GetBounds is: ", x, y, w, h)

			robotgo.TypeStr("moscow temperature")
			robotgo.KeyPress("enter")

			sx, sy := robotgo.GetScreenSize()
			fmt.Println("get screen size: ", sx, sy)

			robotgo.MoveSmooth(10, 10)
			robotgo.ScrollSmooth(-50, 6)

			//bit := robotgo.CaptureScreen(0, 0, 100, 100)
			//defer robotgo.FreeBitmap(bit)

			//img := robotgo.ToImage(bit)
			//robotgo.SavePng(img, "test_1.png")

			robotgo.Sleep(10)
			robotgo.Kill(fpid[0])
		}
	}
}

func main() {
	openSource()
}
