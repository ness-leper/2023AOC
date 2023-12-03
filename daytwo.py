import time
import json

def RenderGames(source:str, games: dict):
    game = source.split(": ")
    games[game[0]] = []
    # Split the dice pulls
    pulls = game[1].split("; ")
    for pull in pulls:
        games[game[0]].append(pull.strip())

    return games

def SolveDayTwo_PartOne(file:str):
    f = open(file, "r")

    # Read in each Game record, break into {ID:[games]}
    games = {}
    for line in f:
        games = RenderGames(line, games)

    f.close()

    # Validate if the game was possible
    valid = {}
    diceAvailable = {"red":12,"blue":14, "green": 13}
    invalidSum = 0
    totalSum = 0
    for game in games:
        gameValid = True
        for inst in games[game]:
            goodGame = True
            pull = inst.split(", ")
            for play in pull:
                for dice in diceAvailable:
                    if dice in play:
                        if diceAvailable[dice] < int(play.split(" ")[0]*1):
                            goodGame = False
            if goodGame == False:
                gameValid = False
                
        if gameValid == True:
            valid[game] = games[game]
        else:
            gameNum = game.split(" ")
            invalidSum += int(gameNum[1])
        gameNum = game.split(" ")
        totalSum += int(gameNum[1])

    sumGames = 0
    
    for game in valid:
        gameNum = game.split(" ")
        sumGames += int(gameNum[1])
            
    print(f"{totalSum}: Valid ({sumGames}) and Invalid ({invalidSum})")

# SolveDayTwo_PartOne("./d2test.txt");
SolveDayTwo_PartOne("./daytwoinput.txt")
