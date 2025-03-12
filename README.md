# note-taking-app
simple note taking app in Go

//before test run,  sqlite3 ./db/database.db < ./db/test.sql

files

- backend
    - api
        - auth.go `api/auth`
        - note.go `api/note`
        - user.go `api/user`

    - db
        - db.go
        - note.go
        - user.go
- fronend
    - electron & axios
- main.go
