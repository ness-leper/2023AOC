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
    diceAvailable = {"red":12,"blue":13, "green": 14}
    for game in games:
        gameValid = True
        for inst in games[game]:
            goodGame = True
            pull = inst.split(", ")
            if len(pull) > 3:
                print(inst)
                exit()
            for play in pull:
                for dice in diceAvailable:
                    if dice in play:
                        if diceAvailable[dice] < int(play.split(" ")[0]*1):
                            goodGame = False
            if goodGame == False:
                gameValid = False
                
        if gameValid == True:
            valid[game] = games[game]
            # print(f"Good Game: {game} {games[game]}")
            # time.sleep(2)

    sumGames = 0
    
    for game in valid:
        gameNum = game.split(" ")
        sumGames += int(gameNum[1])
            
    print(sumGames)

# SolveDayTwo_PartOne("./d2test.txt");
SolveDayTwo_PartOne("./daytwoinput.txt")
