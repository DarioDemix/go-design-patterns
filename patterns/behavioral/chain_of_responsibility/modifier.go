package chain_of_responsibility

import "fmt"

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Apply() {
	if c.next != nil {
		c.next.Apply()
	}
}

// DoubleAttackModifier doubles the attack of a creature
type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier: *NewCreatureModifier(c)}
}

func (d *DoubleAttackModifier) Apply() {
	fmt.Printf("Doubling %s's attack\n", d.creature.Name)
	d.creature.Attack *= 2
	d.CreatureModifier.Apply()
}

// IncreaseDefenseModifier increases the defense of a creature
type IncreaseDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(c *Creature) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{CreatureModifier: *NewCreatureModifier(c)}
}

func (i *IncreaseDefenseModifier) Apply() {
	if i.creature.Attack <= 2 {
		fmt.Printf("Increasing %s's defense\n", i.creature.Name)
		i.creature.Defense++
	}
	i.CreatureModifier.Apply()
}

// NoBonusesModifier disables all the other modifiers
type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{creature: c}}
}

func (n *NoBonusesModifier) Apply() {
	// do not call others modifier
	fmt.Printf("Disabling modifiers for %s\n", n.creature.Name)
}
