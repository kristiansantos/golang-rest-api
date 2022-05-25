### Go version 1.18

## Deployment instructions

```bash
  $ docker-compose up  # Run command to start postgres and pgadmin.
  $ go run main.go # Start server.

  Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.
```

## Services

* List planets (index) [GET/api/v1/planets]
* Get planet (show) [GET/api/v1/planets/:id]
* Create new planet (Create) [POST/api/v1/planets]
* Update planet (Update) [PUT/api/v1/planets/:id]
* Delete planet (Delete) [DELETE/api/v1/planets/:id]
  * Body params
  ```
    {
        "name": "Marte",
        "climate": "hostile",
        "ground": 11,
        "film_appearances": 50
    }
  ```