# Usar la imagen base oficial de Golang
FROM golang:latest as builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el código fuente del proyecto al contenedor
COPY . .

# Compilar la aplicación Go. Asegúrate de reemplazar 'main.go' con el camino y nombre correctos si es diferente
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Usar una imagen Docker 'scratch' para una imagen más pequeña y segura
FROM scratch

# Copiar el ejecutable compilado desde el builder al contenedor final
COPY --from=builder /app/myapp .

# Exponer el puerto 3000 en el contenedor
EXPOSE 3000

# Ejecutar la aplicación Go cuando el contenedor inicie
CMD ["./myapp"]