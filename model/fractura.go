package model

import (
	"time"
)

type Fractura struct {
	Id                      *int
	Id_sondaje              *int
	Id_corrida              *int
	Id_tipo_discontinuidad  *int
	Cod_tipo_discontinuidad *string
	Id_forma                *int
	Cod_forma               *string
	Id_tipo_relleno         *int
	Cod_tipo_relleno        *string
	Id_abertura             *int
	Cod_abertura            *string
	Id_rugosidad            *int
	Cod_rugosidad           *string
	Id_relleno              *int
	Cod_relleno             *string
	Id_meteorizacion        *int
	Cod_meteorizacion       *string
	Angulo_alfa             *string
	Angulo_beta             *float32
	Profundidad             *string
	Id_jr                   *int
	Id_ja                   *int
	Comentario              *string
	Id_tipo_relleno2        *int
	Angulo_gama             *float32
	F_uuid                  *string
	Load_Date               *time.Time
	Loaded_By               *string
	Modified_Date           *time.Time
	Modified_By             *string
	Ts                      *string
}
