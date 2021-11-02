from clearConsole import clearConsole as clear


class Board:
    def __init__(self, player, computer):
        self.player = player
        self.computer = computer
    
    def prtBoard(self):
        pPos = self.player.getPos()

        cPos = self.computer.getPos()

        clear()

        print(f"{self.player.getName()} : ", end="")
        for i in range(0, 30):
            if i == pPos:
                print("♟️", end="")
            else:
                print("⬜", end="")
        print("🎈")

        print(f"{self.computer.getName()} : ", end="")
        for i in range(0, 30):
            if i == cPos:
                print("♟️", end="")
            else:
                print("⬜", end="")
        print("🎈")

    def evaluate(self):
        if(self.player.getPos() >= 30):
            print("You win!")
            return False
        elif(self.computer.getPos() >= 30):
            print("Computer wins!")
            return False
        else:
            return True