package chain_of_responsibility

import "fmt"

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
	fmt.Println("Chain Of Responsibility demo:")

	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	root.Add(NewNoBonusesModifier(goblin))

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreaseDefenseModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))

	root.Apply()
	fmt.Println(goblin.String())
}
