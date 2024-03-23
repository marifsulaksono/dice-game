This is an implementation of a dice game using Go as the programming language.

## Game Description

This game takes input for the number of players (N) and the number of dice per player (M). Each player receives M dice at the beginning of the game. After that, all players roll their dice simultaneously and evaluate based on the following rules:

1. Dice with the number 6 are taken as points for the player.
2. Dice with the number 1 are given to the player sitting next to them.
3. Dice with the numbers 2, 3, 4, and 5 are kept in play by the player.

The game continues until there is only one player left. The player with the most points is the winner.

## How to Run the Game

Make sure you have Go installed on your computer. To run the game, follow these steps:

1. Open a terminal and navigate to the project directory.
2. Type the command ```go run main.go``` to run the game.
3. Follow the instructions to enter the number of players and the number of dice.