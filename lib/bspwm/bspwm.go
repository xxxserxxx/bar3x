package bspwm

import (
	"bufio"
	"os/exec"
	"regexp"
)

type Desktop struct {
	Num     int64       `json:"num"`
	Name    string      `json:"name"`
	Visible bool        `json:"visible"`
	Focused bool        `json:"focused"`
	Urgent  bool        `json:"urgent"`
	Monitor string      `json:"monitor"`
}

//type Monitor struct {
//	Desktops []Desktop `json:"desktops"`
//}

type BSPWM struct {
	//Monitors []Monitor `json:monitors`
	desktops []Desktop
	notify func(b BSPWM)
}


var desktops []Desktop

func New(notify func(b BSPWM)) *BSPWM {
	b := &BSPWM{
		notify: notify,
	}

	go b.subscribe()
	return b
}

func (b *BSPWM) subscribe() {
	desktopRegExp := regexp.MustCompile(`[oOuUfF][^:]+`)
	montiorRegExp := regexp.MustCompile(`^WM[^:]+`)

	cmd := exec.Command("bspc", "subscribe", "report")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		bspcStatus := scanner.Text()
		monitor := montiorRegExp.FindString(bspcStatus)[2:]
		desktopMatches := desktopRegExp.FindAllString(bspcStatus, -1)

		var desktops []Desktop
		for _,bwk := range desktopMatches {
			name := bwk[1]
			status := bwk[0]

			desktops = append(desktops, Desktop{
				int64(name),
				string(name),
				true,
				status == 'O' || status == 'F',
				status == 'O' || status == 'F',
				monitor,
			})
		}

		b.setDesktops(desktops)
	}
	cmd.Wait()
}

func (b *BSPWM) setDesktops(dsks []Desktop) {
	desktops = dsks
	b.notify(*b)
}

func GetDesktops() []Desktop {
	return desktops
}