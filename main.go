package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	dices               []int
	score               int
	additionalDiceCount int
}

func (player *Player) RollDice() {
	rand.Seed(time.Now().UnixNano())

	for i := range player.dices {
		player.dices[i] = rand.Intn(6) + 1
	}
}

func (player *Player) PassingDice(targetPlayer *Player) {
	newDices := []int{}

	for i := range player.dices {
		if player.dices[i] == 1 {
			targetPlayer.ReceiveAdditionalDice()
		} else {
			newDices = append(newDices, player.dices[i])
		}
	}

	player.dices = newDices
}

func (player *Player) ReceiveAdditionalDice() {
	player.additionalDiceCount++
}

func (player *Player) CollectAdditionalDice() {
	for i := 1; i <= player.additionalDiceCount; i++ {
		player.dices = append(player.dices, 1)
	}
	player.additionalDiceCount = 0
}

func (player *Player) ScoreDice() {
	newDices := []int{}

	for i := range player.dices {
		if player.dices[i] == 6 {
			player.score++
		} else {
			newDices = append(newDices, player.dices[i])
		}
	}

	player.dices = newDices
}

func (player *Player) HasDice() bool {
	return len(player.dices) > 0
}

type Game struct {
	players []Player
}

func NewGame(playerNumber int, diceNumber int) *Game {
	game := &Game{}

	for i := 0; i < playerNumber; i++ {
		game.players = append(game.players, Player{dices: make([]int, diceNumber), score: 0, additionalDiceCount: 0})
	}

	return game
}

func (game *Game) RollPlayers() {
	for i := range game.players {
		if len(game.players[i].dices) > 0 {
			game.players[i].RollDice()
			fmt.Printf("Pemain #%d (%d): %v\n", i+1, game.players[i].score, game.players[i].dices)
		} else {
			fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", i+1, game.players[i].score)
		}
	}
}

func (game *Game) GetNextPlayer(currentIndex int) *Player {
	var nextPlayer *Player

	for {
		currentIndex++
		currentIndex = currentIndex % len(game.players)

		currentPlayer := &game.players[currentIndex]
		if len(currentPlayer.dices) > 0 {
			nextPlayer = currentPlayer
			break
		}
	}

	return nextPlayer
}

func (game *Game) EvaluatePlayers() {
	for i := range game.players {
		nextPlayer := game.GetNextPlayer(i)
		game.players[i].PassingDice(nextPlayer)
		game.players[i].ScoreDice()
	}

	for i := range game.players {
		if len(game.players[i].dices) > 0 {
			game.players[i].CollectAdditionalDice()
			fmt.Printf("Pemain #%d (%d): %v\n", i+1, game.players[i].score, game.players[i].dices)
		} else {
			fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", i+1, game.players[i].score)
		}
	}
}

func (game *Game) IsFinish() bool {
	playerHasDiceCount := 0

	for i := range game.players {
		if game.players[i].HasDice() {
			playerHasDiceCount++
		}
	}

	return playerHasDiceCount <= 1
}

func (game *Game) GetHighScore() int {
	hightScore := -1

	for i := range game.players {
		if game.players[i].score > hightScore {
			hightScore = game.players[i].score
		}
	}

	return hightScore
}

func (game *Game) GetWinnersIndex() []int {
	var winnersIndex []int
	highScore := game.GetHighScore()

	for i := range game.players {
		if game.players[i].score == highScore {
			winnersIndex = append(winnersIndex, i)
		}
	}

	return winnersIndex
}

func (game *Game) GetFirstPlayerHasDiceIndex() int {
	index := -1
	for i := range game.players {
		if len(game.players[i].dices) > 0 {
			index = i
			break
		}
	}
	return index
}

func (game *Game) Play() {
	playerNumber := len(game.players)
	diceNumber := len(game.players[0].dices)

	fmt.Printf("Pemain = %d, Dadu = %d\n", playerNumber, diceNumber)

	turnNumber := 1
	for !game.IsFinish() {
		fmt.Println("====================")
		fmt.Printf("Giliran %d lempar dadu:\n", turnNumber)
		game.RollPlayers()

		fmt.Println("Setelah evaluasi:")
		game.EvaluatePlayers()

		turnNumber++
	}

	playerHasDiceIndex := game.GetFirstPlayerHasDiceIndex()

	fmt.Println("====================")

	if playerHasDiceIndex != -1 {
		fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", playerHasDiceIndex+1)
	} else {
		fmt.Println("Game berakhir karena semua pemain kehabisan dadu.")
	}

	for _, value := range game.GetWinnersIndex() {
		fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", value+1)
	}
}

func main() {
	playerNumber := 3
	diceNumber := 4
	game := NewGame(playerNumber, diceNumber)
	game.Play()
}
