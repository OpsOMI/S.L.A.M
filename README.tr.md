# S.L.A.M — Secure Link Anonymous Messaging

S.L.A.M, gizliliğe odaklı, anonim mesajlaşma için geliştirilmiş bir TCP tabanlı iletişim sistemidir. Sistem, uzak bir sunucuda çalışan özel bir TCP server ve bu sunucuya güvenli şekilde bağlanan istemcilerden oluşur.

Her kullanıcı için **tek ve özel bir istemci (client)** üretilir. Başkasının elinde hem kendi giriş bilgileri hem de başka bir istemci olsa bile sisteme erişemez, çünkü istemci-kullanıcı eşleştirmesi benzersizdir.

Yeni bir kullanıcı, sistem yöneticisi tarafından eklendiğinde, o kullanıcıya özel çalıştırılabilir bir istemci otomatik olarak oluşturulur.Bu istemci, sunucuya erişmenin tek yoludur.

S.L.A.M, **USB bellek** veya diğer taşınabilir cihazlardan çalışacak şekilde tasarlanmıştır ve özellikle yüksek güvenlik gerektiren ortamlarda kullanılabilir.

## Amaç

S.L.A.M'in amacı, kullanıcıların güvenli, anonim ve hızlı şekilde mesajlaşabilmesini sağlamaktır. Özellikle hassas operasyonlarda, iz bırakmadan ve merkezi kontrol ile iletişim sağlamak için tasarlanmıştır.

## Dökümantasyon

- [Kurulum](docs/tr/01_installation.md)
- [Özellikler](docs/tr/02_features.md)
- [Komutlar](docs/tr/03_commands.md)

## Destek Olun

Projemizi beğendiyseniz, ⭐️ **Star** atmayı unutmayın!. Bu, bizi motive eder ve projenin daha çok kişiye ulaşmasına yardımcı olur.  
Teşekkürler! 🙌

## Sorumluluk Reddi

Bu proje “olduğu gibi” sunulmaktadır ve herhangi bir garanti içermez. Yazarlar, yazılımın üçüncü kişiler tarafından kötüye kullanılması veya zararlı amaçlarla kullanılmasından sorumlu değildir. Yazılımı sorumluluk bilinciyle ve kendi riskinizde kullanınız.

## Platform Uyumluluğu

- **Sunucu:** Docker kullanıldığı için herhangi bir işletim sisteminde çalışabilir ✅
- **İstemci:** Linux ✅, macOS ✅, Windows ❌ — İstemci şu anda **Windows'ta desteklenmiyor**

## Diller

[İngilizce](README.md)
