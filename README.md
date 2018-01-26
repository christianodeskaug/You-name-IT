# You-name-IT-
IS-105 oppgaver


OPPGAVESETT 1 - 26.01.2018


Oppgave 2:
SVAR: L�st ved at alle tre gruppemedlemmene p� hver sin PC endret "hello.go" til � printe ut sitt eget navn.

Oppgave 3:
3.1. Programmer i go utf�res med kommandoet go run Da genereres det ingen objektfil (en fil som man kan utf�re p� en platform,
og hvis format er platformavhenging). Hvis man �nsker � generere 1 en objektfil,
er kommandoet go build . Finn ut hva objektfiler heter for de mest brukte platformene (Unix/Linux, MS Windows, Mac OS X)!
Hvorfor, etter din mening, har disse platformene s� forskjellige objektfil-formater?


SVAR:
MS Windows bruker PE (Portable Executable)
Linux bruker ELF (Executable and Linkable Format)
Mac OS X: Mach-O.
Disse ulike plattformene er bygget opp p� forskjellige vis og individuelt fra de andre motpartene, derfor er det naturlig at de har sine egne metoder
kj�re filer p�, og derfor egne objektfiler. Dette er kun de mest brukte plattformene, men det finnes mange flere operativsystemer som nytter andre
objektfiler laget kun til den plattformen den er tilegnet.



3.2. G� gjennom https://tour.golang.org/basics/1 Hvilke forskjeller ser du i forhold til programmeringsspr�ket Java?
(forutsetter at alle er kjent med Java, hvis ikke, s� kan man ogs� sammenligne med andre programmeringsspr�k)


SVAR: Vi ser ingen tegn eller bruk av semikolon i koden
"Fmt" ser ut til � erstatte "System.out." n�r det gjelder � printe ut linjer fra koden.
Helhetlig ser Golang til � virke mye mer lettvint enn Java, mindre lange koder og flere forkortelser.
"Func" setter i gang metoden, som er en pakke importert kun ved � skrive "package main", og deretter "import (fmt+math/rand)", veldig enkelt.
Ettersom alle programmerinsspr�k er litt forskjellige p� hvert sitt vis, s� er det utvilsomt flere bemerkninger i andre typer Golang-koder i forhold til
Java som vi ikke har sett enda, men som vi sikkert kommer til � st�te p� ganske snart!
