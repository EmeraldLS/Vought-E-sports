package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/vought-esport-attendance/config"
	"github.com/vought-esport-attendance/model"
)

func InitializeDbContent(c *gin.Context) {
	var data struct {
		TournamentName string `json:"tournament_name,omitempty" bson:"tournament_name,omitempty" validate:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	var validate = validator.New()
	if err := validate.Struct(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	attendance := RepresentDBData(data.TournamentName)
	if err := config.InitializeDbContent(attendance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "content created.",
	})
}

func GetAllTournament(c *gin.Context) {
	tournaments, err := config.GetAllTournament()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, tournaments)
}

func GetTournament(c *gin.Context) {
	id := c.Param("id")
	tournament, err := config.GetTournament(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, tournament)
}

func GetLobbyByDay(c *gin.Context) {
	id := c.Param("id")
	dayNumberString := c.Param("day_number")
	dayNumber, err := strconv.Atoi(dayNumberString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	dayLobbbies, err := config.GetAllLobbyInADay(id, dayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, dayLobbbies)
}

func GetLobbyByID(c *gin.Context) {
	id := c.Param("id")
	dayNumberString := c.Param("day_number")
	lobbyID := c.Param("lobby_id")
	dayNumber, err := strconv.Atoi(dayNumberString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	lobby, err := config.GetLobbyByID(id, lobbyID, dayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, lobby)
}

func GetPlayersInALobbby(c *gin.Context) {

	id := c.Param("id")
	dayNumberString := c.Param("day_number")
	lobbyID := c.Param("lobby_id")
	dayNumber, err := strconv.Atoi(dayNumberString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	players, err := config.GetPlayersInALobbby(id, lobbyID, dayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, players)

}

func GetAllPlayersInAday(c *gin.Context) {
	id := c.Param("id")
	dayNumberString := c.Param("day_number")
	dayNumber, err := strconv.Atoi(dayNumberString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	players, err := config.GetAllPlayersInAday(id, dayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, players)
}

func GetPlayerDetailsFromALobby(c *gin.Context) {

	id := c.Param("id")
	var playerDetails model.PlayerDetails
	if err := c.ShouldBindJSON(&playerDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(playerDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	player, err := config.GetAPlayerFromALobby(id, playerDetails.LobbyID, playerDetails.PlayerID, playerDetails.DayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, player)
}

// func GetLobbyByIndex(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := config.GetLobbyByIndex(id); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"response": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"response": "success",
// 	})
// }

// func InsertPlayerKills(c *gin.Context) {
// 	playerID := c.Param("player_id")
// 	user, err := config.GetUserByID(playerID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"response": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}

// 	player := model.Player{
// 		PlayerID: user.PlayerID,
// 		Name:     user.Name,
// 		Kills:    6,
// 	}
// }

func CreateLobby(c *gin.Context) {
	_id := c.Param("id")
	var lobby model.LobbyCreation
	if err := c.ShouldBindJSON(&lobby); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(lobby); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	lobby.Date = carbon.Now().ToDateString()
	lobby.LobbyID = uuid.NewString()

	allLobby, err := config.CreateLobby(_id, lobby)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, allLobby)
}

func AddPlayerKillsInALobby(c *gin.Context) {

	_id := c.Param("id")

	var playerCreation model.KillCount

	if err := c.ShouldBindJSON(&playerCreation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	var validate = validator.New()
	if err := validate.Struct(playerCreation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	players, err := config.InsertPlayerKillInALobby(_id, playerCreation.LobbyID, playerCreation, playerCreation.DayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, players)

}

func GetTotalPlayerKillsInADay(c *gin.Context) {
	id := c.Param("id")
	dayString := c.Param("day_number")
	day, _ := strconv.Atoi(dayString)
	playerId := c.Param("player_id")
	totalKills, err := config.GetTotalPlayerKillsInADay(id, playerId, day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	player, err := config.GetUserByID(playerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_kills": totalKills,
		"player_name": player.Name,
	})
}

func GetTotalPlayerKillsInWholeTournament(c *gin.Context) {
	id := c.Param("id")
	playerId := c.Param("player_id")
	totalKills, err := config.GetTotalPlayerKillsInWholeTournament(id, playerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	player, err := config.GetUserByID(playerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_kills": totalKills,
		"player_name": player.Name,
	})
}
