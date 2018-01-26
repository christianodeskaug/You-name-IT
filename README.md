# You-name-IT-
IS-105 oppgaver


OPPGAVESETT 1 - 26.01.2018


Oppgave 2: 
SVAR: Løst ved at alle tre gruppemedlemmene på hver sin PC endret "hello.go" til å printe ut sitt eget navn.




Oppgave 3:
3.1. Programmer i go utføres med kommandoet go run Da genereres det ingen objektfil (en fil som man kan utføre på en platform,
og hvis format er platformavhenging). Hvis man ønsker å generere 1 en objektfil, 
er kommandoet go build . Finn ut hva objektfiler heter for de mest brukte platformene (Unix/Linux, MS Windows, Mac OS X)! 
Hvorfor, etter din mening, har disse platformene så forskjellige objektfil-formater?

SVAR: 
MS Windows bruker PE (Portable Executable),
Linux bruker ELF (Executable and Linkable Format),
Mac OS X: Mach-O.
Disse ulike plattformene er bygget opp på forskjellige vis og individuelt fra de andre motpartene, derfor er det naturlig at de har sine egne metoder
kjøre filer på, og derfor egne objektfiler. Dette er kun de mest brukte plattformene, men det finnes mange flere operativsystemer som nytter andre
objektfiler laget kun til den plattformen den er tilegnet.




3.2. Gå gjennom https://tour.golang.org/basics/1 Hvilke forskjeller ser du i forhold til programmeringsspråket Java?
(forutsetter at alle er kjent med Java, hvis ikke, så kan man også sammenligne med andre programmeringsspråk)

SVAR: Vi ser ingen tegn eller bruk av semikolon i koden
"Fmt" ser ut til å erstatte "System.out." når det gjelder å printe ut linjer fra koden.
Helhetlig ser Golang til å virke mye mer lettvint enn Java, mindre lange koder og flere forkortelser.
"Func" setter i gang metoden, som er en pakke importert kun ved å skrive "package main", og deretter "import (fmt+math/rand)", veldig enkelt.
Ettersom alle programmerinsspråk er litt forskjellige på hvert sitt vis, så er det utvilsomt flere bemerkninger i andre typer Golang-koder i forhold til
Java som vi ikke har sett enda, men som vi sikkert kommer til å støte på ganske snart!
