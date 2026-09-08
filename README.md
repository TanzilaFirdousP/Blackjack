# ♠️ Blackjack (Go)

A command-line implementation of the classic **Blackjack** card game built in **Go**. This project simulates a full player vs dealer experience while demonstrating core object-oriented programming concepts in Go such as structs, interfaces, composition, and polymorphism.

## Gameplay

* 52-card deck with randomized shuffling
* Player vs Dealer gameplay
* Hit or Stand decision system
* Automatic dealer turn following Blackjack rules
* Ace value handling (1 or 11)
* Win/Lose/Draw outcome detection
* Input validation for player choices

## Tech Stack

* **Language:** Go
* **Type:** Command Line Interface (CLI)

## Concepts Demonstrated

This project was built to practice fundamental Go programming concepts:

* Structs & Methods
* Interfaces
* Abstraction
* Polymorphism
* Composition (embedded structs)
* Randomization & Shuffling
* Control Flow
* User Input Validation

## How to Run

1. Clone the repository

```bash
git clone https://github.com/TanzilaFirdousP/Blackjack.git
```

2. Navigate into the project

```bash
cd Blackjack
```

3. Run the game

```bash
go run main.go
```

## Sample Gameplay

```text
Your cards:
A♠  8♦
Score: 19

Choose:
(H) Hit
(S) Stand

> s

Dealer reveals:
10♣  7♥
Dealer Score: 17

🎉 You Win!
```
