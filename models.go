package afcclient

type User struct {
	ID         string `json:"auth_source_uuid,omitempty"`
	SourceName string `json:"auth_source_name,omitempty"`
	DistName   string `json:"distinguished_name,omitempty"`
	Role       string `json:"role,omitempty"`
	TokenLife  string `json:"token_lifetime,omitempty"`
	UserName   string `json:"username,omitempty"`
	UUID       string `json:"uuid,omitempty"`
}

type UpdateResult struct {
	MatchedCount  int `json:"matchedCount,omitempty"`
	ModifiedCount int `json:"modifiedCount,omitempty"`
	UpsertedCount int `json:"upsertedCount,omitempty"`
}

type DeleteResult struct {
	DeletedCount int `json:"deletedCount,omitempty"`
}

type InsertedResult struct {
	InsertedID string `json:"insertedID,omitempty"`
}
