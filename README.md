# Automatteori
Denne repositorien er et naivt eksempel på simulering av et endelig tilstandsautomat i programmeriingsspråket Golang.

## Beskrivelse av fremgangsmåten

* Lager en ny modul **automata**
* Knutter module til en repository med navn **automata** 

```
$ mkdir automata
$ cd automata
$ echo “# Automatteori” >> README.md
$ git init
$ go mod init github.com/digitnow/automata
$ git add README.md go.mod
$ git commit -m "commit README and go module configuration file"
$ git branch -M "main"
$ git remote add origin https://github.com/digitnow/automata.git
// "git push" til en ekstern server forutsetter at autentisering (token) er konfigurert
$ git push -u origin main
$ git tag v0.1.0
$ git push origin --tags
$ mkdir dfa 
$ cd dfa
// lage fil dfa_test.go, som inneholder definisjon av én pakke (package dfa),
// import av pakken “testing” (import “testing”) og funksjon for å teste Delta,
// som tar argumenter CurrentState og InputSymbol, og returnerer neste tilstand 
// basert på definisjon av δ ( func Delta(CurrentState, InputSymbol) ), 
// velge datatyper for tilstand og symbol (velger string for alt i første omgang,
// siden det er ikke forventet å gjøre aritmetiske operasjoner med hverken tilstand
// eller symboler fra input, 
// 
// lage fil dfa.go, definer package og funksjon Delta
```
```golang
   func Delta(currentState string, inputSymbol string) string { return “” }
```
// arbeide med go test inntill funksjonen tilfredstiller spesifikasjon til M1, 
// for eksempel
```
Algoritmen for simulering av et automat må starte med å definere input streng og start-tilsand, og lage en hovedløkke, hvor det avleses et symbol av gangen fra input, inntil det er ingen symboler i input igjen. Inne i løkken må man sjekke tilstand (med “if” eller “switch”) og ta en beslutning basert på en eksisterende definisjon av δ (funksjon Delta, som definerer alle overganger fra en tilstand til en annen, basert på symbolet i input streng).

## Vokabular

simulate
: d. To imitate the conditions or behaviour of (a situation or process) by means of a model, esp. for the purpose of study or of training; spec. to produce a computer model of (a process). For example: 1958  C. G. Gotlieb & J. N. P. Hume High-speed Data Processing xiii. 258: "A computer can simulate a warehouse, a factory, an oil refinery, or a river system, and if due regard is paid to detail the imitation can be very exact." (simulate, v. : Oxford English Dictionary, 2021)

## Notater om programmering i Golang

"nil is a predeclared identifier in Go that represents zero values for pointers, interfaces, channels, maps, slices and function types." (go dokumentasjon)

## Referanser

