from Player import Player
from Board import Board

input("Press Enter to start the game...")
player = Player(str(input("플레이어의 이름을 입력하세요 >> ")))
computer = Player("컴퓨터")

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
