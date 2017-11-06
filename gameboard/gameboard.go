package gameboard

var mScroll map[string][]string

var games Allgames

type request struct {
	Key         string
	MatchString string
	Position    int
}

type Allgames struct {
	//TODO: rename, smh
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

//TODO update current player after update board
func UpdateBoard(action int, player int, gamekey string) string {
	ActiveGameBoard := games.games[gamekey]
	if player != games.games[gamekey].currentTurn {
		return "Error: Not your turn"
	}

	switch action {
	case 1:
		{
			if ActiveGameBoard.board[action] > 0 {
				return "Error: position already taken"
			}
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
