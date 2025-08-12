# Komutlar

S.L.A.M istemcisinde kullanıcıların ve yöneticilerin kullanabileceği komutlar aşağıda listelenmiştir.

## Genel Komutlar (Tüm Kullanıcılar İçin)

| Komut                           | Açıklama                                                                                      | Örnek                      |
| ------------------------------- | --------------------------------------------------------------------------------------------- | -------------------------- |
| `/login`                        | Kullanıcı giriş yapmak için kullanılır.                                                       | `/login`                   |
| `/room/create [true/false]`     | Yeni oda oluşturur. `true` yazılırsa şifreli oda, `false` yazılırsa şifresiz oda oluşturulur. | `/room/create true`        |
| `/room/list`                    | Kullanıcının sahibi olduğu tüm odaları listeler.                                              | `/room/list`               |
| `/room/join [code] [password?]` | Belirtilen odaya katılır. Şifreli oda ise şifre girilir, şifresizse boş bırakılabilir.        | `/room/join publicroom123` |
| `/room/clean`                   | Kullanıcının sahibi olduğu odadaki tüm mesajları anlık olarak siler.                          | `/room/clean`              |
| `/me`                           | Kullanıcının kendi bilgilerini (nickname, username) gösterir.                                 | `/me`                      |
| `/reconnect`                    | Bağlantı koparsa sunucuya yeniden bağlanmayı dener.                                           | `/reconnect`               |
| `/clear`                        | Terminal ekranını temizler.                                                                   | `/clear`                   |
| `/logout`                       | Kullanıcıdan çıkış yapılır.                                                                   | `/logout`                  |
| `/exit`, `/quit`                | Uygulamadan tamamen çıkar.                                                                    | `/exit` veya `/quit`       |

💬 **Not:** Eğer bir odaya girdiyseniz, komut yazmadan **düz metin** yazarak odaya mesaj gönderebilirsiniz.

## Yönetici Komutları (Owner Yetkisi Gerektirir)

| Komut       | Açıklama                                                                                              | Örnek       |
| ----------- | ----------------------------------------------------------------------------------------------------- | ----------- |
| `/register` | Yeni bir kullanıcı oluşturur. İstemci binary dosyası otomatik oluşturulur (`client/[username]/main`). | `/register` |
| `/online`   | Sunucuya bağlı anlık online kullanıcıların sayısını listeler.                                         | `/online`   |

[← Geri](./02_features.md)
