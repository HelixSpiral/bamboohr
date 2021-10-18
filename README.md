# bamboohr
An API wrapper for the BambooHR platform

## Usage
Create a new client with `bamboohr.New()`

```go
bbhr, err := bamboohr.New(map[string]interface{}{
  "company": "TestCompany",
  "apikey": "KEYHERE",
}
if err != nil {
  panic(err)
}
```

Use any of the built-in functions to query against your BambooHR instance.
```go
resp, err := bbhr.GetEmployees()
if err != nil {
    panic(err)
}

fmt.Println(string(resp))
```
