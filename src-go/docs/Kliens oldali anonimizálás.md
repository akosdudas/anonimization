# Folyamatosan érkező adatok kliens oldali anonimizálása

## A folyamatos anonimizálás alapkoncepciója
A legtöbb ma létező anonimizálási megoldás a keletkező adatokat eredeti formájukban, anonimizálatlan formában gyűjti egy központi szerveren, ahol bizonyos idő elteltével vagy bizonyos mennyiségű adat összegyűjtése után végzik el az anonimizálást. Ennek legnagyobb veszélye, hogy az anonimizálás előtti időszakban a tárolt személyes adatok (az egyértelmű azonosítók és a hozzájuk tartozó szenzitív adatok) nem megfelelő védelem esetén illetéktelen kezekbe kerülhatnek. Erre a problémára nyújt megoldást a folyamatos anonimizálás, melynek során a szerverre beérkező adatok eredeti formájukban nem kerülnek tárolásra, hanem a szerver minden beérkezó adatot egyesével (vagy kis csoportokban) anonimizál, és az adatbázisba már csak az így anonimizált adatok kerülnek mentésre. Ezáltal még ha az adatbázishoz valaki hozzá is fér, akkor is csak az algoritmus által alkalmazott anonimizálási kritériumnak (pl.: k-anonimitás) megfelelő adatokat szerezhet meg.

## Az alkalmazás által támogatott anonimizálási módok
* [GitHub](Kliens oldali anonimizálás.md)
* Item 2

## A kliens oldali anonimizálás előnyei
A fenti koncepció hátránya, hogy az adatok felküldésekor a szerver továbbra is hozzáfér a még anonimizálatlan adatokhoz, ami pedig maga után vonja azt is, hogy az anonimizálást végző félnek továbbra meg kell felelnie a
GDPR által előírt követelményeknek. Ezen kívül teljes bizalom szükséges a kliens és a szerver között,, hiszen egy rosszindulatú szerver például titokban eltárolhatná a felküldött adatokhoz tartozó egyértelmű azonosítókat is. Ennek kiküszöbölésére érdemes megalkotni egy olyan algoritmust, amely lehetővé teszi a kliens (ahol az adat generálódik) számára azt, hogy a saját adatait már helyben anonimizálja, és csak ezután küldje fel őket a szerverre. Ezáltal a kliens teljes biztonságban tudhatja felküldött adatait, és a szerver sem lép interakcióba személyes adatokkal, így annak üzemeltetője mentesül a GDPR előírásai alól. A szerveren így eltárolt adatok később akár szabadon publikálhatóak, hiszen nem tartalmaznak személyes információt.

## Az alkalmazásban használt alapfogalmak
Az alkalmazás különböző típusú mezőket (attribútumokat) kezel. Egy attribútum besorolása a megfelelő kategóriába tervezői döntés, melyet az szerver oldali adatbázis (a programban dataset) sémájának meghatározásakor kell megadni. Az alkalmazás az alábbi típusú attribútumokat különbözteti meg:
* **Egyértelmű azonosító:** Egy adott személy kizárólagos tulajdonsága; olyan adat, mely e személy kilétét egyértelműen meghatározza (pl. név, személyi igazolvány szám). Ezek az anonimzzálás során eltávolításra kerüklnek. A dataset létrehozásakor az egyértelmű azonosítóknak **drop** módot adjunk meg, kliens oldali anonimizáláskor pedig ne küdjük fel ezeket az attribútumokat
* **Kvázi azonosító:** Olyan adat, mely önmagában nem, más információkkal közösen viszont már képes egy adott személy egyértelmű azonosítására (pl.születési idő, irányítószám).
* **Szenzitív adat:** Egy adott személyről érzékeny információt hordozó ismertető (pl. betegség, politikai nézetek).


* K-anonimitás:
* Item 2

## Az algoritmus működése