from Player import Player
from Board import Board

input("Press Enter to start the game...")
player = Player(str(input("Pleas input player name >> ")))
computer = Player("computer")

board = Board(player, computer)

board.prtBoard()

while(board.evaluate()):
    input("Press Enter to Next turn...")
    player.move()
    board.prtBoard()

    if not board.evaluate():
        break

    input("Press Enter to Next turn...")
    computer.move()
    board.prtBoard()
