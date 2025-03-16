# note-taking-app
simple note taking app in Go

//before test run,  sqlite3 ./db/database.db < ./db/test.sql

files

- backend
    - api
        - auth.go 
            - <b>api/auth</b>
                - `/login`
        - note.go 
            - <b>api/note</b>
                - `/all`
                - `/create`
                - `/update`
                - `/delete`
        - profile.go 
            - <b>api/profile</b>
                - `/me`
    - db (sqlite)
        - db.go
        - note.go
        - profile.go
- fronend
    - electron & axios
- main.go
