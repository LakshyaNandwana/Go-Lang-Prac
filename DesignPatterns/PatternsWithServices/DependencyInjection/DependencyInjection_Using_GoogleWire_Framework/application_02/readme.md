## Dependency Injection using Google WIRE


### How to use
Except for wire_gen.go copy all the files in 
- cmd
- domain
- user

__Use the following command__
#### Installing wire code generator
---
` $ go install github.com/google/wire/cmd/wire@latest `

##### downloading wire module
---
```
    $ go get github.com/google/wire@latest  
    go get: added github.com/google/wire v0.5.0
```

__Finally Run__
` $ wire ./...`

### Dependencies:
- Handler depends on Service
- Service depends on Repository
- Repository depends on Database Connection