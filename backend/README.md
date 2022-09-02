# Appcreative Backend

**Please do not forget to copy the .env.example file to .env file for app's configuration!**

The APP_PORT=8123 is preferable because frontend is going to address to http://localhost:8123, so if you change port here in .env, don't forget to change it in the frontend app ([main.js](https://github.com/arthurshafikov/appcreative/blob/master/frontend/main.js#L2)).

# Commands

## Docker-compose

Run the application
```
make up
```

Down the application
```
make down
```

---
## Tests

Run the unit tests
```
make test
```
