# mojango
A modern and fast Golang wrapper around the Mojang API

## Installation
Use the following command to download the wrapper:
```
go get -u github.com/Lukaesebrot/mojango
```

## Usage
Before you can use the wrapper, you need to create a Client:
```go
func main() {
    // Create a new mojango client
    client := mojango.New()
}
```

Using this client, you can now use the wrapper's functions as you like: 
```go
func main() {
    // Create a new mojango client
    client := mojango.New()
    
    // Fetch a player's UUID
    uuid, err := client.FetchUUID("ksebrt"); if err != nil {
        panic(err)
    }
    fmt.Println(uuid) // Output: 39cc0f91869a486494160d610f18b993
}
```

## Contribution/Help
If you found a bug, have any suggestions or want to improve some code, feel free to create an issue
or pull request! 
