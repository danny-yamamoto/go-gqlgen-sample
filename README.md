# go-gqlgen-sample
Try gqlgen.

### Installation.
```bash
apt-get update
sudo apt-get install -y sqlite3
mv mygraphql.db mygraphql.db.back
./setup.sh 
```

### SQLite
```bash
root ➜ /workspaces/go-gqlgen-sample (main) $ sqlite3 mygraphql.db
SQLite version 3.34.1 2021-01-20 14:10:07
sqlite> select * from users;
U_1|hsaki|
sqlite> 
```
