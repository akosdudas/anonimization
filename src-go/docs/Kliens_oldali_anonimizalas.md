# Folyamatosan érkező adatok kliens oldali anonimizálása

## A folyamatos anonimizálás alapkoncepciója
A legtöbb ma létező anonimizálási megoldás a keletkező adatokat eredeti formájukban, anonimizálatlan formában gyűjti egy központi szerveren, ahol bizonyos idő elteltével vagy bizonyos mennyiségű adat összegyűjtése után végzik el az anonimizálást. Ennek legnagyobb veszélye, hogy az anonimizálás előtti időszakban a tárolt személyes adatok (az egyértelmű azonosítók és a hozzájuk tartozó szenzitív adatok) nem megfelelő védelem esetén illetéktelen kezekbe kerülhetnek. Erre a problémára nyújt megoldást a folyamatos anonimizálás, melynek során a szerverre beérkező adatok eredeti formájukban nem kerülnek tárolásra, hanem a szerver minden beérkezó adatot egyesével (vagy kis csoportokban) anonimizál, és az adatbázisba már csak az így anonimizált adatok kerülnek mentésre. Ezáltal még ha az adatbázishoz valaki hozzá is fér, akkor is csak az algoritmus által alkalmazott anonimizálási kritériumnak (pl.: k-anonimitás) megfelelő adatokat szerezhet meg.

## A kliens oldali anonimizálás előnyei
A fenti koncepció hátránya, hogy az adatok felküldésekor a szerver továbbra is hozzáfér a még anonimizálatlan adatokhoz, ami pedig maga után vonja azt is, hogy az anonimizálást végző félnek továbbra meg kell felelnie a
GDPR által előírt követelményeknek. Ezen kívül teljes bizalom szükséges a kliens és a szerver között, hiszen egy rosszindulatú szerver például titokban eltárolhatná a felküldött adatokhoz tartozó egyértelmű azonosítókat is. Ennek kiküszöbölésére érdemes megalkotni egy olyan algoritmust, amely lehetővé teszi a kliens számára (ahol az adat generálódik) azt, hogy a saját adatait már helyben anonimizálja, és csak ezután küldje fel őket a szerverre. Ezáltal a kliens teljes biztonságban tudhatja felküldött adatait, és a szerver sem lép interakcióba személyes adatokkal, így annak üzemeltetője mentesül a GDPR előírásai alól. A szerveren így eltárolt adatok később akár szabadon publikálhatóak, hiszen nem tartalmaznak személyes információt.

## Az alkalmazásban használt alapfogalmak
### Adathalmaz (dataset)
Az alkalmazás párhuzamosan több adathalmaz anonimizálását és tárolását támogatja. Egy adathalmazhoz tartozó adatokat egy dataset fogja össze. Adathalmaz szinten állíthatók be az anonimizálási paraméterek, valamint az adathalmaz létrehozásakor adhatjuk meg annak sémáját, azaz a benne található attribútumokat és azok típusait. Az adathalmaz sémája később új attribútumok felvételével tovább bővíthető.

### Attribútumok
Az alkalmazás különböző típusú mezőket (attribútumokat) kezel. Egy attribútum besorolása a megfelelő kategóriába tervezői döntés, melyet a szerver oldali adatbázis (a programban dataset) sémájának meghatározásakor kell megadni. Az alkalmazás a következő típusú attribútumokat különbözteti meg:
* **Egyértelmű azonosító:** Egy adott személy kizárólagos tulajdonsága; olyan adat, mely e személy kilétét egyértelműen meghatározza (pl. név, személyi igazolvány szám). Ezek az anonimzálás során eltávolításra kerülnek. A dataset létrehozásakor szerver oldali anonimizálás esetén az egyértelmű azonosítóknak **drop** módot adjunk meg, kliens oldali anonimizáláskor pedig nem küdjük fel ezeket az attribútumokat.
* **Kvázi azonosító:** Olyan adat, mely önmagában nem, más információkkal közösen viszont már képes egy adott személy egyértelmű azonosítására (pl.születési idő, irányítószám). Kvázi azonosítók esetén a rendszer különböző adattípusokat támogat. 
    * **Intervallum attribútum:** Jelenleg a szám típusú kvázi azonosítók intervallum attribútumként adhatóak meg, ekkor az adott mező beállításához a **int** módot használhatjuk. Az intervallum attribútumok az intervallum méretének csökkentésével finomíthatóak, ezáltal az adathalmaz információtartalma és hasznossága növelhető. Ilyen intervallum attribútum például az életkor, a fizetés vagy a súly. 
    * **Kategorikus attribútum:** A szöveges attribútumok megadása kategorikus attribútumként történik, melyek definiálásához a **cat** módot kell megadni. A kategorikus attribútumok (a jelenlegi implementációban) nem finomíthatóak, ezekre példa a lakóhely települése vagy a nem.
* **Szenzitív adat:** Egy adott személyről érzékeny információt hordozó ismertető (pl. betegség, politikai nézetek). Mivel az adathalmaz lényegét ezek a szenzitív adatok adják, így ezek eredeti formájukban kerülnek bele az anonimizált adatbázisba. Definiálásukhoz a **keep** módot használhatjuk.

### K-anonimitás
Egy anonimizált adatbázisról akkor mondjuk, hogy k-anonim, ha a benne található kvázi azonosító értékek bármely konkrét előfordulása esetén van még legalább k-1 db rekord, melyben ugyanazok a kvázi azonosító értékek fordulnak elő. Azaz bármely rekordhoz van még k-1 db hasonló rekord, abban az értelemben, hogy a kvázi azonosító értékeik megegyeznek.

### Ekvivalencia osztály
A fent megadott k-anonimitás definícióban az azonos attribútum értékekkel rendelkező rekordok egy ekvivalencia osztályt alkotnak. Ekkor külön ekvivalencia osztályt határoznak meg az adatbázis sémájában található kvázi azonosítók lehetséges értékeinek kombinációi. 
Például az alábbi 2-anonim adatbázisban (amennyiben az életkort és a nemet tekintjük kvázi azonosítónak) ekvivalencia osztályt alkotnak az 1.-2., a 3.-4 és az 5.-6. rekordok.

Életkor | Nem | Tesztpontszám 
------------ | ------------- | -------------
12 | Férfi | 79
12 | Férfi | 98
13 | Nő | 99
13 | Nő | 99 
14 | Férfi | 82
14 | Férfi | 85

Egy k-anonim adatbázisban minden ekvivalencia osztály számossága legalább k. Az algoritmus működése során az intervallum attribútumokhoz az ekvivalencia osztályokban nem egy konkrét érték tartozik, hanem egy alsó és felső korlátokkal adott intervallum. Például az életkorra vonatkozó ekvivalencia osztályok lehetnek az alábbiak: 

Ekvivalencia osztály id | Életkor
------------ | -------------
1 | 0 - 18
2 | 19 - 64
3 | 65 - 100

## Az algoritmus alapvető működése
Az algoritmus célja egy olyan adatbázis építése szerver oldalon, melyben minden ekvivalencia osztály legalább k-elemet tartalmaz. Ezáltal az adatbázis k-anonimitása minden pillanatban garantált lesz. Kliens oldali anonimizálás esetén a klienseket önálló ágenseknek tekintjük, melyek adatokat gyűjtenek. Amennyiben a kliensek rendelkezésre áll felküldendő adat, megpróbálja azt elhelyezni a szerveren található ekvivalencia osztályok valamelyikében. Ehhez először lekérdezi a szerveren található ekvivalencia osztályokat, majd megnézi, hogy az anonimizálandó adat beleillik-e valamelyik osztályba. 
* Amennyiben talál olyan ekvivalencia osztályt, melyben már van legalább k db elem, akkor egyszerűen abba az osztályba kell feltöltenie a keletkezett szenzitív adatokat. Ezt a kiválasztott ekvivalencia osztály egyedi azonosítójának megadásával és a szenzitív adatok feltöltésével teszi meg.
* Ha a kliens talál egy olyan ekvivalencia osztályt, ami illeszkedik a feltöltendő adataira, de abban még nincs legalább k db elem (ekkor ténylegesen még 0 elem lesz az adott ekvivalencia osztályba feltöltve, hiszen különben sérülne a k-anonimitás feltétele), akkor a feltöltéssel várnia kell. A kliens jelzi a szervernek a feltöltési igényét (az ekvivalencia osztály azonosítójának megadásával), a szerver pedig minden ekvivalencia osztályhoz nyilvántartja a feltöltési igények számát. Amennyiben az igények száma valamely osztálynál eléri a k értéket, akkor a szerver jelzi a klienseknek (erről részletesebben később), hogy megkezdhetik a feltöltést, akik erre egyszerre feltöltik az adataikat az ekvivalencia osztályba, ahol ennek hatására már legalább k db elem lesz.
* Ha a kliens által generált adatokra egyetlen ekvivalencia osztály sem illeszkedik, akkor a kliens generál egyet, és azt (az érzékeny adatok nélkül) felküldi a szerverre. Ilyenkor a szerver menti az ekvivalencia osztályt, és a feltöltési igények számát 1-re állítja (az azt létrehozó kliens biztos akar adatot küldeni). Ezután a kliens az előző esethez hasonlóan addig vár, amígy legalább k darab feltöltési igény össze nem gyűlik.

![Anonimization](img/anonimization.png)

## Az működés részletei
### A várakozási ciklus
Ha egy kliens adatot szeretne feltölteni egy még üres (például általa létrehozott) ekvivalencia osztályba, akkor a k-anonimitás megtartása érdekében várakoznia kell, amíg másik legalább k-1 db kliens nem jelezte a feltöltési igényét ugyanabba az ekvivalencia osztályba. Felmerülhet a kérdés, hogy honnan fogja tudni a kliens, hogy már összegyűlt-e a megfelelő számú igény. Ebben az implementációban erre szolgál a központi tábla. A szerver a publikusan elérhető (lekérdezhető) központi táblában tartja nyilván azoknak az ekvivalencia osztályoknak az azonosítóit, melyekbe már legalább k db feltöltési igény érkezett. A kliens ciklikusan lekérdezi a központi tábla tartalmát (jelen implementációban csak egy konkrét ekvivalencia osztály ID benne létére kérdez rá), és ha a központi tábla tartalmazza az ekvivalecia osztályt, akkor a kliens megkezdheti a feltöltést. Valójában mivel a kliensek különböző időpontban kérdezik le a tábla tartalmát, így ekkor még különböző időpontban töltenék fel az adataikat. Ennek elkerülésére a központi tábla rekordjai tartalmaznak egy, a szerver által meghatározott (jövőbeli) időpontot is. A kliensek a központi tábla lekérdezésekor, ha megtalálták a keresett azonosítót, akkor eltárolják a hozzá tartozó időpontot, és abban az időpontban fogják megkezdeni a szenzitív adatok feltöltését.

![Loop](img/loop.png)

### Hálózati hibák kezelése
A feltöltés időpontjában előfordulhat, hogy néhány kliens például hálózati hiba miatt nem képes az adatok feltöltésére. Ilyenkor sérülne az adatbázis k-anonimitása. Az ilyen esetekre került bevezetésre az adathalmaz definiálásakor megadható epszilon érték, melynek jelentősége, hogy a szerver nem k, hanem (k + epszilon) db feltöltési igényre vár, mielőtt közzéteszi az ekvivalencia osztály azonosítóját a központi táblában. Így ha legfeljebb epszilon darab kliensnek nem is sikerül feltöltenie az adatait a megfelelő időpontban, az adatbázis k-anonimitása akkor is biztosítva marad.

### Ekvivalencia osztályok finomítása
Ahhoz, hogy az adatok információtartalmát (a k-anonimitás megőrzése mellett) maximalizálni tudjuk, idővel érdemes lehet az ekvivalencia osztályok méretének csökkentése. Minnél több és kisebb ekvivalencia osztályunk van, annál nagyobb az adatok hasznossága, mivel annál pontosabban tudjuk egy adott szenzitív adathoz tartozó itervallum attribútumok értékét. Az ekvivalencia osztályok elemszámának alacsonyan tartására a szerver két módot támogat:
* A **szerver oldali ekvivalencia osztály finomítás** esetében az algoritmus paraméterként megadható (**max** néven) maximális elemszámú ekvivalencia osztályokat enged meg. Ha egy ekvivalencia osztályba feltöltött adatok száma eléri ezt a maximális értéket, akkor az adott ekvivalencia osztály inaktív lesz, azaz nem jelenik meg az ekvivalencia osztályok listázásakor. Helyette a szerver két másik új ekvivalencia osztályt hoz létre. Ekvivalencia osztály finomításakor program választ egy intervallum attribútumot, és annak méretét megfelezve hozza létre a két új ekvivalencia osztályt.

Például ha egy ekvivalencia osztályt a magasság attribútum mentén bontunk ketté, akkor az alábbi eredményt kapjuk:

Id | Magasság | Életkor | Aktív-e
------------ | ------------- | ------------- | -------------
1 | 1 - 50 | 0 - 18 | Nem
2 | 1 - 25 | 0 - 18 | Igen
3 | 25 - 50 | 0 - 18 | Igen

Ez jól működik mindaddig, amíg véges intervallumokat akarunk megfelezni. Mivel egy végtelen intervallumnak nem tudnánk, hogy hol van a felezőpontja, ezért a legelső ekvivalenciaosztályt (és azokat is, amik ennek az intervallumain kívül esnek) mindenképp kliens oldalon kell létrehozni, és a szerverre feltölteni. Ezután már a szerver oldalon lehet újabb ekvivalencia osztályokat generálni, ha valamely osztály mérete elérné a beállított maximum értéket.

* A másik lehetőség egy **preferált intervallumméret megadása** minden intervallum atrribútumhoz. Ekkor az új ekvivalencia osztályok generálása csak kliens oldalon történik. A kliens a dataset lekérdezésével hozzáfér az előre konfigurált preferált intervallumméretekhez, és adatait úgy általánosítja, hogy azok mindig a preferált méretű intervallumokba essenek. Például ha a preferált intervallumméretek a magasságra 20, az életkorra 5, akkor a magasság=200, életkor=49 konkrét értékeket a kliens például az alábbi módon általánosíthatja:

Magasság | Életkor 
------------- | ------------- 
190 - 210 | 45 - 50

Ennek a módnak a használata abban az esetben javasolt, ha azonos intervallum méretű ekvivalencia osztályokat akarunk létrehozni vagy az adott intervallum korlátai előre nem ismertek.

## Technológiák
A megvalósítás során felhasznált technológiák [itt](Tech.md) találhatók.

## Konfigurációs lehetőségek
### Dataset beállításai
Dataset létrehozásakor az alábbi beállítások adhatók meg:
* max: A maximális ekvivalencia osztály elemszám. (integer)
* k: A k-anonimimitásnál használt k értéke. (integer)
* e: Az epszilon érték, hogy k-nál mennyivel több igény érkezésére várjon a szerver. (integer)
* algorithm: Az anonimizáláshoz használt algoritmus, lehetséges értékei:
    * "client-side": A fent dokumentált kliens oldali anonimizálás, az ekvivalencia osztályok szerver oldali finomításával.
    * "client-side-custom": Kliens oldali anonimizálás, preferált intervallumméretekkel.
    * "mondrian": Szerver oldali anonimizálás, a Mondrian algoritmus használatával.
* mode: Az anonimizálás módja, amely lehet folytonos vagy egyszeri. Kliens oldali anonimizáláshoz a folytonos módot kell választani. Értékei:
    * "continuous": Folytonos
    * "single": Egyszeri

Továbbá a dataset létrehozásakor adhatóak meg a sémában szereplő attribútumok is.

### Attribútum beállítások
Minden attribútumra külön-külön az alábbi beállítások adhatók meg:
* name: A mező neve. (string)
* mode: A mező viselkedését adja meg az anonimizálás során. Lehetséges értékei:
    * "id": Egyedi azonosító, amit törölni kell.
    * "qid": Kvázi azonosító, ezek alapján történik az ekvivalencia osztályok kialakítása. (Csak szerver oldali anonimizáláshoz.)
    * "keep": Szenzitív adat, a szerver nem anonimizálja, hanem változatlanul tartja meg.
    * "cat": Intervallum típusú kvázi azonosító, az ekvivalencia osztályokban intervallumként lesz anonimizálva. (Csak kliens oldali anonimizáláshoz.)
    * "int": Kategorikus típusú kvázi azonosító, az ekvivalencia osztályokban a konkrét értéke jelenik meg, tovább nem anonimizálható. (Csak kliens oldali anonimizáláshoz.)
* type: A mező típusa. Kliens oldali anonimizálás esetén a "numeric" és a "string" támogatottak. Intervallum esetében szám típust kell megadni.
* preferedSize: Az adott mező intervallumának preferált mérete. Csak "client-side-custom" módú anonimizálás esetén van jelentősége.

## A kliens működése
A Proof-Of-Concept jelleggel, .Net Core platformon elkészített kliens két projektből épül fel.
* **Anonimization:** Az anonimizáláshoz szükséges modell és service osztályok valamint a Refit könyvtár felhasználásával készített REST kliens. Az újrafelhasználhatóság érdekében ezt a projektet Class Library-ként valósítottam meg.
* **AnonimizationClient:** A fent leírt Class library-t használó konzolos alkalmazás, ami bemutató jelleggel néhány előre beégetett dokumentum anonimizálását végzi el.

Mivel a megvalósított algoritmusban több kliens, egymástól függetlenül végezhet anonimizálásokat, illetve egymástól függetlenül jelezheti az adatfeltöltési igényeit a szerver felé, így ezt egy többszálú alkalmazással modelleztem. A kipróbálhatóság érdekében a kliens a központi tábla lekérdezése után (amennyiben az tartalmazza a keresett ekvivalencia osztályt) nem vár a szerver által kitűzött időpontig, hanem egyből megkezdi a szenzitív adatainak feltöltését a kiválasztott ekvivalencia osztályba.
