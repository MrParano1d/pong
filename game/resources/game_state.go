package resources

type GameState struct {
	Paused       bool
	Player1Score int
	Player2Score int
}

func NewGameState() *GameState {
	return &GameState{
		Paused:       false,
		Player1Score: 0,
		Player2Score: 0,
	}
}

func (g *GameState) Player1Scored() {
	g.Player1Score++
}

func (g *GameState) Player2Scored() {
	g.Player2Score++
}

func (g *GameState) Reset() {
	g.Player1Score = 0
	g.Player2Score = 0
}
