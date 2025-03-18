package model

type Mutant struct {
	DNA      string `json:"dna" bson:"dna"`
	IsMutant bool   `json:"isMutant" bson:"isMutant"`
}
