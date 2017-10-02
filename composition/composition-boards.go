package main

import (
	"fmt"
)

/*
We have a carpenter that we hired to renovate our house. In particular, there are some boards in the house that have rotted and need to be yanked out, as well as new boards that need to be nailed in. The carpenter will be given a supply of nails, boards to work with and tools to perform the work.
*/

// Board is a representation of a board with NailsNeeded showing nails needed to make it driven completely and total nails driven at present
type Board struct {
	NailsNeeded int
	NailsDriven int
}

// NailDriver is an interface for the tools which can do nail driving
type NailDriver interface {
	DriveNail(board *Board, n *int)
}

// NailPuller is an interface for the tools which can pull the nails
type NailPuller interface {
	PullNail(board *Board, n *int)
}

// NailDriverPuller is an interface for the tools which can both drive and pull the nails
type NailDriverPuller interface {
	NailDriver
	NailPuller
}

// Crowbar is a tool used for pulling nails
type Crowbar struct{}

// Mallet is a hammer tool used for pounding nails
type Mallet struct{}

// Crowbar implementing NailDriver interface
func (m *Mallet) DriveNail(b *Board, nailSupply *int) {
	b.NailsDriven++
	*nailSupply--
	fmt.Println("Mallet pounded a nail")
}

// Mallet implementing NailPuller interface
func (c *Crowbar) PullNail(b *Board, nailSupply *int) {
	b.NailsDriven--
	*nailSupply++
	fmt.Println("Crowbar pulled out a nail")
}

// Toolbox can contain any type of nail driver and puller and any amount of nails
type ToolBox struct {
	NailDriver
	NailPuller

	nails int
}

// Carpenter
type Carpenter struct{}

func (c *Carpenter) PullOutNails(b *Board, np NailPuller, nailSupply *int) {
	for b.NailsDriven > b.NailsNeeded {
		np.PullNail(b, nailSupply)
	}
}

func (c *Carpenter) DriveInNails(b *Board, nd NailDriver, nailSupply *int) {
	for b.NailsDriven < b.NailsNeeded {
		nd.DriveNail(b, nailSupply)
	}
}

func (c *Carpenter) ProcessBoards(boards []Board, ndp NailDriverPuller, nailSupply *int) {
	fmt.Println("Processing boards...")
	for i := range boards {
		board := &boards[i]
		switch {
		case board.NailsDriven < board.NailsNeeded:
			c.DriveInNails(board, ndp, nailSupply)
		case board.NailsDriven > board.NailsNeeded:
			c.PullOutNails(board, ndp, nailSupply)
		}
	}
}

func displayState(tb *ToolBox, boards []Board) {
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}

	fmt.Println()
}

func main() {

	boards := []Board{
		//rotten boards
		{NailsDriven: 2},
		{NailsDriven: 3},
		{NailsDriven: 4},

		//new boards
		{NailsNeeded: 1},
		{NailsNeeded: 3},
		{NailsNeeded: 4},
	}

	toolBox := ToolBox{
		NailDriver: &Mallet{},
		NailPuller: &Crowbar{},
		nails:      10,
	}

	displayState(&toolBox, boards)
	carpentar := &Carpenter{}
	carpentar.ProcessBoards(boards, &toolBox, &toolBox.nails)
	displayState(&toolBox, boards)

}
