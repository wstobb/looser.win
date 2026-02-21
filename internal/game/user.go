package game

type User struct {
	*Inventory
}

func NewUser() *User {
	return &User{
		Inventory: NewInventory(),
	}
}

func (u *User) Draw(deck *Deck)               {}
func (u *User) Discard(index int, deck *Deck) {}
func (u *User) Fold()                         {}

func (u *User) Bet(amount int) (int, error) {
	return 0, nil
}

func (u *User) Get() *Inventory {
	return u.Inventory
}

func (u *User) IsActive() bool {
	return true
}
