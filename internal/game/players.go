package game

type Player interface {
	Draw(deck *Deck)
	Discard(index int, deck *Deck)
	Fold()
	Bet(amount int) (int, error)
	GetSeat() *Seat
	IsActive() bool
}

func newPlayers() []Player {
	players := make([]Player, 0, 4)
	players = append(players, newUser())
	players = append(players, newSquirrelJak())
	players = append(players, newCobson())
	players = append(players, newChud())

	return players
}

type Seat struct {
	Hand     []*Card
	Chips    int
	Discards int
	Folded   bool
}

func newSeat() *Seat {
	return &Seat{
		Hand:     make([]*Card, 0, 5),
		Chips:    1000,
		Discards: 0,
		Folded:   false,
	}
}

type NPC interface {
	Player
	DecideAction(g *Game)
}

type User struct {
	*Seat
	Number int
}

func newUser() *User {
	return &User{
		Seat:   newSeat(),
		Number: 0,
	}
}

func (u *User) Draw(deck *Deck) {}

func (u *User) Discard(index int, deck *Deck) {}

func (u *User) Fold() {}

func (u *User) Bet(amount int) (int, error) {
	return 0, nil
}

func (u *User) GetSeat() *Seat {
	return u.Seat
}

func (u *User) IsActive() bool {
	return !u.Folded
}

type SquirrelJak struct {
	*Seat
	Number int
}

func newSquirrelJak() *SquirrelJak {
	return &SquirrelJak{
		Seat:   newSeat(),
		Number: 1,
	}
}

func (s *SquirrelJak) Draw(deck *Deck) {}

func (s *SquirrelJak) Discard(index int, deck *Deck) {}

func (s *SquirrelJak) Fold() {}

func (s *SquirrelJak) Bet(amount int) (int, error) {
	return 0, nil
}

func (s *SquirrelJak) GetSeat() *Seat {
	return s.Seat
}

func (s *SquirrelJak) IsActive() bool {
	return !s.Folded
}

func (s *SquirrelJak) DecideAction(g *Game) {}

type Cobson struct {
	*Seat
	Number int
}

func newCobson() *Cobson {
	return &Cobson{
		Seat:   newSeat(),
		Number: 2,
	}
}

func (c *Cobson) Draw(deck *Deck) {}

func (c *Cobson) Discard(index int, deck *Deck) {}

func (c *Cobson) Fold() {}

func (c *Cobson) Bet(amount int) (int, error) {
	return 0, nil
}

func (c *Cobson) GetSeat() *Seat {
	return c.Seat
}

func (c *Cobson) IsActive() bool {
	return !c.Folded
}

func (c *Cobson) DecideAction(g *Game) {}

type Chud struct {
	*Seat
	Number int
}

func newChud() *Chud {
	return &Chud{
		Seat:   newSeat(),
		Number: 3,
	}
}

func (c *Chud) Draw(deck *Deck) {}

func (c *Chud) Discard(index int, deck *Deck) {}

func (c *Chud) Fold() {}

func (c *Chud) Bet(amount int) (int, error) {
	return 0, nil
}

func (c *Chud) GetSeat() *Seat {
	return c.Seat
}

func (c *Chud) IsActive() bool {
	return !c.Folded
}

func (c *Chud) DecideAction(g *Game) {}
