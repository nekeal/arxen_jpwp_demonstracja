# Arxen demonstracja

Powyższe repozytoroum zawiera zadania oraz prezentacje dotyczące projektu [arxen](https://github.com/bartQu9/arxen).
W katalogu `zadania` znajduje się szereg plików których nazwy odpowiadają numerom zadań. 
Poprawność wykonania każdego zadania można sprawdzić przez wywołanie testu, a jak to zrobić jest napisane 
w sekcji Zadania. Gwiazdki obok numeru zadania symbolizują jego trudność.

## Zadania

### Zadanie 00*: Aktualizowanie czatu

W pliku ```exc00.go``` znajduje się uproszczony schemat działania `receivedPayloadChan()` 
z projektu (tutaj nazwany jest `Handle()`).
Zadaniem jest uzupełnienie funkcji `Handle` w taki sposób by za każdym razem gdy w kanale `receivedMessageChan`
pojawia się nowa wiadomość, była ona dodawana do `MessageList` w odpowiednim czacie.
Wszystkie czaty przechowywane są w `ChatList`.
  
#### Przykład
W kanale `receivedMessageChan` pojawia się wiadomość:
```go
Message {
	Text:   "przykladowa tresc wiadomosci",
	ChatID: "1",
}
```
to zostanie ona zapisana w 
```go
ChatList["1"].MessageList
```
jako kolejny element listy

#### Test
W celu przetestowania kodu należy:
```bash
cd zadania
go test -v exc00.go exc00_test.go
```



### Zadanie 01***: Optymalizacja kodu przez dodanie goroutines

W pliku `exc01.go` znajduje się funkcja `PrimesCounter` która na podstawie zadanego przedziału od `start` do `limit`
 zwraca ilość liczb pierwszych w tym przedziale. Funkcja `PrimesCounter` jest dość powolna przez fakt, 
 że w prostej pętli sprawdza każdą liczbę za pomocą `isPrime`. Aby zoptymalizować program 
 wstępnie zostały zaimplementowane funkcje `PrimesCounterWorker` oraz `PrimesCounterHandler`,
 dzięki którym program ma wykonać się szybciej
 
 Zadanie polega na dokończeniu implementacji `PrimesCounterWorker` oraz `PrimesCounterHandler`
 tak aby `PrimesCounterHandler` dzielił przedział na mniejsze części i każdą z tych części
  obliczał równolegle za pomocą funkcji `PrimesCounterWorker` wywołanej kilka razy. Należy skorzystać 
  z channelów oraz goroutines. 
  
#### Test i benchmark
W celu sprawdzenia poprawności działania kodu można wywołać test:
```bash
cd zadania
go test -v exc01.go exc01_test.go
```

Następnie aby sprawdzić czy kod faktycznie stał się bardziej wydajny wystarczy:
```bash
cd zadania
go test -bench=. -run=^PrimesCounterHandler
```

### Zadanie 02**: Prosty serwer http
W pliku ```exc02.go``` znajduje się przykładowa implementacja serwera z request routerem. Dzięki temu
można odwiedzać dynamicznie generowane na podstawie URL podstrony naszej domeny (podobne zagadnienie pojawiło się 
na WWWiJS gdzie odpowiedzialny za to był serwer Apache).  
Zadanie polega implementacji funkcji
 ```NameHandler```, która na podstawie żądania będzie dodawać do naszej struktury `Server`
 nowego użytkownika na podstawie adresu.
 
 #### Przykład
  
 request `localhost:8000/grzegorz/kuli` utworzy instancje struktury  
 ```go
PersonalInformation {
		Name:    "grzegorz",
		Surname: "kuli",
	}
```
Następnie doda utworzony element do tablicy `s.usersList`.

#### Test
W celu przetestowania kodu należy:
```bash
cd zadania
go test -v exc02.go exc02_test.go
```

---
**NOTE**

Test sprawdza tylko działanie `mux routera`. Jeśli chcesz zobaczyć działanie programu po
przez przeglądarkę należy wykonac
```bash
go build exc10.go
go run exc10.go
```

---

### Zadanie dodatkowe: Docker

Zadanie dodatkowe polega na utworzeniu pliku Dockerfile, który będzie kompilował zadania 
i przeprowadzał wybrane testy. Przykładowy [Dockerfile](https://github.com/bartQu9/arxen/blob/master/arxen-gui-golang/Dockerfile), 
może okazać się przydatny przy rozwiązywaniu.

#### Setup
Utwórz w wybranej dla siebie lokalizacji w projekcie plik `Dockerfile`. 