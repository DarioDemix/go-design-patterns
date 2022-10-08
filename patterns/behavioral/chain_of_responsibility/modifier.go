package chain_of_responsibility

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func NewCreatureModifier(game *Game, creature *Creature) *CreatureModifier {
	return &CreatureModifier{game: game, creature: creature}
}

// DoubleAttackModifier doubles the attack of a creature
type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{*NewCreatureModifier(g, c)}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

type IncreaseDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(g *Game, c *Creature) *IncreaseDefenseModifier {
	idm := &IncreaseDefenseModifier{*NewCreatureModifier(g, c)}
	idm.game.Subscribe(idm)
	return idm
}

func (i *IncreaseDefenseModifier) Handle(q *Query) {
	if q.CreatureName == i.creature.Name && q.WhatToQuery == Defense {
		q.Value++
	}
}

func (i *IncreaseDefenseModifier) Close() error {
	i.game.Unsubscribe(i)
	return nil
}
