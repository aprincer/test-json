package response

import (
	"time"
)

type FracturaResponse struct {
	Id                      *int       `json:"id"`
	Id_sondaje              *int       `json:"id_sondaje"`
	Id_corrida              *int       `json:"id_corrida"`
	Id_tipo_discontinuidad  *int       `json:"id_tipo_discontinuidad"`
	Cod_tipo_discontinuidad *string    `json:"cod_tipo_discontinuidad"`
	Id_forma                *int       `json:"id_forma"`
	Cod_forma               *string    `json:"cod_forma"`
	Id_tipo_relleno         *int       `json:"id_tipo_relleno"`
	Cod_tipo_relleno        *string    `json:"cod_tipo_relleno"`
	Id_abertura             *int       `json:"id_abertura"`
	Cod_abertura            *string    `json:"cod_abertura"`
	Id_rugosidad            *int       `json:"id_rugosidad"`
	Cod_rugosidad           *string    `json:"cod_rugosidad"`
	Id_relleno              *int       `json:"id_relleno"`
	Cod_relleno             *string    `json:"cod_relleno"`
	Id_meteorizacion        *int       `json:"id_meteorizacion"`
	Cod_meteorizacion       *string    `json:"cod_meteorizacion"`
	Angulo_alfa             *string    `json:"angulo_alfa"`
	Angulo_beta             *float32   `json:"angulo_beta"`
	Profundidad             *string    `json:"profundidad"`
	Id_jr                   *int       `json:"id_jr"`
	Id_ja                   *int       `json:"id_ja"`
	Comentario              *string    `json:"comentario"`
	Id_tipo_relleno2        *int       `json:"Id_tipo_relleno2"`
	Angulo_gama             *float32   `json:"angulo_gama"`
	F_uuid                  *string    `json:"f_uuid"`
	Load_Date               *time.Time `json:"load_Date"`
	Loaded_By               *string    `json:"loaded_By"`
	Modified_Date           *time.Time `json:"modified_Date"`
	Modified_By             *string    `json:"modified_By"`
	Ts                      *string    `json:"ts"`
}
