A program tesztel�s�hez az egy k�nyt�rban l�v� adatf�jlokat kell haszn�lni. 
Ahhoz hogy a tesztek megfelel�en fussanak alap helyzetbe kell �ll�tani a programot, az a configdir tartalma csak a k�vetkez� lehet:
 - mod.py
 - treedata.xml
 - az �ltalunk bem�solt szerkezetet le�r� json f�jl( ha �tnevezz�k recordstruct.json-r�l, akkor a scriptnek -s kapcsol�val kell megadni a f�jlnevet) 

A k�l�n k�nyvt�rakban l�v� json f�jlok �rj�k le az adatszerkezetet, a teszthez a megfelel� k�nyvt�rban l�v� json f�jl a configdir 
mapp�ba kell m�solni, a mellette l�v� csv �llmon�nyt ami az adatokat tartalmazza pedig a datadir k�nyvt�rba. 
Az anonimyze.py-vel ind�that� a program, a readme.txt tartalmazza a r�szletes le�r�st.