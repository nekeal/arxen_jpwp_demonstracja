

# Zadanie x
W pliku ```exc10.go``` znajduje się przykładowa implementacja serwera z request routerem. Dzięki temu
można odwiedzać dynamicznie generowane na podstawie URL podstrony naszej domeny (podobne zagadnienie pojawiło się 
na WWWiJS gdzie odpowiedzialny za to był serwer Apache).  
Zadanie polega implementacji funkcji
 ```NameHandler```, która na podstawie żądania będzie dodawać do naszej struktury `Server`
 nowego użytkownika na podstawie adresu.
 
 ###Przykład
  
 request `localhost:8000/grzegorz/kuli` utworzy instancje struktury  
 ```go
PersonalInformation {
		Name:    "grzegorz",
		Surname: "kuli",
	}
```
Następnie doda utworzony element do tablicy `s.usersList`.

### Test
W celu przerestowania kodu należy:
```bash
cd zadania
go test -v exc10.go exc10_test.go
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