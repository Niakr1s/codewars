package kata

import (
	"fmt"
	"log"
)

type Bullet struct {
	Speed     uint32
	Travelled uint32
}

func NewBullet(speed uint32) *Bullet {
	return &Bullet{
		Speed:     speed,
		Travelled: 0,
	}
}

func (b *Bullet) IsHit(dist uint32) bool {
	return dist <= b.Travelled
}

func (b *Bullet) Step() {
	b.Travelled += b.Speed
}

func (b *Bullet) String() string {
	return fmt.Sprint(*b)
}

type Soldier struct {
	InitialPos int
	RifleSpeed uint32
	IsDead     bool
}

func NewSoldier(initialPos int, rifleSpeed uint32) *Soldier {
	return &Soldier{
		InitialPos: initialPos,
		RifleSpeed: rifleSpeed,
	}
}

func MakeSoldiers(troop []uint32) []*Soldier {
	res := []*Soldier{}
	for i, rifleSpeed := range troop {
		res = append(res, NewSoldier(i, rifleSpeed))
	}
	return res
}

func (s *Soldier) String() string {
	return fmt.Sprint(*s)
}

type Army struct {
	InitialPos int
	Soldiers   []*Soldier
	IncBullets []*Bullet
}

func NewArmy(troop []uint32, initialPos int) *Army {
	return &Army{
		InitialPos: initialPos,
		Soldiers:   MakeSoldiers(troop),
		IncBullets: []*Bullet{},
	}
}

func (a *Army) ShiftSoldiers() {
	if len(a.Soldiers) == 0 {
		return
	}
	first := a.Soldiers[0]
	a.Soldiers = a.Soldiers[1:]
	if !first.IsDead {
		a.Soldiers = append(a.Soldiers, first)
	}
}

func (a *Army) IsEliminated() bool {
	return len(a.Soldiers) == 0
}

func (a *Army) String() string {
	res := fmt.Sprintf("Army with initialPos=%d\n", a.InitialPos)
	res += fmt.Sprintf("Soldiers: %v\n", a.Soldiers)
	res += fmt.Sprintf("IncBullets: %v\n", a.IncBullets)
	return res
}

func MakeArmies(troops ...[]uint32) []*Army {
	armies := make([]*Army, len(troops))
	for i, troop := range troops {
		armies[i] = NewArmy(troop, i)
	}
	return armies
}

type Game struct {
	Armies    []*Army
	Distance  uint32
	iteration int
}

func NewGame(armies []*Army, dist uint32) *Game {
	return &Game{
		Armies:   armies,
		Distance: dist,
	}
}

func (g *Game) Shoot() {
	for i, army := range g.Armies {
		adjacentArmy := g.Armies[(i+1)%len(g.Armies)]
		firstSoldier := army.Soldiers[0]
		if !firstSoldier.IsDead {
			adjacentArmy.IncBullets = append(adjacentArmy.IncBullets, NewBullet(army.Soldiers[0].RifleSpeed))
		}
	}
}

func (g *Game) StepBullets() {
	for _, army := range g.Armies {
		for _, bullet := range army.IncBullets {
			bullet.Step()
		}
	}
}

func (g *Game) MarkDead() {
loop:
	for _, army := range g.Armies {
		for _, bullet := range army.IncBullets {
			if bullet.IsHit(g.Distance) {
				army.Soldiers[0].IsDead = true
				continue loop
			}
		}
	}
}

func (g *Game) RemoveHitBullets() {
	for _, army := range g.Armies {
		nonHitBullets := []*Bullet{}
		for _, bullet := range army.IncBullets {
			if !bullet.IsHit(g.Distance) {
				nonHitBullets = append(nonHitBullets, bullet)
			}
		}
		army.IncBullets = nonHitBullets
	}
}

func (g *Game) ShiftSoldiers() {
	for _, army := range g.Armies {
		army.ShiftSoldiers()
	}
}

func (g *Game) RemoveDeadArmies() {
	aliveArmies := []*Army{}
	anyEleminated := false
	for _, army := range g.Armies {
		if !army.IsEliminated() {
			aliveArmies = append(aliveArmies, army)
		} else {
			anyEleminated = true
		}
	}
	g.Armies = aliveArmies
	if anyEleminated {
		g.removeAllBullets()
	}
}

func (g *Game) removeAllBullets() {
	for _, army := range g.Armies {
		army.IncBullets = []*Bullet{}
	}
}

func (g *Game) Step() {
	g.StepBullets()
	g.MarkDead()
	g.RemoveHitBullets()
	g.Shoot()
	g.ShiftSoldiers()
	g.RemoveDeadArmies()
	g.iteration++
}

func (g *Game) IsFinished() bool {
	return len(g.Armies) <= 1
}

func (g *Game) GetResult() (int, []uint32) {
	if len(g.Armies) == 0 {
		return -1, []uint32{}
	}
	army := g.Armies[0]

	armyInitialPos := army.InitialPos
	survivedSoldiers := []uint32{}

	for _, survivedSoldier := range army.Soldiers {
		survivedSoldiers = append(survivedSoldiers, uint32(survivedSoldier.InitialPos))
	}

	return armyInitialPos, survivedSoldiers
}

func (g *Game) String() string {
	res := ""
	res += fmt.Sprintf("---------------\n")
	res += fmt.Sprintf("Iteration #%d\n\n", g.iteration)
	for _, army := range g.Armies {
		res += fmt.Sprint(army)
		res += "\n"
	}
	res += fmt.Sprintf("---------------\n")
	return res
}

func QueueBattle(dist uint32, troops ...[]uint32) (int, []uint32) {
	game := NewGame(MakeArmies(troops...), dist)

	log.Print(game)
	log.Print("START")

	for !game.IsFinished() {
		game.Step()
		log.Print(game)
	}

	return game.GetResult()
}
