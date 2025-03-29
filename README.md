Enlace la Repositorio: Backend    Fronted
Parte 1: Backend (Go) Instalación
Creacion
Se desarrolló una aplicación backend en Golang usando el framework Gin y el ORM GORM, para la creación de una API REST para el manejo de imágenes clasificadas.
Despliegue en AWS Lightsail
1.	Se creó una instancia Ubuntu en AWS Lightsail.
2.	Se instalaron Go y PostgreSQL en la instancia.
3.	Se subió el proyecto al servidor (vía git clone).
4.	Se compilaron y ejecutaron los servicios:
5.	Se utilizó screen para mantener el servicio activo en segundo plano.
 
![image](https://github.com/user-attachments/assets/632544ca-eb75-49c1-acf7-0de054ee792d)


Comprobación en Postman
Las imágenes eran subidas correctamente vía POST /imagenes.
 ![image](https://github.com/user-attachments/assets/8fce2980-0c20-464a-b9a6-e5e77efc1180)

La ruta GET /imagenes devolvía un JSON con los registros.
 ![image](https://github.com/user-attachments/assets/02ec40ac-7831-44b6-827f-917a86e8c334)

Las imágenes eran accesibles mediante la ruta pública:
Imagen de prueba
http://107.23.71.140:8080/uploads/img006.jpg


Peticiones a la instancia
![image](https://github.com/user-attachments/assets/fdc134dd-536f-4da3-aa1f-f8ffa93dc6fc)


 
Parte 2: Función Lambda con Clasificación de Imágenes
Se implementó una función Lambda en Node.js para recibir una imagen codificada en base64 desde el frontend, clasificarla por nombre y enviarla al backend desarrollado en Go para su almacenamiento.
El entorno de ejecución seleccionado fue: Node.js 18.x, utilizando bibliotecas como axios, form-data, fs y path.

![image](https://github.com/user-attachments/assets/c9be738a-28c3-44dc-9255-426602175a4e)

 
Lógica de clasificación
Se definió una función de clasificación simulada (classifyImage) que identifica si el archivo es una:
•	Factura (si contiene factura o invoice en el nombre)
•	Documento (si contiene document)
•	Foto (por defecto)
Integración con backend en Go
La imagen se guarda temporalmente en /tmp, y se utiliza form-data para enviarla al endpoint del backend: http://107.23.71.140:8080/imagenes
La solicitud HTTP incluye:
•	nombre: nombre del archivo
•	tipo_detectado: tipo de imagen detectado
•	descripcion: descripción generada
•	archivo: el archivo como multipart/form-data
Exposición mediante Function URL
La función Lambda fue publicada mediante Function URL pública, sin autenticación, para permitir que cualquier cliente pudiera enviar una imagen:
https://i2ux3umjlk4z2uxhfyfmucmsii0zebqe.lambda-url.us-east-2.on.aws/
Comprobación en Postman
La función fue probada usando Postman, comprobando el envío correcto de los datos al backend.
![image](https://github.com/user-attachments/assets/4316695f-b1c3-4308-b5e1-13ea722b40d8)

 
Parte 3: Integración
Desarrollo del frontend
Se desarrolló una interfaz web sencilla utilizando HTML, Bootstrap y JavaScript para:
•	Permitir al usuario subir una imagen desde el navegador
•	Enviar la imagen codificada en base64 a la función Lambda
•	Mostrar el resultado de la clasificación (tipo y descripción)
•	Mostrar el listado de imágenes almacenadas en la base de datos a través del backend
La interfaz se integró con:
•	La Lambda Function URL para enviar las imágenes
•	El backend en Go desplegado en AWS Lightsail para mostrar el historial de imágenes
Flujo de funcionamiento
•	El usuario selecciona una imagen desde la interfaz.
•	La imagen se convierte a base64 y se envía a la función Lambda.
•	Lambda clasifica la imagen y reenvía los datos al backend Go.
•	El backend guarda el archivo y los metadatos.
•	El frontend obtiene el listado de imágenes desde el backend y las muestra.


![image](https://github.com/user-attachments/assets/474fee79-499e-4a3c-b6dd-ca43f59605b0)
![image](https://github.com/user-attachments/assets/637d4c65-c009-4c63-8a72-9cfa39acb06e)
 

