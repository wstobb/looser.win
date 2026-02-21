package game

import "fmt"

type User struct {
	*Inventory
}

func NewUser() *User {
	return &User{
		Inventory: NewInventory(),
	}
}

func (u *User) Draw(deck *Deck) {
	u.Hand = append(u.Hand, deck.Draw())
}

func (u *User) Discard(index int, deck *Deck) {
	discarded := u.Hand[index]
	u.Hand = append(u.Hand, deck.Draw())
	deck.Use(discarded)
}

func (u *User) Fold() {
	u.Folded = true
}

func (u *User) Bet(amount int) (int, error) {
	if u.Chips-amount < 0 {
		return 0, fmt.Errorf("bet amount too high")
	}

	u.Chips = u.Chips - amount

	return amount, nil
}

func (u *User) Get() *Inventory {
	return u.Inventory
}

func (u *User) IsActive() bool {
	return !u.Folded
}
