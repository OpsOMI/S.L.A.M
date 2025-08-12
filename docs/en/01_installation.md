## 📦 Installation with Docker

This project can be easily run using **Docker** and **Docker Compose**. Before getting started, make sure the following are installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- **Make** (optional) — Used for running commands more easily.
  If you don’t have Make on your system, you can also start the server using `docker compose` commands directly.

### 1️⃣ Clone the Project

```bash
git clone https://github.com/OpsOMI/S.L.A.M
cd S.L.A.M
```

### 2️⃣ Create Certificates

The project uses your own certificates for secure communication with TLS.

1. Copy the `cert.example.conf` file in `/certs/example` folder and save it as `cert.conf`:

```bash
cp certs/example/cert.example.conf certs/example/cert.conf
```

2. Edit the `DNS` field under the **alt_names** section inside `cert.conf` to match your server address.
   This field is critical for TLS verification.

3. Generate the certificates:

```bash
openssl req -x509 -nodes -days 365 \
  -newkey rsa:2048 \
  -keyout server.key \
  -out server.crt \
  -config cert.conf
```

### 3️⃣ Set Environment Variables (ENV)

Fill in the variables located in the `env/` folder:

```env
JWT_ISSUER=slam                # Identity for JWT package
JWT_SECRET=                    # Secret key for JWT
TSL_SERVER_NAME=               # DNS name from cert.conf
PRIVATE_ROOM_PASS=             # Password for the "private" room

MANAGEMENT_NICKNAME=           # Display name for admin
MANAGEMENT_USERNAME=           # Admin username
MANAGEMENT_PASSWORD=           # Admin password
```

Notes:

- **JWT_ISSUER / JWT_SECRET** → Used for token verification.
- **TSL_SERVER_NAME** → Must be the same as the DNS name in the certificate.
- **PRIVATE_ROOM_PASS** → Password for the default private room.
- **MANAGEMENT\_\*** → Default system administrator credentials.

### 4️⃣ Configure Database Credentials

Check the `.env.example` files under `./deployment/dev` and `./deployment/prod` directories.
You only need to fill in the following fields:

```env
DEV_DB_USER=
DEV_DB_PASSWORD=

PROD_DB_USER=
PROD_DB_PASSWORD=
```

Do not modify other fields.

### 5️⃣ Start the Server

If you have **Make** installed, you can start the server in detached mode with the commands below:

- For development mode:

```bash
make dev-build-d
```

- For production mode:

```bash
make prod-build-d
```

If you don’t have Make, you can also run the server using the following `docker compose` commands:

- For development mode:

```bash
docker compose -f ./deployment/dev/docker-compose.yml up --build -d
```

- For production mode:

```bash
docker compose -f ./deployment/prod/docker-compose.yml up --build -d
```

### 6️⃣ Admin User and Client

After the server is started:

- An **"owner"** role admin user is automatically created using the **MANAGEMENT\_\*** credentials you provided.
- The **client binary** file for this user is generated under the `./clients` directory.
- You can connect to the server using this generated client and the admin credentials.
- Two rooms are automatically created for the admin:

  - One with the code **public**
  - Another with the code **private**
  - The **private** room’s password is the value you set for **PRIVATE_ROOM_PASS** in your `env` file.
