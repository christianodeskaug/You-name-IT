**Obligatorisk oppgave 1**

**Gruppemedlemmer:**

Christian �deskaug - christianodeskaug@hotmail.com
Henrik Lindseth - henrik.lindseth1@gmail.com
Kristian S�rensen - kristian_h_sorensen@hotmail.com


**Oppgave 1**

Vi har hatt problemer med innf�ring av tabeller og bilder i MD, s� vi legger heller ved bilder gjennom linker, h�per det er �lreit!
BILDE: http://oi66.tinypic.com/21b0hom.jpg

**Oppgave 1.A**
Bin�re tall er det simpleste tallsystemet vi har ettersom det kun bruker tallene 0 og 1. 
En kort og grei metode for � g� fra bin�re tall til b�de heksadesimal og desimaltall er � gange opp et bin�rt tall hver for seg. Eksempel: 1110.
Opph�y bakerste tallet med 1 (for � markere digits), og deretter doble eksponenten for hvert tall fremover. 
1*8
1*4 
1*2 
0*1 
Dette blir til sammen 8420. Pluss disse tallene sammen 8+4+2+0 og vi f�r tallet 14 (desimaltall), som igjen indikerer hexadesimaltallet �E� (alle hexadesimaltall over 10 er bokstaver, s� 14=E)
For � g� tilbake fra hexadesimal/desimal til bin�re m� man bare f�lge tallsystemet til bin�re/heksadesimal/desimal. N�r man har heksadesimalen E, s� vet man at dette er tallet 14, og 14 skrives som 1110 i bin�re tall.

**Oppgave 1.B**
Heksadesimaler f�lger sekstentallssystemet, s� for hvert ledd i en heksadesimal s� opph�yes tallene bakfra med 16. Eksempel: A4D.
A*256
4*16
D*1
Man vet at A = 10, og D = 13, s� regnestykket blir alts� 2560+64+13 = **2637.**

Metoden for det motsatte er ikke s� ulikt ettersom vi fortsatt holder oss i sekstentallsystemet. Eksempel: gj�re 680 fra heksadesimal til desimaltall.
Sekstentallsystemet = 16, 16^1, 16^2 (256) osv.
Man deler 680 p� det h�yest mulige tallet i systemet, som da er 256.
680/256 gir oss **2** f�r det gir oss mange desimaler, og det resterende desimaltaller er da 680-256*2=168. Deretter deler vi 168 med det lavere tallet i systemet (16). Dette gir oss **10,** og dette er som kjent **A** i heksadesimaltall. Da sitter vi igjen med tallet 8.
N�r vi sl�r i sammen disse tallene vi har f�tt for hvert delestykke, s� har vi alts� f�tt **2A8,** som er heksadesimaltallsformen til **680.**


**Oppgave 2**

**Oppgave 2.A:**
Vi har skrevet f�lgende Bubble-sort funksjon: 
BILDE: http://oi67.tinypic.com/2m30f2r.jpg


**Oppgave 2.B**
Vi byttet ut �BubbleSort� i testfilen med v�r egen funksjon, slik at vi til slutt kunne teste v�r kode men bench-testingen.
BILDE: http://oi68.tinypic.com/2ebzptf.jpg

**Oppgave 2.C**
Etter � ha kj�rt tester av b�de original sorting.go og v�r egen kode, s� er dette resultatet av v�r egen funksjon.

BILDE1: http://oi63.tinypic.com/2v8n3vq.jpg
BILDE2: http://oi64.tinypic.com/72yizb.jpg

Det vi la merke til n�r vi s� an fors�kene sammenlignet med Big-O tabellen, var at BubbleSort og BubbleSortModified algoritmene presterte i hovedsak d�rlige enn QuickSort metoden. Dette er ogs� noe man ser er vanlig p�  http://bigocheatsheet.com/. 
Space Complexity ser ogs� ut til � v�re st�rre for BubbleSort-funksjonene enn Quicksort, noe som ogs� stemmer i forhold til hva tabellen tilsier.


**Oppgave 3**
Vi har laget et program som kalles �uendelig.go� som simpelt nok genererer en uendelig loop med teksten �uendelig loop�. Gjennom CMD og med en Intel Core i7-7700 Kaby Lake prosessor s� bruker dette programmet ca 15-16% av CPU�en stabilt, noe vi synes er ganske mye for et s� lite og �ubetydelig� program, spesielt med tanke p� hvor god prosessoren er. Derimot s� bruker programmet kun 11,4 mb minne. Om man derimot kj�rte loopen via GoLand s� brukte programmet skyh�ye 50% av CPU og 512mb minne. P� CMD kunne man kun avslutte programmet ved � trykke �x� eller ctrl + c, noe som ikke ga en avslutningsmelding. P� GoLand kunne man klikke p� et r�dt stop-symbol og loopen avsluttet, og ga meldingen �Process finished with exit code 2�.
Filen brukt i oppgaven finnes under **�Opg3.go�**-filen i Github repositoryet v�rt.

**Oppgave 4**
Vi har dessverre ikke klart � svare p� denne oppgaven, ettersom den var altfor utfordrende basert p� kompetansen v�r. Vi pr�vde s� godt vi kunne � finne tips og � utnytte ressursene godt, men vi skj�nte rett og slett ikke oppgaven.
