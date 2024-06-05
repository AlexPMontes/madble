# Etapa de construcción
FROM golang:1.18 AS builder

# Instala las dependencias del sistema
RUN apt update
RUN apt install -y libopenal-dev
RUN apt install -y libopus-dev

# Establece el directorio de trabajo
WORKDIR /app

# Copia solo los archivos necesarios para descargar las dependencias
COPY go.mod go.sum ./

# Descarga las dependencias de Go
RUN go mod download

# Copia todo el código fuente
COPY . .

# Construye la aplicación
RUN go build -o madble_client

# Etapa de producción
FROM golang:1.18

# Copia solo los binarios necesarios de la etapa de construcción
COPY --from=builder /app/madble_client /app/

# Establece el directorio de trabajo en la nueva etapad
WORKDIR /app

RUN apt update && apt install -y libopenal-dev

COPY asound.conf /etc/asound.conf



# Comando para iniciar la aplicación
CMD ["./madble_client", "-server", "192.168.1.3:64738", "-username", "mft-04", "-insecure"]
