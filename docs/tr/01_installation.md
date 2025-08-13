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

1. `/certs/example` dizinindeki `cert.example.conf` dosyasını `/certs/real` içine kopyalayıp `cert.conf` olarak kaydedin:

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
  -config ./certs/real/cert.conf
```

### 3️⃣ Sunucu Config Ayarları

Kurulumdan önce `/configs/server.yaml` dosyasındaki önemli ayarları kontrol edin:

```yaml
# Server configuration including host, port, and TLS certificate paths
server:
  external_host:
  host: 192.168.1.27
  port: 6666
  tls_cert_path: "./certs/real/server.crt"
  tls_key_path: "./certs/real/server.key"
```

- **external_host** → Eğer server’i başka bir sunucuda açtıysanız, o sunucunun **public IP**’sini buraya yazın.

  - Eğer bu alan boş değilse, client **external_host**’a bağlanır.
  - Boş ise client, server’in `host` alanına bağlanır.

- **host** → Kendi bilgisayarınızda açtığınız server için kullanılır.

  - Sunucu başka bir makinede açıldıysa, local IP’nizi verin.
  - Aynı ağdaysanız client, bu IP’ye bağlanacaktır.

- **port** → Server’in hangi portta çalışacağını belirler.
- **tls_cert_path / tls_key_path** → Sertifika ve anahtar yolu.

  - Kurulum dokümanındaki adımları takip ettiyseniz değiştirmeye gerek yoktur.

Diğer config ayarlarına dokunmanıza gerek yoktur.

### 4️⃣ Ortam Değişkenlerini (ENV) Ayarlama

Aşağıdaki değişkenleri doldurmak için `env/example/.env.example` dosyasını `env/real/.env` olarak kopyalayın ve ardından değerlerini güncelleyin.

```env
MESSAGE_SECRET=                # Mesaj şifreleme için 16/24/32 karakter uzunluğunda güçlü gizli anahtar
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

### 5️⃣ Veritabanı Bilgilerini Düzenleme

`./deployment/dev` ve `./deployment/prod` dizinlerindeki `.env.example` dosyalarına bakın.
Yalnızca aşağıdaki alanları doldurmanız yeterlidir:

```env
DEV_DB_USER=
DEV_DB_PASSWORD=

PROD_DB_USER=
PROD_DB_PASSWORD=
```

Diğer alanları değiştirmeyin.

### 6️⃣ Sunucuyu Başlatma

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

### 7️⃣ Yönetici Kullanıcı ve Client

Sunucu başlatıldıktan sonra:

- Verdiğiniz **MANAGEMENT\_\*** bilgilerine göre otomatik olarak **"owner"** rolünde bir yönetici kullanıcı oluşturulur.
- Bu kullanıcıya ait **client binary** dosyası `./clients` dizininde oluşturulur.
- Oluşan client ile yönetici bilgileri kullanılarak sunucuya bağlanabilirsiniz.
- Yönetici için otomatik olarak **2 adet oda** oluşturulur:

  - Birinin kodu **public**
  - Diğerinin kodu **private**
  - **Private** odanın şifresi, sizin `env` dosyanızda belirttiğiniz **PRIVATE_ROOM_PASS** değeridir.

[← Geri](../../README.tr.md)   |   [İleri →](./02_features.md)
