# A simple golang/gin bookstore api 

### To create and a new migration run:

```sql-migrate new -config=config/migrations.yml MIGRATION-NAME```

### To execute the migrations run "make up" or:

Up: ```sql-migrate up -config=config/migrations.yml```

Down: ```sql-migrate down -config=config/migrations.yml```

### To generate new models use:

```sqlboiler -c config/sqlboiler.yml --wipe psql```

### To start the api run:
```
make start
```