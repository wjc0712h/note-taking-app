# note-taking-app
simple note taking app in Go

//before test run,  sqlite3 ./db/database.db < ./db/test.sql

files

- backend
    - api
        - auth.go <b>api/auth</b>
            - `/login`
        - note.go <b>api/note</b>
            - `/all`
            - `/create`
            - `/update`
            - `/delete`
        - profile.go <b>api/profile</b>
            - `/me`
            - `/new`
    - db (sqlite)
        - db.go
        - note.go
        - profile.go
- fronend
    - electron & axios
- main.go

build - `npx electron-packager . note --platform=darwin --arch=x64 --out=release-build --overwrite --extra-resource=./dist/server --icon=icon.icns`