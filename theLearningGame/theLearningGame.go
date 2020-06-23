package learninggameone

type Last struct {
	Cmd     int
	No      int
	learned bool
}

type Machine struct {
	triedCommands   map[int]int // [Cmd]No
	learnedCommands map[int]int // [Cmd]No
	last            Last
}

//Function called to get your machine initialised
func NewMachine() Machine {
	return Machine{triedCommands: make(map[int]int), learnedCommands: make(map[int]int)}
}

func (m *Machine) Command(cmd int, num int) int {
	if id, ok := m.learnedCommands[cmd]; ok {
		m.last = Last{Cmd: cmd, No: id, learned: true}
	} else {
		m.last = Last{Cmd: cmd, No: m.triedCommands[cmd], learned: false}
	}
	return Actions()[m.last.No](num)
}

func (m *Machine) Response(res bool) {
	if res {
		m.learnedCommands[m.last.Cmd] = m.last.No
		return
	}
	m.triedCommands[m.last.Cmd]++
	delete(m.learnedCommands, m.last.Cmd)
}
