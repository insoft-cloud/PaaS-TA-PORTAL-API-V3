package relationships

type Relationship struct {
	Guid string
}

type ToOneRelationship struct {
	Data Relationship
}

type ToManyRelationship struct {
	Data []Relationship
}
