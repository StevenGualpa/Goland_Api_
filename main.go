package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Imagen struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Nombre        string `json:"nombre"`
	TipoDetectado string `json:"tipo_detectado"`
	Descripcion   string `json:"descripcion"`
	Archivo       string `json:"archivo"`
}

var db *gorm.DB

func main() {
	var err error

	// 💡 Configura esto con tu conexión real
	dsn := "host=localhost user=apiuser password=1234 dbname=imagenes port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}

	// Crear tabla automáticamente si no existe
	db.AutoMigrate(&Imagen{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Sirve archivos estáticos desde la carpeta 'uploads'
	r.Static("/uploads", "./uploads")
	// Endpoint para subir imagen
	r.POST("/imagenes", subirImagen)

	// Endpoint para listar imágenes
	r.GET("/imagenes", listarImagenes)

	// Ejecutar en el puerto 8080
	r.Run(":8080")
}

func subirImagen(c *gin.Context) {
	nombre := c.PostForm("nombre")
	tipo := c.PostForm("tipo_detectado")
	descripcion := c.PostForm("descripcion")

	file, err := c.FormFile("archivo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Archivo no proporcionado"})
		return
	}

	path := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar el archivo"})
		return
	}

	imagen := Imagen{
		Nombre:        nombre,
		TipoDetectado: tipo,
		Descripcion:   descripcion,
		Archivo:       path,
	}

	db.Create(&imagen)

	c.JSON(http.StatusOK, imagen)
}

func listarImagenes(c *gin.Context) {
	var imagenes []Imagen
	db.Find(&imagenes)
	c.JSON(http.StatusOK, imagenes)
}
