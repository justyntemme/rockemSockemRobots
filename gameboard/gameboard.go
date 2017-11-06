package gameboard

var mScroll map[string][]string

var games Allgames

type request struct {
	Key         string
	MatchString string
	Row         int
	Col         int
}

type Allgames struct {
	games map[string]gameboard
}

type gameboard struct {
	/*
		1,2,3
		4,5,6
		7,8,9



	*/
	board [9]int
	//1 = player1, 2= player2
	currentTurn int
}

func UpdateBoard(action int, player int, gamekey string) string {
	ActiveGameBoard := games.games[gamekey]

	switch action {
	case 1:
		{

		}
	}

	return "Success"

}

func VerifyRequest(req request) string {
	var player = 0
	if mScroll[req.MatchString][0] == req.Key {
		player = 1
	} else if mScroll[req.MatchString][1] == req.Key {
		player = 2
	} else {
		return ("Error Incorrect PrivateKey/MatchKey")
	}

	return string(player)

}

func Init() {

	mScroll := make(map[string][]string)

}
