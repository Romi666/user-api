# USER API

# SIMPLE GO REST API

## What you need
- IntelliJ, Eclipse, Or VSCode
- Postman application for sending http request and JSON
- MySQL or Postgre for Database Management System
- Go Language

### Installation
- Clone project
```sh
https://github.com/Romi666/user-api.git
```

- Go to directory
- Configure database in file .env.example
- Change the name to .env
- Run the program by typing the command on terminal
```sh
go run main.go
```


# Example
### Create User

http://localhost:8080/api/v1/user [POST]

JSON Body

```
{
    "nama_user" : "Romi",
    "tanggal_lahir" : "1999-11-16",
    "no_ktp" : "1234123412341234",
    "pekerjaan" : "PNS",
    "pendidikan_terakhir" : "SMA"
}
```

Output

```
{
    "status": "200",
    "message": "Success",
    "data": {
        "message": "success create user"
    }
}
```

### Update User

http://localhost:8080/api/v1/user?user_id=1 [POST]

JSON Body

```
{
    "nama_user" : "Romi Abdul Yusuf",
    "tanggal_lahir" : "1999-11-16",
    "no_ktp" : "1234123412341234",
    "pekerjaan" : "PNS",
    "pendidikan_terakhir" : "SMA"
}
```

Output

```
{
    "status": "200",
    "message": "Success",
    "data": {
        "message": "success update user"
    }
}
```

### Get All User

http://localhost:8080/api/v1/users [GET]

Output
```
{
    "status": "200",
    "message": "Success",
    "data": [
        {
            "id": "1",
            "nama_user": "Romi Abdul Yusuf",
            "tanggal_lahir": "1999-11-16",
            "no_ktp": "1234123412341234",
            "pekerjaan": "PNS",
            "pendidikan_terakhir": "SMA"
        },
        {
            "id": "2",
            "nama_user": "Dewi",
            "tanggal_lahir": "2000-12-24",
            "no_ktp": "1234123412341234",
            "pekerjaan": "PNS",
            "pendidikan_terakhir": "SMA"
        }
    ]
}
```

### Get User By ID

http://localhost:8080/api/v1/user?user_id=1 [GET]

Output

```
{
    "status": "200",
    "message": "Success",
    "data": {
        "id": "1",
        "nama_user": "Romi Abdul Yusuf",
        "tanggal_lahir": "1999-11-16",
        "no_ktp": "1234123412341234",
        "pekerjaan": "PNS",
        "pendidikan_terakhir": "SMA"
    }
}
```
