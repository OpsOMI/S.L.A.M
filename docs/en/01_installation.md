## 📦 Docker Installation

This project can be easily run using **Docker** and **Docker Compose**. Before starting, make sure the following are installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- **Make** (optional) — Used to run commands more easily.
  If your system does not have Make, you can start the server with `docker compose` commands directly.

### 1️⃣ Clone the Project

```bash
git clone https://github.com/OpsOMI/S.L.A.M
cd S.L.A.M
```

### 2️⃣ Generate Certificates

The project uses TLS for secure communication. You need to generate your own certificates.

1. Copy the example certificate config:

```bash
cp certs/example/cert.example.conf certs/real/cert.conf
```

2. Edit the **DNS** field under **alt_names** in `cert.conf` to match your server address.
   This is critical for TLS validation.

3. Generate the certificates:

```bash
openssl req -x509 -nodes -days 365 \
  -newkey rsa:2048 \
  -keyout server.key \
  -out server.crt \
  -config ./certs/real/cert.conf
```

### 3️⃣ Server Config Settings

Before running, check the important settings in `/configs/server.yaml`:

```yaml
# Server configuration including host, port, and TLS certificate paths
server:
  external_host:
  host: 192.168.1.27
  port: 6666
  tls_cert_path: "./certs/real/server.crt"
  tls_key_path: "./certs/real/server.key"
```

- **external_host** → If you are running the server on another machine, enter the **public IP** of that server here.

  - If this field is not empty, the client will connect to `external_host`.
  - If empty, the client will connect to the `host` field.

- **host** → Used when running the server on your own computer.

  - If running on a separate server, provide the server’s local IP.
  - If on the same network, the client will connect to this IP.

- **port** → Port on which the server will listen.
- **tls_cert_path / tls_key_path** → Path to certificate and key files.

  - If you followed the installation instructions exactly, you don’t need to change these.

Other configuration options do not need to be modified.

### 4️⃣ Set Environment Variables (ENV)

Copy `env/example/.env.example` to `env/real/.env` and update the values:

```env
MESSAGE_SECRET=                # Strong 16/24/32 character secret for message encryption
JWT_ISSUER=slam                # Required for JWT token
JWT_SECRET=                    # JWT secret key
TSL_SERVER_NAME=               # Must match the DNS name in cert.conf
PRIVATE_ROOM_PASS=             # Password for the "private" room

MANAGEMENT_NICKNAME=           # Display name for admin
MANAGEMENT_USERNAME=           # Admin username
MANAGEMENT_PASSWORD=           # Admin password
```

Notes:

- **JWT_ISSUER / JWT_SECRET** → Used for token validation
- **TSL_SERVER_NAME** → Must match the DNS in your certificate
- **PRIVATE_ROOM_PASS** → Default password for the private room
- **MANAGEMENT\_\*** → Default admin account information

### 5️⃣ Database Settings

Check the `.env.example` files in `./deployment/dev` and `./deployment/prod`.
You only need to fill in:

```env
DEV_DB_USER=
DEV_DB_PASSWORD=

PROD_DB_USER=
PROD_DB_PASSWORD=
```

Do not modify other fields.

### 6️⃣ Start the Server

If **Make** is installed, start the server in the background:

- Development mode:

```bash
make dev-build-d
```

- Production mode:

```bash
make prod-build-d
```

**Without Make**, you can use `docker compose` directly:

- Development mode:

```bash
docker compose -f ./deployment/dev/docker-compose.yml up --build -d
```

- Production mode:

```bash
docker compose -f ./deployment/prod/docker-compose.yml up --build -d
```

### 7️⃣ Admin User and Client

After the server starts:

- An **"owner"** admin user is automatically created using the **MANAGEMENT\_\*** values.
- The corresponding **client binary** is generated in `./clients`.
- Use this client and the admin credentials to connect to the server.
- Two rooms are automatically created for the admin:

  - One with code **public**
  - One with code **private**
  - The private room password is the **PRIVATE_ROOM_PASS** from your `.env` file

[← Back](../../README.md)   |   [Next →](./02_features.md)
