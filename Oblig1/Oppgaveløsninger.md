**Obligatorisk oppgave 1**

**Gruppemedlemmer:**

Christian Ødeskaug - christianodeskaug@hotmail.com
Henrik Lindseth - henrik.lindseth1@gmail.com
Kristian Sørensen - kristian_h_sorensen@hotmail.com


**Oppgave 1**

Vi har hatt problemer med innføring av tabeller og bilder i MD, så vi legger heller ved bilder gjennom linker, håper det er ålreit!
BILDE: http://oi66.tinypic.com/21b0hom.jpg

**Oppgave 1.A**
Binære tall er det simpleste tallsystemet vi har ettersom det kun bruker tallene 0 og 1. 
En kort og grei metode for å gå fra binære tall til både heksadesimal og desimaltall er å gange opp et binært tall hver for seg. Eksempel: 1110.
Opphøy bakerste tallet med 1 (for å markere digits), og deretter doble eksponenten for hvert tall fremover. 
1*8
1*4 
1*2 
0*1 
Dette blir til sammen 8420. Pluss disse tallene sammen 8+4+2+0 og vi får tallet 14 (desimaltall), som igjen indikerer hexadesimaltallet “E” (alle hexadesimaltall over 10 er bokstaver, så 14=E)
For å gå tilbake fra hexadesimal/desimal til binære må man bare følge tallsystemet til binære/heksadesimal/desimal. Når man har heksadesimalen E, så vet man at dette er tallet 14, og 14 skrives som 1110 i binære tall.

**Oppgave 1.B**
Heksadesimaler følger sekstentallssystemet, så for hvert ledd i en heksadesimal så opphøyes tallene bakfra med 16. Eksempel: A4D.
A*256
4*16
D*1
Man vet at A = 10, og D = 13, så regnestykket blir altså 2560+64+13 = **2637.**

Metoden for det motsatte er ikke så ulikt ettersom vi fortsatt holder oss i sekstentallsystemet. Eksempel: gjøre 680 fra heksadesimal til desimaltall.
Sekstentallsystemet = 16, 16^1, 16^2 (256) osv.
Man deler 680 på det høyest mulige tallet i systemet, som da er 256.
680/256 gir oss **2** før det gir oss mange desimaler, og det resterende desimaltaller er da 680-256*2=168. Deretter deler vi 168 med det lavere tallet i systemet (16). Dette gir oss **10,** og dette er som kjent **A** i heksadesimaltall. Da sitter vi igjen med tallet 8.
Når vi slår i sammen disse tallene vi har fått for hvert delestykke, så har vi altså fått **2A8,** som er heksadesimaltallsformen til **680.**


**Oppgave 2**

**Oppgave 2.A:**
Vi har skrevet følgende Bubble-sort funksjon: 
BILDE: http://oi67.tinypic.com/2m30f2r.jpg


**Oppgave 2.B**
Vi byttet ut “BubbleSort” i testfilen med vår egen funksjon, slik at vi til slutt kunne teste vår kode men bench-testingen.
BILDE: http://oi68.tinypic.com/2ebzptf.jpg

**Oppgave 2.C**
Etter å ha kjørt tester av både original sorting.go og vår egen kode, så er dette resultatet av vår egen funksjon.

BILDE1: http://oi63.tinypic.com/2v8n3vq.jpg
BILDE2: http://oi64.tinypic.com/72yizb.jpg

Det vi la merke til når vi så an forsøkene sammenlignet med Big-O tabellen, var at BubbleSort og BubbleSortModified algoritmene presterte i hovedsak dårlige enn QuickSort metoden. Dette er også noe man ser er vanlig på  http://bigocheatsheet.com/. 
Space Complexity ser også ut til å være større for BubbleSort-funksjonene enn Quicksort, noe som også stemmer i forhold til hva tabellen tilsier.


**Oppgave 3**
Vi har laget et program som kalles “uendelig.go” som simpelt nok genererer en uendelig loop med teksten “uendelig loop”. Gjennom CMD og med en Intel Core i7-7700 Kaby Lake prosessor så bruker dette programmet ca 15-16% av CPU’en stabilt, noe vi synes er ganske mye for et så lite og “ubetydelig” program, spesielt med tanke på hvor god prosessoren er. Derimot så bruker programmet kun 11,4 mb minne. Om man derimot kjørte loopen via GoLand så brukte programmet skyhøye 50% av CPU og 512mb minne. På CMD kunne man kun avslutte programmet ved å trykke “x” eller ctrl + c, noe som ikke ga en avslutningsmelding. På GoLand kunne man klikke på et rødt stop-symbol og loopen avsluttet, og ga meldingen “Process finished with exit code 2”.
Filen brukt i oppgaven finnes under **“Opg3.go”**-filen i Github repositoryet vårt.

**Oppgave 4**
Vi har dessverre ikke klart å svare på denne oppgaven, ettersom den var altfor utfordrende basert på kompetansen vår. Vi prøvde så godt vi kunne å finne tips og å utnytte ressursene godt, men vi skjønte rett og slett ikke oppgaven.
