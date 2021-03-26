package main

//
// Go has some limitations WRT to getting pointers.
//
// https://utcc.utoronto.ca/~cks/space/blog/programming/GoAddressableValues
func myName() string {
	return "Goku"
}

type Warrior interface {
	fightingStyle() string
	makeDisciple() *Warrior
}

type Monk struct {
	school string
}

func (monk *Monk) fightingStyle() string {
	return monk.school + " Wing Chun"
}

// Because interfaces are not about Types but Methods
// this is not allowed in Golang.
func (master *Monk) makeDisciple() *Warrior {
	return &Monk{
		school: master.school,
	}
}

// Here we return a Pointer to the struct, which is ok
func makeDisciple(master *Monk) *Monk {
	return &Monk{
		school: master.school,
	}
}

type Samurai struct {
	master string
}

func (samurai *Samurai) fightingStyle() string {
	return samurai.master + " Iai Jutsu"
}

type Ninja struct {
	village string
}

func main() {
	ages := map[string]int{
		"Mary":   30,
		"Zorgon": 100,
	}

	age := &ages["Mary"]
	name := &myName()

	nameAgain := myName()
	nameAgainPtr := &nameAgain
	println(nameAgainPtr)
}
