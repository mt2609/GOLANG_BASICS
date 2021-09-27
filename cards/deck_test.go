package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	//T1
	if len(d) != 52 {
		t.Errorf("expected 52, but got %v", len(d))
	}
	//T2

	if d[0] != "Ace of spades" {
		t.Errorf("exoected AofS but got %v", d[0])
	}
	//T3
	if d[len(d)-1] != "King of hearts" {
		t.Errorf("expected KofH but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") //Removes existing tests

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 50 {
		t.Errorf("expected 52 but got %v", len(loadedDeck))

	}
	os.Remove("_decktesting")
}
