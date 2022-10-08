package chain_of_responsibility

import (
	"fmt"
	"sync"
)

type Argument int

const (
	Attack = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value any) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type Creature struct {
	game            *Game
	Name            string
	attack, defense int
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{game: game, Name: name, attack: attack, defense: defense}
}

func (c *Creature) Attack() int {
	q := &Query{c.Name, Attack, c.attack}
	c.game.Fire(q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := &Query{c.Name, Defense, c.defense}
	c.game.Fire(q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}
