# Weather API with Redis caching

baby version(without concurrency concepts and rate limiting)

## How to use?

### 1. Install Go from official source

### 2. Install Redis from official source

### 3. Configure GOPATH and GOROOT

### 4. Register on [3rd Party Weather API](https://www.visualcrossing.com/sign-up)

### 5. Clone project into your IDE

```bash
git clone https://github.com/PureTeamLead/weather-api-redis.git
```

### 6. Get your API token within 3rd party API(step 4)

Sign up or Sign in, -> account -> key

### 7. Set up file with environment variables

*Create file .env
*Create variable with name: WEATHER_API_TOKEN
*Paste your key(API token)
*It would look like this:

```bash
WEATHER_API_TOKEN=4352HJ252JDH2JHFJD
```

*Add environment variable for Redis host -> localhost in your example could be

```bash
REDIS_HOST=localhost:3000
```

### 8. Start redis server on terminal

```bash
$ redis-server
```

### 9. Run application on another terminal window

```bash
$ cd cmd
$ go run main.go
```

### 10. Type in your web-address in your browser

For example: localhost:3000

### 11. Satisfaction zone
