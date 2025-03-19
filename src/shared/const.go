package shared

import "strconv"

const (
	MainLayer   = "main_layer"
	AppName     = "challenge_prueba_biblioteca"
	HealthLayer = "health_layer"
	MongoLayer  = "Mongo_Layer"
)

const (
	DatabaseCollection string = "books_challenge"
	MutantCollection   string = "book_collection"
)

const (
	BadRequestMsg      string = "Solicitud inválida"
	ExistRequestMsg    string = "El ID del libro ya existe"
	OKRequestMsg       string = "Libro creado exitosamente"
	NotFoundRequestMsg string = "El ID del libro no existe"
)

const (
	BookAlreadyExists string = "el documento: %v ya existe"
)

func GetIntFromString(val string) (int, error) {
	rst, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return rst, nil
}
