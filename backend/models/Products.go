package models

import (
	"encoding/xml"
)

type Products struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"Soap,attr"`
	Body    struct {
		Text               string `xml:",chardata"`
		ReadMultipleResult struct {
			Text               string `xml:",chardata"`
			Xmlns              string `xml:"xmlns,attr"`
			ReadMultipleResult struct {
				Text         string         `xml:",chardata"`
				PRODUCTOS_WS []PRODUCTOS_WS `xml:"PRODUCTOS_WS"`
			} `xml:"ReadMultiple_Result"`
		} `xml:"ReadMultiple_Result"`
	} `xml:"Body"`
}

type PRODUCTOS_WS struct {
	Text                string `xml:",chardata"`
	Key                 string `xml:"Key"`
	CodNav              string `xml:"N_x00BA__producto"`
	Nombre              string `xml:"Descripcion"`
	NombreBoton         string `xml:"Descripcion_2"`
	PrecioUnidad        string `xml:"Precio_venta"`
	UnidadesPorLote     string `xml:"Unidades_por_lote"`
	NombreTiket         string `xml:"Descripcion_Venta"`
	CombinadoId         string `xml:"Combinado_mezcla"`
	Mezcla              string `xml:"Mezcla"`
	ConsumibleVendible  string `xml:"Consumible_vendible"`
	CategoriaId         string `xml:"Clasificacion_TPV"`
	Vendible            string `xml:"Producto_venta_directa"`
	Comprable           string `xml:"Prod_compra_inventario"`
	CodAvalon           string `xml:"Cod_Avalon"`
	PrecioCompra        string `xml:"Precio_compra_ref_1"`
	NombreProveedor     string `xml:"Nombre_proveedor_1"`
	CodProveedor        string `xml:"N_x00BA__proveedor_1"`
	Escandallo          string `xml:"Cant_baja_escandallo"`
	SubclasificacionTPV string `xml:"Subclasificacion_TPV"`
}
