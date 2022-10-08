package chain_of_responsibility

import (
	"fmt"
	"sync"
)

/*
Command Query Separation

Command = asking for an action or change (e.g. set your attack value to 2)
Query = asking for information (e.g. give me your attack value)
CQS = having separate means of sending commands and queries to e.g., direct field access
*/

type Modifier interface {
	Add(m Modifier)
	Apply()
}

type ChainOfResponsibility struct {
}

func (c *ChainOfResponsibility) Demo() {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.String())

	dam := NewDoubleAttackModifier(game, goblin)
	fmt.Println(goblin.String())
	_ = dam.Close()

	idm := NewIncreaseDefenseModifier(game, goblin)
	fmt.Println(goblin.String())
	_ = idm.Close()

	fmt.Println(goblin.String())
}
