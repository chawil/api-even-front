# front-end

## Docker 

```bash
$ docker build -t my-apache2 .
$ docker run -d --name api-even-front -p 8080:80 my-apache2
```

Then visit http://localhost:8080/