package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Pagination struct {
	Total int
	Start int
	End   int
	Sql   string
}

// MODELS

type Room struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IsDirect bool   `json:"is_direct"`
}

func (Room) TableName() string {
	return "tbl_room"
}

type Participant struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (Participant) TableName() string {
	return "tbl_participant"
}

type Message struct {
	Text    string `json:"text"`
	Created string `json:"created"`
	Author  int    `json:"author" gorm:"default:null"`
	Room    int    `json:"room" gorm:"default:null"`
	Id      int    `json:"id"`
}

func (Message) TableName() string {
	return "tbl_message"
}

type RoomHasParticipants struct {
	Id          int `json:"id"`
	Room        int `json:"room" gorm:"default:null"`
	Participant int `json:"participant" gorm:"default:null"`
}

func (RoomHasParticipants) TableName() string {
	return "tbl_room_has_participants"
}

// Relations

// Value objects

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbuser := "root"
	dbpass := "root"
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	serverport := os.Getenv("SERVER_PORT")

	dbstr := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbport, dbname)

	db, err := gorm.Open("mysql", dbstr)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	r := gin.Default()
	r.Use(SetDBtoContext(db))
	r.Use(cors.Default())

	r.GET("/room", ListRoom)
	r.POST("/room", CreateRoom)
	r.GET("/room/:id", ReadRoom)
	r.PUT("/room/:id", UpdateRoom)
	r.DELETE("/room/:id", DeleteRoom)

	r.GET("/participant", ListParticipant)
	r.POST("/participant", CreateParticipant)
	r.GET("/participant/:id", ReadParticipant)
	r.PUT("/participant/:id", UpdateParticipant)
	r.DELETE("/participant/:id", DeleteParticipant)

	r.GET("/message", ListMessage)
	r.POST("/message", CreateMessage)
	r.GET("/message/:id", ReadMessage)
	r.PUT("/message/:id", UpdateMessage)
	r.DELETE("/message/:id", DeleteMessage)

	r.GET("/room-has-participants", ListRoomHasParticipants)
	r.POST("/room-has-participants", CreateRoomHasParticipants)
	r.GET("/room-has-participants/:id", ReadRoomHasParticipants)
	r.PUT("/room-has-participants/:id", UpdateRoomHasParticipants)
	r.DELETE("/room-has-participants/:id", DeleteRoomHasParticipants)

	r.Run(":" + serverport)
}

func Paginate(c *gin.Context) Pagination {
	start, _ := strconv.Atoi(c.Query("_start"))
	end, _ := strconv.Atoi(c.Query("_end"))
	var p Pagination
	p.Sql = fmt.Sprintf(" limit %d, %d", start, end-start)
	p.Start = start
	p.End = end

	return p
}

func ListRoom(c *gin.Context) {
	db := DBInstance(c)
	var listRoom []Room
	pagination := Paginate(c)
	query := fmt.Sprintf("%s %s ", "SELECT * FROM tbl_room", pagination.Sql)
	db.Raw(query).Scan(&listRoom)
	db.Raw("SELECT count(*) as total FROM tbl_room").Scan(&pagination)
	c.Header("X-Total-Count", fmt.Sprintf("%d", pagination.Total))
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, listRoom)
}

func CreateRoom(c *gin.Context) {
	db := DBInstance(c)
	var room Room

	if err := c.Bind(&room); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&room).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, room)
}

func ReadRoom(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var room Room
	if db.First(&room, id).Error != nil {
		content := gin.H{"error": "room with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, room)
}

func UpdateRoom(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var room Room
	if db.First(&room, id).Error != nil {
		content := gin.H{"error": "room with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if err := c.Bind(&room); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&room).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, room)

}

func DeleteRoom(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var room Room
	if db.First(&room, id).Error != nil {
		content := gin.H{"error": "room with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	if err := db.Delete(&room).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, "deleted")
}

func ListParticipant(c *gin.Context) {
	db := DBInstance(c)
	var listParticipant []Participant
	pagination := Paginate(c)
	query := fmt.Sprintf("%s %s ", "SELECT * FROM tbl_participant", pagination.Sql)
	db.Raw(query).Scan(&listParticipant)
	db.Raw("SELECT count(*) as total FROM tbl_participant").Scan(&pagination)
	c.Header("X-Total-Count", fmt.Sprintf("%d", pagination.Total))
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, listParticipant)
}

func CreateParticipant(c *gin.Context) {
	db := DBInstance(c)
	var participant Participant

	if err := c.Bind(&participant); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&participant).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, participant)
}

func ReadParticipant(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var participant Participant
	if db.First(&participant, id).Error != nil {
		content := gin.H{"error": "participant with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, participant)
}

func UpdateParticipant(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var participant Participant
	if db.First(&participant, id).Error != nil {
		content := gin.H{"error": "participant with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if err := c.Bind(&participant); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&participant).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, participant)

}

func DeleteParticipant(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var participant Participant
	if db.First(&participant, id).Error != nil {
		content := gin.H{"error": "participant with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	if err := db.Delete(&participant).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, "deleted")
}

func ListMessage(c *gin.Context) {
	db := DBInstance(c)
	var listMessage []Message
	pagination := Paginate(c)
	query := fmt.Sprintf("%s %s ", "SELECT * FROM tbl_message", pagination.Sql)
	db.Raw(query).Scan(&listMessage)
	db.Raw("SELECT count(*) as total FROM tbl_message").Scan(&pagination)
	c.Header("X-Total-Count", fmt.Sprintf("%d", pagination.Total))
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, listMessage)
}

func CreateMessage(c *gin.Context) {
	db := DBInstance(c)
	var message Message

	if err := c.Bind(&message); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&message).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, message)
}

func ReadMessage(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var message Message
	if db.First(&message, id).Error != nil {
		content := gin.H{"error": "message with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, message)
}

func UpdateMessage(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var message Message
	if db.First(&message, id).Error != nil {
		content := gin.H{"error": "message with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if err := c.Bind(&message); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&message).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, message)

}

func DeleteMessage(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var message Message
	if db.First(&message, id).Error != nil {
		content := gin.H{"error": "message with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	if err := db.Delete(&message).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, "deleted")
}

func ListRoomHasParticipants(c *gin.Context) {
	db := DBInstance(c)
	var listRoomHasParticipants []RoomHasParticipants
	pagination := Paginate(c)
	query := fmt.Sprintf("%s %s ", "SELECT * FROM tbl_room_has_participants", pagination.Sql)
	db.Raw(query).Scan(&listRoomHasParticipants)
	db.Raw("SELECT count(*) as total FROM tbl_room_has_participants").Scan(&pagination)
	c.Header("X-Total-Count", fmt.Sprintf("%d", pagination.Total))
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, listRoomHasParticipants)
}

func CreateRoomHasParticipants(c *gin.Context) {
	db := DBInstance(c)
	var room_has_participants RoomHasParticipants

	if err := c.Bind(&room_has_participants); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&room_has_participants).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, room_has_participants)
}

func ReadRoomHasParticipants(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var room_has_participants RoomHasParticipants
	if db.First(&room_has_participants, id).Error != nil {
		content := gin.H{"error": "room_has_participants with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, room_has_participants)
}

func UpdateRoomHasParticipants(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var room_has_participants RoomHasParticipants
	if db.First(&room_has_participants, id).Error != nil {
		content := gin.H{"error": "room_has_participants with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if err := c.Bind(&room_has_participants); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&room_has_participants).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, room_has_participants)

}

func DeleteRoomHasParticipants(c *gin.Context) {
	db := DBInstance(c)
	id := c.Params.ByName("id")
	var room_has_participants RoomHasParticipants
	if db.First(&room_has_participants, id).Error != nil {
		content := gin.H{"error": "room_has_participants with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	if err := db.Delete(&room_has_participants).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, "deleted")
}

func SetDBtoContext(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}
