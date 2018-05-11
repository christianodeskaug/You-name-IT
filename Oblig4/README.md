# Obligatorisk Oppgave 4 #
    
**Gruppemedlemmer**
* Christian Ødeskaug - christianodeskaug@hotmail.com
* Kristian Holmberg Sørensen - kristian_h_sorensen@hotmail.com
* Henrik Lindseth - henrik.lindseth1@gmail.com
    
**Systemspesifikasjon, hva går appen vår ut på?**

Vår gruppe endte til slutt på å planlegge og til slutt utvikle en værmeldingsapplikasjon, der man gjennom utfylling av sted og lokasjon vil finne ut nøyaktig hvordan værforholdet og temperaturen er på det gitte området. Vi tenker at dette er en bra idé, spesielt nå som det begynner å bli strålende fint og varmt utendørs! Vår applikasjon skal være lett å bruke, og oversiktlig nok til at enhver person kan finne ut hvordan den fungerer, og raskt få tilbakemelding på det han eller hun trenger å vite om værforholdet i dette sekund. Appen skal videføre sanntidsinformasjon til brukeren, som vil si at informasjonen som blir motatt skal være fersk og riktig, noe som forventes i en værmeldingsapp. Det skal være mulig å bruke appen i webleseren.
    
**Systemarkitektur, hvordan er appen vår bygget opp?**

Applikasjonen vår er en filbasert webside som forteller oss den nåværende værsituasjonen hvor som helst i verden. Filen er i html-format, og inneholder html-kode som fremstiller JSON-data fra en API-nettside (ligger i Oblig4 her på github) til strukturert data som kan leses som normal tekst. API'en viste oss mange funksjoner vi kunne legge til i vår værmelding, men vi valgte å strukturere applikasjonen vår slik at den er mest mulig forståelig, og dermed ha med de mest etterspurte og nødvendige funksjonene. Enhver populær nettleser kan brukes, men vi i gruppen gikk for Google Chrome som standardnettleser for applikasjonen vår. Vi har enhetstestet applikasjonen vår, og den fungerer som den skal innad i gruppen.

    
**Hvordan bruke vår værmeldingsapplikasjon?**

Appen åpnes i nettleser. Dette skal fungere i alle nettlesere, men vi opplevde at Chrome først måtte få tillatelse til å åpne programmet. Dette gjøres ved å endre egenskapene til Chrome og gi kommandoen "-allow-file-access-from-files" i "målfeltet". Når applikasjonen kjører er det bare å søke på ønsket by og land for så å utføre søket med søkeknappen (get weather). Byer i USA krever staten oppgitt for at søket skal bli godkjent. Eksempelvis: "City: Boston, State/Country: MA", da vil appen gi en presis værrapport fra Boston, Massachusetts.
Kodefilen lastes altså ned i html-format, og åpnes i Chrome, hvor du da vil umiddelbart få mulighet til søke etter værforholdene.

**Eksempel fra appen vår:**

![alt text](http://i67.tinypic.com/v5hcbb.png)

(åpnet i nettleser på PC, brukt Chrome-verktøy for å se hvordan den hadde sett ut på mobil)
