package components

import (
	"phoenixbuilder/omega/components/qqGroupLink"
	"phoenixbuilder/omega/defines"
)

type BasicComponent struct {
	Config   *defines.ComponentConfig
	Frame    defines.MainFrame
	Ctrl     defines.GameControl
	Listener defines.GameListener
}

func (bc *BasicComponent) Init(cfg *defines.ComponentConfig) {
	bc.Config = cfg
}

func (bc *BasicComponent) Inject(frame defines.MainFrame) {
	bc.Frame = frame
	bc.Listener = frame.GetGameListener()
}

func (bc *BasicComponent) Activate() {
	bc.Ctrl = bc.Frame.GetGameControl()
}

func (bc *BasicComponent) Stop() error {
	return nil
}

func GetComponentsPool() map[string]func() defines.Component {
	return map[string]func() defines.Component{
		"Bonjour": func() defines.Component {
			return &Bonjour{BasicComponent: &BasicComponent{}}
		},
		"ChatLogger": func() defines.Component {
			return &ChatLogger{BasicComponent: &BasicComponent{}}
		},
		"Banner": func() defines.Component {
			return &Banner{BasicComponent: &BasicComponent{}}
		},
		"FeedBack": func() defines.Component {
			return &FeedBack{BasicComponent: &BasicComponent{}}
		},
		"Memo": func() defines.Component {
			return &Memo{BasicComponent: &BasicComponent{}}
		},
		"PlayerTP": func() defines.Component {
			return &PlayerTP{BasicComponent: &BasicComponent{}}
		},
		"BackToHQ": func() defines.Component {
			return &BackToHQ{BasicComponent: &BasicComponent{}}
		},
		"SetSpawnPoint": func() defines.Component {
			return &SetSpawnPoint{BasicComponent: &BasicComponent{}}
		},
		"Respawn": func() defines.Component {
			return &Respawn{BasicComponent: &BasicComponent{}}
		},
		"AboutMe": func() defines.Component {
			return &AboutMe{BasicComponent: &BasicComponent{}}
		},
		"Portal": func() defines.Component {
			return &Portal{BasicComponent: &BasicComponent{}}
		},
		"Immortal": func() defines.Component {
			return &Immortal{BasicComponent: &BasicComponent{}}
		},
		"Kick": func() defines.Component {
			return &Kick{BasicComponent: &BasicComponent{}}
		},
		"Shop": func() defines.Component {
			return &Shop{BasicComponent: &BasicComponent{}}
		},
		"QGroupLink": func() defines.Component {
			return &qqGroupLink.QGroupLink{}
		},
		"Recycle": func() defines.Component {
			return &Recycle{BasicComponent: &BasicComponent{}}
		},
		"FakeOP": func() defines.Component {
			return &FakeOp{BasicComponent: &BasicComponent{}}
		},
		"SimpleCmd": func() defines.Component {
			return &SimpleCmd{BasicComponent: &BasicComponent{}}
		},
		"Schedule": func() defines.Component {
			return &Schedule{BasicComponent: &BasicComponent{}}
		},
		"TimeSync": func() defines.Component {
			return &TimeSync{BasicComponent: &BasicComponent{}}
		},
		"MoneyTransfer": func() defines.Component {
			return &MoneyTransfer{BasicComponent: &BasicComponent{}}
		},
		"StructureBackup": func() defines.Component {
			return &StructureBackup{BasicComponent: &BasicComponent{}}
		},
		"Crash": func() defines.Component {
			return &Crash{BasicComponent: &BasicComponent{}}
		},
		"IntrusionDetectSystem": func() defines.Component {
			return &IntrusionDetectSystem{BasicComponent: &BasicComponent{}}
		},
		"WhoAreYou": func() defines.Component {
			return &WhoAreYou{BasicComponent: &BasicComponent{}}
		},
		"ContainerScan": func() defines.Component {
			return &ContainerScan{BasicComponent: &BasicComponent{}}
		},
		"OpCheck": func() defines.Component {
			return &OpCheck{BasicComponent: &BasicComponent{}}
		},
	}
}
