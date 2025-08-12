# Özellikler

S.L.A.M (Secure Link Anonymous Messaging) projesi, gizlilik ve güvenlik odaklı anonim mesajlaşma ihtiyacına yönelik olarak geliştirilmiştir. Aşağıda projenin öne çıkan özellikleri listelenmiştir:

## Genel Özellikler

- **TCP tabanlı güvenli iletişim:**
  Uzak bir sunucu ile TCP üzerinden bağlantı kurulur. İletişim TLS (Transport Layer Security) ile şifrelenir.

- **Kullanıcıya özel client binary:**
  Her kullanıcı için benzersiz, o kullanıcıya özel çalışan bir client executable (çalıştırılabilir dosya) otomatik olarak oluşturulur.

- **Client-User eşleştirme:**
  Client dosyası sadece ilgili kullanıcıya ait olduğundan, başkalarının kullanıcı bilgileri olsa dahi client olmadan giriş yapılamaz.

- **USB ve taşınabilir cihaz desteği:**
  Client uygulaması USB gibi taşınabilir ortamlarda kolayca çalıştırılabilir, ekstra kurulum gerekmez.

- **Yönetici kontrollü client binary**
  Sadece yönetici yeni bir kullanıcı dolayısıyla client oluşturabilir.

- **Oda bazlı iletişim:**
  Kullanıcılar public ve private olmak üzere oda sistemi ile organize olur.
  Private odanın şifresi, sistem yöneticisi tarafından belirlenir.

- **24 saatlik mesaj saklama ve otomatik temizleme:**
  Mesajlar veritabanında 24 saat tutulur, sonrasında otomatik olarak silinir. Bu sayede iz bırakmadan iletişim sağlanır.

## Güvenlik Özellikleri

- **TLS ile uçtan uca şifreleme:**
  Ağdaki tüm iletişim TLS üzerinden şifrelenerek dinleme ve müdahaleye karşı korunur.

- **JWT tabanlı kimlik doğrulama:**
  Kullanıcı kimlik doğrulaması JWT tokenlarıyla güvence altına alınır.

- **Veritabanı mesaj şifreleme:**
  Mesajlar veritabanında şifrelenmiş olarak tutulur.

## Kullanım ve Yönetim

- **Otomatik yönetici ve client oluşturma:**
  Sunucu başlangıcında yönetici bilgilerine göre "owner" rolünde bir kullanıcı ve ona bağlı client otomatik yaratılır.

- **Yönetici için otomatik public ve private odalar:**
  Yönetici hesabı için iki oda otomatik oluşturulur; public ve private. Private odanın şifresi yönetici tarafından belirlenir.

- **Kolay kurulum ve dağıtım:**
  Docker ve Docker Compose ile kolayca kurulabilir, taşınabilir ortamlarda çalışabilir.
