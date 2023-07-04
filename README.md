Docker üzerinde 1 adet MongoDB container’ı ayağa kaldırılmıştır ve bu container 3 Replica Set olarak High Available olarak yapılandırılmıştır.. MongoDB cluster yapılandırılmasından sonra bir init script ile stajdb adında bir database oluşuturulacak ve authentication aktif edilmiştir. Staj database’inde her 2 internet sitesi içinde iller ve ülkeler adında iki collection oluşturulmuştur ve 10’ar veri girilmiştir, bu tanımlar uygulamaların temel kodlarında yazılmıştır.
Python Flask kullanılarak bir web uygulaması geliştirilicek docker imajı oluşturulmuştur bu uygulama 5555 portunda çalışır durumdadır /path’ine istek atıldığında “Merhaba Python!” yanıtı döndürmektedir /staj path’ine istek atıldığında MongoDB’de oluşturulan staj database’indeki iller collection’ından rastgele bir il verisi döndürmektedir.
Go Gin kullanılarak  bir web uygulaması geliştirilmiştir, docker container’i oluşturulup bu container 4444 portunda çalışmaktadır /path’ine istek atıldığında “Merhaba Go!” yanıtı döndürmektedir /staj path’ine istek atıldığında MongoDB’de oluşturulan staj database’indeki iller collection’ından rastgele bir ülke verisi döndürmektedir.
Uygulama yazılım aşaması tamamlandıktan sonra Go uygulamasına ait metrikleri toplamak için Prometheus tercih edilmiştir. Prometheus aynı şekilde docker üzerinde bir konteynerda koşmaktadır. Prometheus’un metrikleri toplayaması açısından, Docker network teknolojisi kullanılmıştır. Docker’da net1 ve net2 adında 2 network ağı oluşturulmuştur. Net2 ağında uygulamaların test ve stage’leri koşturulmuştur ardından oluşturulan pipeline kapatılmıştır.

![](https://hackmd.io/_uploads/B1dm82WFh.png)


net1 ve net2 arasında bir köprü bulunmamaktadır. Net1 ortamında aynı metriklere sahip olması için Go web uygulaması, Prometheus, Grafana ve MongoDB koşmaktadır. 

![](https://hackmd.io/_uploads/HJQUInZY3.png)

Prometheus Go uygulamasının doğru bir şekilde çalışması için özel olarak yapılandırılmıştır. Var olan prometheus uygulaması, mertprometheus olarak set edilmiştir, ardından bu set edilen mertprometheus versiyon atlayarak mertprom2 olarak deploy edilmiştir. Aynı zamanda Go uygulamasının localhostta çalışan ismi değiştirilmiş olunup gouygulamasi olarak tekrar canlıya alınmıştır.

![](https://hackmd.io/_uploads/ryrY83ZFn.png)

Gerekli config düzenlemeleri yapıldıktan sonra Prometheus’ta gouygulamasi:4444 tanımı yapılmıştır. Save&Test sürecine sokularak gerekli test stage’leri atlandıktan sonra metrikler okunur hale gelmiştir.

![](https://hackmd.io/_uploads/rkxp82WYn.png)


Prometheus ana ekranında Go Uygulamamızda bir problem oluştuğunda ya da ayakta olduğunda görüntüleyebiliyoruz.

![](https://hackmd.io/_uploads/BJxkwnZKh.png)

Burada ilk başta metrikler okunurken hatalar bize gösterilmektedir. Gerekli düzenlemeler yapıldıktan sonra podumuz tekrar ayaktadır. Sorunsuz halde çalıştığını bizlere göstermektedir.

![](https://hackmd.io/_uploads/rJxew2ZK2.png)




Metriklerin okunma işlemlerinin ardından Grafana Dashboard üzerinde Test Data’ları aktif edilmiştir. Artık uygulamalarımız izlenir hale gelmiştir.
