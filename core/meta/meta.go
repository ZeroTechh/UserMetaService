package meta

import (
	"context"
	"time"

	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/ZeroTechh/hades"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZeroTechh/UserMetaService/core/types"
)

var (
	config   = hades.GetConfig("main.yaml", []string{"config", "../config", "../../config"})
	dbConfig = config.Map("database")

	log = logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)
)

// New returns a new meta handler struct.
func New() *Meta {
	m := Meta{}
	m.init()
	return &m
}

// Meta handles meta data of user.
type Meta struct {
	coll *mongo.Collection
}

// init initializes.
func (m *Meta) init() {
	c := utils.CreateMongoDB(dbConfig.Str("address"), log)
	db := c.Database(dbConfig.Str("db"))
	m.coll = db.Collection(dbConfig.Str("collection"))
}

// Create generates and adds new meta data into db.
func (m Meta) Create(ctx context.Context, userID string) error {
	_, err := m.coll.InsertOne(
		ctx,
		types.Meta{
			UserID:             userID,
			AccountStatus:      config.Map("accountStatuses").Str("unverified"),
			AccountCreationUTC: time.Now().Unix(),
		},
	)
	return errors.Wrap(err, "Error while adding data into db")
}

// Get returns meta data of user.
func (m Meta) Get(ctx context.Context, userID string) (types.Meta, error) {
	var data types.Meta
	err := m.coll.FindOne(ctx, types.Meta{UserID: userID}).Decode(&data)
	return data, errors.Wrap(err, "Error while finding from db")
}

// SetStatus sets account status to something.
func (m Meta) SetStatus(ctx context.Context, userID, status string) error {
	_, err := m.coll.UpdateOne(
		ctx,
		types.Meta{UserID: userID},
		map[string]types.Meta{"$set": types.Meta{AccountStatus: status}},
	)
	return errors.Wrap(err, "Error while updating data in db")
}
