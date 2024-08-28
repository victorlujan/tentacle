package models

import "encoding/xml"

type UserEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"Soap,attr"`
	Body    struct {
		Text               string `xml:",chardata"`
		ReadMultipleResult struct {
			Text               string `xml:",chardata"`
			Xmlns              string `xml:"xmlns,attr"`
			ReadMultipleResult struct {
				Text       string `xml:",chardata"`
				USUARIOSWS []struct {
					Text              string `xml:",chardata"`
					Key               string `xml:"Key"`
					Usuario           string `xml:"Usuario"`
					NombreUsuario     string `xml:"Nombre_usuario"`
					Clave             string `xml:"Clave"`
					TipoUsuario       string `xml:"Tipo_Usuario"`
					RolEmotivaNet     string `xml:"Rol_Emotiva_Net"`
					Rotaturnos        string `xml:"Rotaturnos"`
					NombreEmpleado    string `xml:"Nombre_empleado"`
					ApellidosEmpleado string `xml:"Apellidos_empleado"`
					DNIEmpleado       string `xml:"DNI_empleado"`
					Delegacion        string `xml:"Delegacion"`
					Inactivo          string `xml:"Inactivo"`
					CodEmpleado       string `xml:"Cod_empleado"`
				} `xml:"USUARIOS_WS"`
			} `xml:"ReadMultiple_Result"`
		} `xml:"ReadMultiple_Result"`
	} `xml:"Body"`
}
