# Felhasznált technológiák

## Szerver
### Go programozási nyelv
A Google által fejlesztett, 2009-ben megjelent programozási nyelv iránt, hatékonysága, open-source jellege és C közeli szintaxisa miatt az utóbbi időben jelentősen nőtt az érdeklődések száma.
![Go](https://jaxenter.com/wp-content/uploads/2018/04/HNranking.jpg)
A Go kihasználja a statikusan típusos nyelvek előnyeit, és nyelvi szinten tartalmaz szálkezelő és szinkronizációs megoldásokat. Ezáltal célszerű választás többszálú és nagy hatékonyságot igénylő hálózati alkalmazások készítése esetén. A Go nem objektum orientált nyelv, nem támogatja az osztályok definiálását, helyette interfészekkel és azok viselkedését megadó extension metódusokkal érhetünk el osztályszerű, egységbezárt viselkedést.

### Docker
A Docker egy konténerizációs technológia, mely megkönnyíti a mutiplatform alkalmazások készítését. A Docker segítségével az alkalmazásunkat és annak függőségeit egy konténerbe csomagoljuk. Ezáltal elegendő az alkalmazásunkat egy adott konténerben való futásra felkészítenünk, a konténer futtatását a különböző platformokon a keretrendszerre bízhatjuk.

### MongoDB
A MongoDB egy általános célú, dokumentum alapú adatbázikezelő rendszer. A hagyományos SQL adatbázismotorokhoz képest a MongoDB előnye, hogy segítségévela magas szintű programozási nyelvekből megszokott objektumok különösebb transzformációk vagy több táblára való bontás nélkül JSON formában tárolhatók. A MongoDB saját, szintén a JSON formátumon alapuló lekérdezőnyelvet használ az adatok kinyeréséhez és feldolgozásához. Számtalan keretrendszerhez és programozási nyelvhez létezik osztálykönyvtár a Mongo adatbázisok eléréséhez, többek között .Net, Java és Go esetében is.

## Kliens
### .Net Core
A Microsoft által 2016-ban bemutatott .NET Core egy ingyenesen hozzáférhető és nyílt forráskódú, általános célú szoftverfejlesztési keretrendszer. A teljes .NET Frameworkhöz képest jelentősen átdolgozott .NET Core multiplatform szoftverfejlesztést tesz lehetővé, egyaránt léteznek implementáció Windows, Linux és macOS operációs rendszerekre. A .NET nyelvfüggetlen, a C#, VB és F# mellett több mint 30 különböző nyelvet támogat. A különböző nyelvekről különböző platformokra történő fordítás egyszerűsítéséhez a .NET Framework és a .NET Core esetében is két lépésben történik a kód fordítása. A fordítás során először egy köztes IL (Common Intermediate Language) kód keletkezik, majd a program végrehajtásakor Just-in-time módon áll elő az IL kódból a gépi kód. A .NET Frameworkkel ellentétben a .NET Core moduláris felépítésű, így csak az alapvető osztályokat tartalmazza, további komponensek Nuget csomagok formájában tölthetőek le. A .NET Core a legdinamikusabban fejlődő .NET platform, legújabb, 3.0-s verziója, amely már Windows Forms és WPF alkalmazásokat is támogat, 2019-ben jelent meg. A moduláris felépítés több előnnyel is jár. Lerövidül a fordítási idő, a kész alkalmazások pedig gyorsabban indulhatnak és kevesebb tárhelyet igényelnek.

### Refit
A Refit egy .Net platformra (Core és Xamarin is támogatott) készült REST keretrendszer, amellyel deklaratív formában interfészek definiálásával készíthetünk REST kéréseket végrehajó programokat. A Refit használata során a szerveren támogatott kérésekhez egy interfészt adunk meg. Az interfész egy függvénye egy elérési utat és Http műveletet fog reprezentálni. Ezeket az információkat attribútumok segítségével adhatjuk meg. A programunkban ezután dinamikusan kérünk el egy, az általunk definiált interfészt megvalósító objektum példányt, melynek implementációját a keretrendszer generálja. A Refit teljesen típusbiztos megoldást nyújt REST kérések végrehajtásához, és a modern async-await minta követését is támogatja.