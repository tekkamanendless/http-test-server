# HTTP Test Server
This is an HTTP test server that you can use for testing, debugging, etc.

## Endpoints

### /status/{status}
This responds to any method (`GET`, `POST`, etc.) with the status code specifed in `{status}`.
For example, `GET /status/120` would return an empty page with status code 120.

