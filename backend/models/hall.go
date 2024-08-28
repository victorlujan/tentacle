package models

import (
	"encoding/xml"
)

type Halls struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"Soap,attr"`
	Body    struct {
		Text               string `xml:",chardata"`
		ReadMultipleResult struct {
			Text               string `xml:",chardata"`
			Xmlns              string `xml:"xmlns,attr"`
			ReadMultipleResult struct {
				Text           string           `xml:",chardata"`
				CONFSALONES_WS []CONFSALONES_WS `xml:"CONFSALONES_WS"`
			} `xml:"ReadMultiple_Result"`
		} `xml:"ReadMultiple_Result"`
	} `xml:"Body"`
}

type CONFSALONES_WS struct {
	Text                   string `xml:",chardata"`
	Key                    string `xml:"Key"`
	CodigoSalon            string `xml:"Cod_salon"`
	Nombre                 string `xml:"Descripcion"`
	SinDNI                 string `xml:"Sin_DNI"`
	DireccionIP            string `xml:"Direccion_IP"`
	ActivoEmotivanet       string `xml:"Activo_EMOTIVA_NET"`
	ClubFumadores          string `xml:"Club_fumadores"`
	HoraApertura           string `xml:"Hora_apertura"`
	HoraCierre             string `xml:"Hora_cierre"`
	InvitacionesSinCliente string `xml:"Invitaciones_sin_cliente"`
	SalonSocio             string `xml:"Salon_socio"`
	ExcluirControlAcceso   string `xml:"Excluir_control_poblacion"`
	RegionCode             string `xml:"Siglas_comunidad"`
	IsHall                 string `xml:"Salon_Emotiva"`
	RutaMusica             string `xml:"Ruta_FTP_musica_salon"`
	Correoresponsable      string `xml:"Correo_responsable_sala"`
	CorreoCoordinador      string `xml:"Correo_coordinador_sala"`
	Localidad              string `xml:"Localidad_Salon"`
	Direccion              string `xml:"Direccion_salon"`
	Empresa                string `xml:"Empresa"`
	Cif                    string `xml:"_x003C_CIF_Empresa_x003E_"`
	Telefono               string `xml:"N_x00BA__telefono"`
	Externo                string `xml:"Externo"`
}
