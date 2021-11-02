import random

class Player:
    def __init__(self, name, pos=0):
        self.name = name
        self.pos = pos
    
    def getName(self):
        return self.name
    
    def getPos(self):
        return self.pos
    
    def move(self):
        self.pos += random.randint(0,5)