## Buy Event Cli Application

---

### Database

- Postgresql

### Commands

- customer
  - add
  - get
  - update
  - delete


- order
  - add
  - get
  - update
  - delete

- log
  - display
  - clear

### Example

Create a new customer
```bash
go run main.go customer add -p "+998903599940" -e "t.mannonov@gmail.com"
```

Create a new order
```bash
go run main.go order add -c 1a880800-10dc-4c62-825c-9a0479163e90 -p "Iphone 11" -t 12000
```

Show logs
```bash
go run main.go log display
```
