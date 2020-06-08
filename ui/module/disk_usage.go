package module

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/shimmerglass/bar3x/ui"
	"github.com/shimmerglass/bar3x/ui/markup"
	"github.com/shirou/gopsutil/disk"
)

type DiskUsage struct {
	moduleBase

	mk    *markup.Markup
	clock *Clock

	Txt *TextUnit
}

func NewDiskUsage(p ui.ParentDrawable, mk *markup.Markup, clock *Clock) *DiskUsage {
	return &DiskUsage{
		mk:         mk,
		clock:      clock,
		moduleBase: newBase(p),
	}
}

func (m *DiskUsage) Init() error {
	_, err := m.mk.Parse(m, m, `
		<Sizer ref="Root" Height="{height}">
			<Row>
				<Icon>{icons["disk"]}</Icon>
				<Sizer PaddingLeft="{h_padding}">
					<TxtUnit ref="Txt" />
				</Sizer>
			</Row>
		</Sizer>
	`)
	if err != nil {
		return err
	}

	m.clock.Add(m, 10*time.Second)
	return nil
}

func (m *DiskUsage) Update() {
	stat, err := disk.Usage("/")
	var p float64
	if err != nil {
		log.Println(err)
		p = 0
	} else {
		p = stat.UsedPercent
	}

	m.Txt.Set(fmt.Sprintf("%.0f", p), "%")
}