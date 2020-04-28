package types

// Meta stores meta data of user.
type Meta struct {
	UserID             string `bson:"_id,omitempty,-"`
	AccountStatus      string `bson:"AccountStatus,omitempty,-"`
	AccountCreationUTC int64  `bson:"AccountCreationUTC,omitempty,-"`
}
