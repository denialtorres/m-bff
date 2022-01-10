package bff

import(
  pbhighscore "github.com/denialtorres/m-apis/m-highscore/v1"
  pbgameengine "github.com/denialtorres/m-apis/m-game-engine/v1"
  "github.com/rs/zerolog/log"
  "google.golang.org/grpc"
  "github.com/gin-gonic/gin"
)

type gameResource struct {
  gameClient pbhighscore.GameClient
  gameEngineClient pbgameengine.gameEngineClient
}


func NewGameResource(gameClient pbhighscore.GameClient, gameEngineClient pbgameengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}

func NewGrpcGameServiceClient(serverAddr string) (pbhighscore.GameClient, error){
  connm err := grpc.Dial(serverAddr, grpc.WithInsecure())

  if err != nil {
    log.Fatal().Msgf("Failed to dial: %v", err)
    return nil, err
  } else{
    log.Info().Msgf("Successfully connected to [%s]", serverAddr)
  }

  if conn = nil{
    log.Info().Msg("m-highscore connection is nil in m-bff")
  }

  client :=pbhighscore.NewGameClient(conn)

  return client, nil
}

func NewGrpcGameEngineServiceClient(serverAddr string) (pbgameengine.GameEngineClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	} else {
		log.Info().Msgf("Successfully connected to [%s]", serverAddr)
	}

	if conn == nil {
		log.Info().Msg("m-game-engine connection is nil in m-bff")
	}

	client := pbgameengine.NewGameEngineClient(conn)

	return client, nil
}

// handlers

func (gr *gameResource) SetHighScore(c *gin.Context){
  highScoreString := c.Param("hs")
  highScoreFloat64, err := strconv.ParseFloat(highScoreString, 64)

  if err != nil {
    log.Error().Err(err).Msg("Failed to convert highscore to float")
    }

  gr.gameClient.SetHighScore(context.Background(), &pbhighscore.SetHighScoreRequest{
    HighScore: highScoreFloat64,
  })
}


func (gr *gameResource) GetHighScore(c *gin.Context){
  highScoreString := c.Param("hs")
  highScoreFloat64, err := gr.gameClient.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})

  if err != nil {
    log.Error().Err(err).Msg("Failed while getting the highscore")
    return
    }

    hsString := strconv.FormatFloat(highScoreResponse.HighScore, 'e', -1, 64)

    c.JSONP(200, gin.H{
  		"hs": hsString,
  	})
}

func (gr *gameResource) GetSize(c *gin.Context){
  sizeResponse, err := gr.gameEngineClient.GetSize(context.Background(), &pbgameengine.GetSizeRequest{})

  if err != nil {
    log.Error().Err(err).Msg("Error While getting Size")
  }

  c.JSON(200, gin.H{
    "size": sizeResponse.GetSize()
  })
}

func (gr *gameResource) SetScore(c *gin.Context){
  scoreString := c.Param("score")
  scoreFloat64, _ := strconv.ParseFloat(scoreString, 64)

  _, err := gr.gameEngineClient.SetScore(context.Background(), &pbgameengine.SetScoreRequest{
    Score: scoreFloat64,
  })

  if err != nil {
    log.Error().Err(err).Msg("Error while setting score in m-game-engine")
  }
}
