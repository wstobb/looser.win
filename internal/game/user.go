package game

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

func (u *User) Draw(deck *Deck) {
	u.Hand = append(u.Hand, deck.Draw())
}

func (u *User) Discard(index int, deck *Deck) {
	discarded := u.Hand[index]
	deck.Use(discarded)
	u.Draw(deck)
}

func (u *User) Fold() {
	u.Folded = true
}

func (u *User) Bet(amount int) (int, error) {
	return 0, nil
}

func (u *User) GetSeat() *Seat {
	return u.Seat
}

func (u *User) IsActive() bool {
	return !u.Folded
}
