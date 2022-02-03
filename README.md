# GoMon
A simple uptime monitoring service written in go.

---
## Quick start:
To run the service in a test environemnt you can simply use the provided docker compose file:
```bash
cd http-server  # pwd -> <proj-repo>/http-broker/
dokcer-compose up -d
```
*Warning: this compose file might not be suitable for use in prodcution environments and should only be used for testing purposes*

---
## API
Clients can interact witht he service trough a REST http api. The api currently supports:

- User signup
- User login with JWT authentication
- Register url to be monitored
- List registered urls
- Get number of failed and successful calls for specific url

For more details on the exact endpoints please refer to the provided postman collection.

---
### Deployment
