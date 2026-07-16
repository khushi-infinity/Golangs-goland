# Mini Casino

A simple terminal-based slot machine game written in Go.

## Overview

You start with a balance of **$200**, enter your name, place a bet, and spin a 3x3 slot grid. If a row contains the same symbol across all three slots, you win based on that symbol's multiplier.

## Features

- Console-based gameplay
- Player name prompt
- Bet validation against current balance
- 3x3 random slot spins
- Win payouts based on matching rows
- Quit anytime by entering `0`

## Symbol Rules

The game uses four symbols with different frequencies and payouts:

| Symbol | Count | Multiplier |
|--------|-------|------------|
| A      | 4     | 20x        |
| B      | 7     | 10x        |
| C      | 12    | 5x         |
| D      | 20    | 2x         |

## How to Run

```bash
go run .
```

Or build and run:

```bash
go build -o mini-casino
./mini-casino
```

## Gameplay Flow

1. Enter your name.
2. Choose a bet amount.
3. View the generated slot grid.
4. Receive winnings for any matching rows.
5. Continue until you quit or your balance reaches zero.

## Requirements

- Go 1.23.4 or newer

