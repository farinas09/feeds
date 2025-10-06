# Feeds API

Una ejemplo de una API distribuida de feeds construida con Go que implementa arquitectura de microservicios con comunicación basada en eventos.

## 🏗️ Arquitectura

Este proyecto implementa una arquitectura de microservicios con los siguientes componentes:

- **Feed Service**: Servicio para crear feeds
- **Query Service**: Servicio para consultar y buscar feeds
- **Pusher Service**: Servicio WebSocket para notificaciones en tiempo real
- **Nginx**: Proxy reverso para balancear la carga
- **PostgreSQL**: Base de datos principal
- **Elasticsearch**: Motor de búsqueda
- **NATS**: Sistema de mensajería para eventos

## 🚀 Características

- ✅ Creación de feeds con eventos
- ✅ Búsqueda de feeds usando Elasticsearch
- ✅ Notificaciones en tiempo real vía WebSocket
- ✅ Arquitectura de microservicios
- ✅ Comunicación asíncrona con NATS
- ✅ Dockerización completa
- ✅ Proxy reverso con Nginx

## 📋 Requisitos

- Docker y Docker Compose
- Go 1.24.3+ (para desarrollo local)

## 🛠️ Instalación y Uso

### Usando Docker Compose (Recomendado)

1. Clona el repositorio:
```bash
git clone <repository-url>
cd feeds-api
```

2. Ejecuta todos los servicios:
```bash
docker-compose up --build
```

3. La API estará disponible en `http://localhost:8080`

### Desarrollo Local

1. Instala las dependencias:
```bash
go mod download
```

2. Configura las variables de entorno:
```bash
export POSTGRES_DB=feeds
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
export NATS_ADDRESS=localhost:4222
export ELASTICSEARCH_URL=localhost:9200
```

## 📡 API Endpoints

### Feed Service (Puerto 8080)
- `POST /feeds` - Crear un nuevo feed

**Ejemplo de request:**
```json
{
  "title": "Mi Feed",
  "description": "Descripción del feed"
}
```

### Query Service (Puerto 8080)
- `GET /feeds` - Listar todos los feeds
- `GET /search?query=<término>` - Buscar feeds

### Pusher Service (Puerto 8080)
- `GET /ws` - Conexión WebSocket para notificaciones en tiempo real

## 🔧 Servicios

### Feed Service
- Maneja la creación de feeds
- Publica eventos cuando se crea un feed
- Almacena datos en PostgreSQL

### Query Service
- Proporciona endpoints para consultar feeds
- Implementa búsqueda usando Elasticsearch
- Escucha eventos de creación para indexar feeds

### Pusher Service
- Maneja conexiones WebSocket
- Notifica en tiempo real cuando se crean feeds
- Escucha eventos de NATS

## 🗄️ Base de Datos

### PostgreSQL
Almacena los feeds con la siguiente estructura:
```sql
CREATE TABLE feeds (
    id VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL
);
```

### Elasticsearch
Indexa los feeds para búsquedas rápidas con el siguiente mapeo:
```json
{
  "mappings": {
    "properties": {
      "id": {"type": "keyword"},
      "title": {"type": "text"},
      "description": {"type": "text"},
      "created_at": {"type": "date"}
    }
  }
}
```

## 🔄 Flujo de Eventos

1. Cliente envía POST a `/feeds`
2. Feed Service crea el feed en PostgreSQL
3. Feed Service publica evento `CreatedFeed` en NATS
4. Query Service recibe el evento e indexa el feed en Elasticsearch
5. Pusher Service recibe el evento y notifica a clientes WebSocket

## 🐳 Docker

El proyecto incluye:
- `Dockerfile` para construir las imágenes de los servicios
- `docker-compose.yml` para orquestar todos los servicios
- Configuración de Nginx para balancear carga

## 📦 Dependencias Principales

- **Gorilla Mux**: Router HTTP
- **NATS**: Sistema de mensajería
- **Elasticsearch**: Motor de búsqueda
- **PostgreSQL**: Base de datos
- **WebSocket**: Comunicación en tiempo real

## 🔍 Monitoreo y Logs

Los servicios incluyen logging estructurado para facilitar el debugging y monitoreo.

## 👨‍💻 Autor

Desarrollado por [farinas09](https://github.com/farinas09)
