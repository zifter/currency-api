```
heroku git:remote -a zifter-currency-api
```

Deploy
```
git push heroku master
```

```
docker image build -t currency-api:1.0 .
docker container run --publish 8000:8080 --detach --name api currency-api:1.0
docker container run --publish 8000:8080 --name api currency-api:1.0
```