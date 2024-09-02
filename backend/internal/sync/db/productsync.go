package db

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/victorlujan/tentacle/backend/internal/sync/utils"
	"github.com/victorlujan/tentacle/backend/models"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"log/slog"

	"github.com/jmoiron/sqlx"
)

func UpdateProducts(ctx context.Context, db *sqlx.DB, products models.Products, logger *logrus.Logger) error {

	slog.Info("Updating Products")

	var productsUpdated int = 0
	var productsInserted int = 0
	//ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, "ALTER TABLE productos AUTO_INCREMENT=1")
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT codNav, nombre FROM productos WHERE codNav IS NOT NULL ORDER BY id")
	if err != nil {
		return err
	}

	productsList := make(map[string]string)
	for rows.Next() {
		var codNav string
		var nombre string
		err = rows.Scan(&codNav, &nombre)
		if err != nil {
			return err
		}
		productsList[codNav] = nombre
	}

	rows, err = db.Query("SELECT id, codigo FROM categoriasProducto WHERE codigo IS NOT NULL ORDER BY id")
	if err != nil {
		return err
	}

	categoriasProductos := make(map[string]string)
	for rows.Next() {
		var id string
		var codigo string
		err = rows.Scan(&id, &codigo)
		if err != nil {
			return err
		}
		categoriasProductos[codigo] = id
	}
	for _, product := range products.Body.ReadMultipleResult.ReadMultipleResult.PRODUCTOS_WS {

		var categoriaProducto string
		var categoriaMezcla string
		var combinadoMezcla string
		var subclasificacionTPV string

		if product.CategoriaId != "" {
			categoriaProducto = categoriasProductos[product.CategoriaId]

		}
		if product.Mezcla != "" {
			categoriaMezcla = categoriasProductos[product.Mezcla]

		}

		if product.CombinadoId != "" {
			combinadoMezcla = productsList[product.CombinadoId]
		}

		if product.SubclasificacionTPV != "" && categoriaProducto != "" {
			var id string
			row := tx.QueryRowContext(ctx, "SELECT id FROM subCategoriasProducto WHERE codigo = ? AND categoria_id = ?", product.SubclasificacionTPV, categoriaProducto)
			err = row.Scan(&id)
			if err != nil {
				return err
			}
			subclasificacionTPV = id
		}

		//CHECK IF PRODUCT EXISTS
		_, productExists := productsList[product.CodNav]

		if productExists {
			_, err = tx.ExecContext(ctx, "UPDATE productos SET updated_on = ?, catTPV_id = ?, subcatTPV_id = ?,nombreTicket = ?, mezcla_id = ?, combinado_id = ?, nombre = ?, nombreBoton = ?, precioUnidad = ?, vendible = ?, comprable = ?, cantidadPorCaja = ?, codProveedor = ?, nombreProveedor = ?, precioCompra = ?, consumible_vendible = ?, escandallo = ? WHERE codNav = ?",
				time.Now(), utils.NewNullString(categoriaProducto), utils.NewNullString(subclasificacionTPV), product.NombreTiket, utils.NewNullString(categoriaMezcla), utils.NewNullString(combinadoMezcla), product.Nombre, product.NombreBoton, product.PrecioUnidad, product.Vendible, utils.IsComprable(product.Comprable), product.UnidadesPorLote, product.CodProveedor, product.NombreProveedor, product.PrecioCompra, utils.IsVendible(product.ConsumibleVendible), product.Escandallo, product.CodNav)

			if err != nil {
				return err
			}
			logger.Info("Updated product: ", product.CodNav)
			runtime.EventsEmit(ctx, "productUpdate", fmt.Sprintf("Product updated: %v , %v", product.CodNav, product.Nombre))

			productsUpdated++

		} else {
			_, err = tx.ExecContext(ctx, "INSERT INTO productos (created_on, updated_on,categoria_id, catTPV_id, subcatTPV_id, mezcla_id, combinado_id, nombre, nombreBoton,precioUnidad, vendible, comprable, cantidadPorCaja, codProveedor, nombreProveedor, precioCompra, consumible_vendible, escandallo, codNav, nombreTicket) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
				time.Now(), time.Now(), utils.NewNullString(categoriaProducto), utils.NewNullString(categoriaProducto), utils.NewNullString(subclasificacionTPV), utils.NewNullString(categoriaMezcla),
				utils.NewNullString(combinadoMezcla), product.Nombre, product.NombreBoton, product.PrecioUnidad, utils.IsVendible(product.Vendible), utils.IsComprable(product.Comprable),
				product.UnidadesPorLote, product.CodProveedor, product.NombreProveedor, product.PrecioCompra, utils.IsVendible(product.ConsumibleVendible), product.Escandallo, product.CodNav, product.NombreTiket)
			if err != nil {
				return err
			}
			logger.Info("Inserted product: ", product.CodNav)
			runtime.EventsEmit(ctx, "productUpdate", fmt.Sprintf("Product inserted: %v , %v", product.CodNav, product.Nombre))

			productsInserted++
		}
		progress := (float64(productsUpdated) / float64(len(products.Body.ReadMultipleResult.ReadMultipleResult.PRODUCTOS_WS)) * 100)
		runtime.EventsEmit(ctx, "progress", progress)
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}
