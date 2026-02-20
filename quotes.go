package main

import "fmt"

// this does not need to be its own file but i dont like looking at long walls of quote text
// going to spend a bit of time finding fav quotes
var quotes = []string{
	"I do bite my thumb, sir",
	"Would thou wert clean enough to spit upon!",
	"I do desire we may be strangers",
	"Thou whoreson zed, thou unnecessary letter!",
	"You Banbury cheese!",
	"I am sick when I do look on thee",
	"I do desire we may be better strangers",
	"There’s no more faith in thee than in a stewed prune",
	"To be, or not to be, that is the question",
	"All the world's a stage",
	"And all the men and women merely players",
	"The web of our life is of a mingled yarn, good and ill together",
	"I love long life better than figs",
	"This tiger-footed rage",
	"Come not between the dragon and his wrath!",
	"I understand a fury in your words",
	"But not the words",
	"Who is man that is not angry",
	"Heat not a furnace for your foe so hot",
	"That it do singe yourself",
	"Convert to anger",
	"Blunt not the heart",
	"Enrage it",
	"There is no following her in this fierce vein",
	"I would my father look'd but with my eyes.",
	"It is a wise father that knows his own child",
}

func count() {
	fmt.Printf("count of quotes: %d", len(quotes))
}
