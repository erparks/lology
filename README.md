# lology
WIP: League of Legends API client

# API key
Riot invalidates development API keys every 24 hours. To easily update your API key, follow the instructions below.

Sign in at https://developer.riotgames.com/ to find your API key

```
go run cmd/keyset/main.go YOUR_KEY_HERE
```

# Is it working?

To confirm that everything is set up correctly, start the server:
```
go run cmd/lology/main.go
```
Visit: http://localhost:8080/matches?name=doublelift
