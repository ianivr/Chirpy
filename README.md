# Chirpy

Chirpy es un backend HTTP en Go para una mini red social. El proyecto permite registrar usuarios, autenticarse con JWT, crear y consultar chirps, renovar y revocar refresh tokens, y manejar un webhook de upgrade de usuario.

## Tecnologías utilizadas

- Go
- PostgreSQL
- net/http
- sqlc
- JWT
- Argon2id
- godotenv
- UUID

## Estructura del proyecto

```text
.
├── main.go
├── index.html
├── requests.http
├── go.mod
├── go.sum
├── sql/
│   ├── queries/
│   │   ├── chirps.sql
│   │   ├── refresh_tokens.sql
│   │   └── users.sql
│   └── schema/
│       ├── 001_users.sql
│       ├── 002_chirps.sql
│       ├── 003_users.sql
│       ├── 004_refresh_tokens.sql
│       └── 005_users.sql
└── internal/
    ├── auth/
    ├── database/
    └── server/
```

## Descripción de carpetas

- `main.go`: punto de entrada de la aplicación.
- `internal/server`: contiene el router HTTP, handlers, middleware y configuración del servidor.
- `internal/auth`: lógica relacionada con JWT, bearer tokens, refresh tokens, hashing de contraseñas y validaciones.
- `internal/database`: código generado por sqlc para interactuar con PostgreSQL.
- `sql/schema`: migraciones SQL para crear y modificar tablas.
- `sql/queries`: consultas SQL usadas por sqlc.
- `requests.http`: ejemplos de peticiones HTTP para probar la API.

## Variables de entorno

La aplicación espera estas variables:

```env
DB_URL=postgres://usuario:password@localhost:5432/chirpy?sslmode=disable
PLATFORM=dev
JWT_SECRET=tu_secreto
POLKA_KEY=tu_clave_polka
```

## Modelo de datos

### Usuarios

- `id`: UUID
- `created_at`: timestamp
- `updated_at`: timestamp
- `email`: texto único
- `hashed_password`: contraseña hasheada
- `is_chirpy_red`: booleano

### Chirps

- `id`: UUID
- `created_at`: timestamp
- `updated_at`: timestamp
- `body`: texto
- `user_id`: UUID del autor

### Refresh tokens

- `token`: texto
- `created_at`: timestamp
- `updated_at`: timestamp
- `user_id`: UUID
- `expires_at`: timestamp
- `revoked_at`: timestamp nullable

## Endpoints principales

### Salud y administración

- `GET /api/healthz`
- `GET /admin/metrics`
- `POST /admin/reset`

### Usuarios

- `POST /api/users`
- `POST /api/login`
- `PUT /api/users`

### Chirps

- `GET /api/chirps`
- `GET /api/chirps/{chirpID}`
- `POST /api/chirps`
- `DELETE /api/chirps/{chirpID}`

### Tokens

- `POST /api/refresh`
- `POST /api/revoke`

### Webhook

- `POST /api/polka/webhooks`

## Flujo de autenticación

1. El usuario crea una cuenta con `POST /api/users`.
2. Luego inicia sesión con `POST /api/login`.
3. El servidor devuelve un JWT de acceso y un refresh token.
4. Los endpoints protegidos usan el header:

   ```http
   Authorization: Bearer <token>
   ```

5. El token de refresco se puede usar en `POST /api/refresh` para obtener un nuevo token de acceso.
6. `POST /api/revoke` revoca el refresh token.

## Cómo correr el proyecto

1. Tener PostgreSQL en ejecución.
2. Crear la base de datos.
3. Configurar las variables de entorno.
4. Ejecutar las migraciones SQL ubicadas en `sql/schema`.
5. Ejecutar el proyecto:

```bash
go run .
```

## Notas adicionales

- El endpoint `POST /admin/reset` solo funciona cuando `PLATFORM=dev`.
- Los chirps pasan por una función de censura de palabras al responderse.
- El middleware de métricas cuenta cuántas veces se sirvió contenido estático desde `/app/`.
