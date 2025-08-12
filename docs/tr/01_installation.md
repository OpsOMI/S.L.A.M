## 📦 Docker ile Kurulum

Bu proje, **Docker** ve **Docker Compose** kullanılarak kolayca çalıştırılabilir. Başlamadan önce aşağıdakilerin yüklü olduğundan emin olun:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- **Make** (opsiyonel) — Komutları kolayca çalıştırmak için kullanılır.
  Eğer sisteminizde Make yoksa, sunucuyu `docker compose` komutlarıyla da başlatabilirsiniz.

### 1️⃣ Projeyi İndirin

```bash
git clone https://github.com/OpsOMI/S.L.A.M
cd S.L.A.M
```

### 2️⃣ Sertifika Oluşturma

Proje, TLS ile güvenli iletişim için kendi sertifikalarınızı kullanır.

1. `/certs/example` dizinindeki `cert.example.conf` dosyasını kopyalayıp `cert.conf` olarak kaydedin:

```bash
cp certs/example/cert.example.conf certs/real/cert.conf
```

2. `cert.conf` içindeki **alt_names** bölümündeki `DNS` alanını kendi sunucu adresinize göre düzenleyin.
   Bu alan TLS doğrulaması için kritik öneme sahiptir.

3. Sertifikaları oluşturun:

```bash
openssl req -x509 -nodes -days 365 \
  -newkey rsa:2048 \
  -keyout server.key \
  -out server.crt \
  -config cert.conf
```

### 3️⃣ Ortam Değişkenlerini (ENV) Ayarlama

`env/` klasöründe bulunan aşağıdaki değişkenleri doldurun:

```env
JWT_ISSUER=slam                # JWT paketi için gerekli kimlik
JWT_SECRET=                    # JWT'nin gizli anahtarı
TSL_SERVER_NAME=               # cert.conf içindeki DNS ismi
PRIVATE_ROOM_PASS=             # "private" odasının şifresi

MANAGEMENT_NICKNAME=           # Yönetici görünen adı
MANAGEMENT_USERNAME=           # Yönetici kullanıcı adı
MANAGEMENT_PASSWORD=           # Yönetici şifresi
```

Açıklamalar:

- **JWT_ISSUER / JWT_SECRET** → Token doğrulama için kullanılır.
- **TSL_SERVER_NAME** → Sertifikadaki DNS ismiyle aynı olmalı.
- **PRIVATE_ROOM_PASS** → Varsayılan özel odanın şifresi.
- **MANAGEMENT\_\*** → Varsayılan sistem yöneticisi bilgileri.

### 4️⃣ Veritabanı Bilgilerini Düzenleme

`./deployment/dev` ve `./deployment/prod` dizinlerindeki `.env.example` dosyalarına bakın.
Yalnızca aşağıdaki alanları doldurmanız yeterlidir:

```env
DEV_DB_USER=
DEV_DB_PASSWORD=

PROD_DB_USER=
PROD_DB_PASSWORD=
```

Diğer alanları değiştirmeyin.

### 5️⃣ Sunucuyu Başlatma

Eğer **Make** yüklüyse, aşağıdaki komutlarla arka planda sunucuyu başlatabilirsiniz:

- Geliştirme modunda:

```bash
make dev-build-d
```

- Production modunda:

```bash
make prod-build-d
```

**Make yoksa**, aşağıdaki `docker compose` komutlarıyla da çalıştırabilirsiniz:

- Geliştirme modu için:

```bash
docker compose -f ./deployment/dev/docker-compose.yml up --build -d
```

- Production modu için:

```bash
docker compose -f ./deployment/prod/docker-compose.yml up --build -d
```

### 6️⃣ Yönetici Kullanıcı ve Client

Sunucu başlatıldıktan sonra:

- Verdiğiniz **MANAGEMENT\_\*** bilgilerine göre otomatik olarak **"owner"** rolünde bir yönetici kullanıcı oluşturulur.
- Bu kullanıcıya ait **client binary** dosyası `./clients` dizininde oluşturulur.
- Oluşan client ile yönetici bilgileri kullanılarak sunucuya bağlanabilirsiniz.
- Yönetici için otomatik olarak **2 adet oda** oluşturulur:

  - Birinin kodu **public**
  - Diğerinin kodu **private**
  - **Private** odanın şifresi, sizin `env` dosyanızda belirttiğiniz **PRIVATE_ROOM_PASS** değeridir.

- [İleri](docs/tr/02_features.md)
