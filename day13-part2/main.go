package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

const directionNorth = "^"
const directionEast = ">"
const directionSouth = "v"
const directionWest = "<"

const trackHorizontal = "-"
const trackVertical = "|"
const trackJunction = "+"
const trackForwardSlash = "/"
const trackBackSlash = "\\"

const turnLeft = "LEFT"
const turnStraight = "STRAIGHT"
const turnRight = "RIGHT"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Cart struct {
	position  Point
	direction string
	turnIndex int
	collided  bool
}

func (c *Cart) Move(tracks []string) {
	switch c.direction {
	case directionNorth:
		c.position.y--
	case directionEast:
		c.position.x++
	case directionSouth:
		c.position.y++
	case directionWest:
		c.position.x--
	}
	trackRow := tracks[c.position.y]
	switch string(trackRow[c.position.x]) {
	case trackJunction:
		c.TurnAtJunction()
	case trackForwardSlash:
		switch c.direction {
		case directionNorth:
			c.TurnRight()
		case directionEast:
			c.TurnLeft()
		case directionSouth:
			c.TurnRight()
		case directionWest:
			c.TurnLeft()
		}
	case trackBackSlash:
		switch c.direction {
		case directionNorth:
			c.TurnLeft()
		case directionEast:
			c.TurnRight()
		case directionSouth:
			c.TurnLeft()
		case directionWest:
			c.TurnRight()
		}
	}
}

func (c *Cart) TurnAtJunction() {
	turns := []string{turnLeft, turnStraight, turnRight}
	turn := turns[c.turnIndex]
	if turn == turnLeft {
		c.TurnLeft()
	} else if turn == turnRight {
		c.TurnRight()
	}
	c.turnIndex = (c.turnIndex + 1) % len(turns)
}

func (c *Cart) TurnLeft() {
	switch c.direction {
	case directionNorth:
		c.direction = directionWest
	case directionEast:
		c.direction = directionNorth
	case directionSouth:
		c.direction = directionEast
	case directionWest:
		c.direction = directionSouth
	}
}

func (c *Cart) TurnRight() {
	switch c.direction {
	case directionNorth:
		c.direction = directionEast
	case directionEast:
		c.direction = directionSouth
	case directionSouth:
		c.direction = directionWest
	case directionWest:
		c.direction = directionNorth
	}
}

func (c Cart) HasCollided(carts []Cart) bool {
	for _, otherCart := range carts {
		if otherCart.collided {
			continue
		}
		if c.position == otherCart.position {
			return true
		}
	}
	return false
}

type Mine struct {
	carts      []Cart
	tracks     []string
	collisions []Point
}

type Point struct {
	x int
	y int
}

func trackForCartDirection(cartDirection string) string {
	switch cartDirection {
	case directionNorth:
		return trackVertical
	case directionSouth:
		return trackVertical
	case directionEast:
		return trackHorizontal
	case directionWest:
		return trackHorizontal
	}
	return ""
}

func (m *Mine) ParseMap(data string) {
	rows := strings.Split(string(data), "\n")
	for i, row := range rows {
		m.tracks = append(m.tracks, row)
		re := regexp.MustCompile("([<>v^])")
		cartIndexes := re.FindAllStringIndex(row, -1)
		for _, cartIndex := range cartIndexes {
			m.AddCart(cartIndex[0], i, string(row[cartIndex[0]]))
		}
	}
	m.AddTilesForCarts()
}

func (m *Mine) AddCart(x int, y int, direction string) {
	cart := Cart{Point{x, y}, direction, 0, false}
	m.carts = append(m.carts, cart)
}

func (m *Mine) AddTilesForCarts() {
	for _, cart := range m.carts {
		m.replaceCartTrack(cart)
	}
}

func (m *Mine) replaceCartTrack(c Cart) {
	row := m.tracks[c.position.y]
	row = row[:c.position.x] + trackForCartDirection(c.direction) + row[c.position.x+1:]
	m.tracks[c.position.y] = row
}

func (m *Mine) FirstCollision() Point {
	for {
		collisionPoint := m.MoveCarts()
		if collisionPoint != nil {
			return *collisionPoint
		}
	}
}

func (m *Mine) LastCart() Point {
	for {
		lastCartPoint := m.MoveCarts()
		if lastCartPoint != nil {
			return *lastCartPoint
		}
	}
}

func (m *Mine) MoveCarts() *Point {
	sort.Sort(CartsByPosition(m.carts))
	var lastPoint *Point
	collisionDuringRound := false
	for i := range m.carts {
		cart := &m.carts[i]
		if cart.collided {
			continue
		}
		cart.Move(m.tracks)
		otherCarts := append(make([]Cart, len(m.carts)-1), m.carts[:i]...)
		otherCarts = append(otherCarts, m.carts[i+1:]...)
		collided := cart.HasCollided(otherCarts)
		collisionDuringRound = collisionDuringRound || collided
		if collided {
			fmt.Println("Before Collision!!!!!", m.NoCollisionCount())
			m.MarkCartsCollided(cart.position)
			fmt.Println("After Collision!!!!!", m.NoCollisionCount())
			for _, cart := range m.carts {
				if !cart.collided {
					fmt.Println(cart)
				}
			}
		} else {
			lastPoint = &cart.position
		}
	}
	if m.NoCollisionCount() == 1 {
		return lastPoint
	}
	return nil
}

func (m *Mine) MarkCartsCollided(point Point) {
	for i := range m.carts {
		if m.carts[i].collided {
			continue
		}
		if m.carts[i].position == point {
			m.carts[i].collided = true
		}
	}
}

func (m Mine) NoCollisionCount() int {
	count := 0
	for _, cart := range m.carts {
		if !cart.collided {
			count++
		}
	}
	return count
}

type CartsByPosition []Cart

func (a CartsByPosition) Len() int {
	return len(a)
}

func (a CartsByPosition) Less(i int, j int) bool {
	return a[i].position.y < a[j].position.y ||
		a[i].position.y == a[j].position.y && a[i].position.x < a[j].position.x
}

func (a CartsByPosition) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (m Mine) PrintTrack() {
	for _, row := range m.tracks {
		fmt.Println(row)
	}
}

func (m Mine) PrintAll() {
	for y, row := range m.tracks {
		for _, cart := range m.carts {
			if cart.position.y == y {
				track := cart.direction
				switch string(row[cart.position.x]) {
				case directionNorth:
					track = "X"
				case directionEast:
					track = "X"
				case directionSouth:
					track = "X"
				case directionWest:
					track = "X"
				}
				row = row[:cart.position.x] + track + row[cart.position.x+1:]
			}
		}
		fmt.Println(row)
	}
}

func newMine() Mine {
	return Mine{}
}

func main() {
	data, err := ioutil.ReadFile("input/d13-input.txt")
	check(err)

	m := newMine()
	m.ParseMap(string(data))
	collisionPoint := m.LastCart()
	fmt.Println("Last Cart", collisionPoint)
	m.PrintAll()

	fmt.Println("Done")
}
