package metaDB

import (
	"context"
	"time"

	"github.com/ZeroTechh/VelocityCore/utils"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZeroTechh/UserMetaService/core/types"
)

// New returns a new metaDB handler struct
func New() *MetaDB {
	metaDB := MetaDB{}
	metaDB.init()
	return &metaDB
}

// MetaDB is used to handle user extra data
type MetaDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// init initializes client and database
func (metaDB *MetaDB) init() {
	metaDB.client = utils.CreateMongoDB(dbConfig.Str("address"), log)
	metaDB.database = metaDB.client.Database(dbConfig.Str("db"))
	metaDB.collection = metaDB.database.Collection(dbConfig.Str("collection"))
}

// Create is used to generate and add new user meta data into database
func (metaDB MetaDB) Create(userID string) {
	data := types.Meta{
		UserID:             userID,
		AccountStatus:      config.Map("accountStatuses").Str("unverified"),
		AccountCreationUTC: time.Now().Unix(),
	}
	metaDB.collection.InsertOne(context.TODO(), data)
}

// Get is used to a users data
func (metaDB MetaDB) Get(userID string) (data types.Meta) {
	filter := types.Meta{UserID: userID}
	metaDB.collection.FindOne(context.TODO(), filter).Decode(&data)
	return
}

// ChangeStatus changed account status to something
func (metaDB MetaDB) ChangeStatus(userID, status string) {
	update := types.Meta{AccountStatus: status}
	metaDB.collection.UpdateOne(
		context.TODO(),
		types.Meta{UserID: userID},
		map[string]types.Meta{"$set": update},
	)
}
