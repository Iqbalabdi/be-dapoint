# be-dapoint
Repository for dapoint app API

## How to use
1. clone this repo
2. run `go mod verify`
3. run `go mod tidy`
4. rename file `config.toml.example` to `config.toml` in folder `/config`
5. fill username, password, name (database name) in `config.toml`
6. move to `/app` folder in terminal
7. run `go run main.go`

## Reference
This repository using clean architecture known as hexagonal architecture  
![image](https://github.com/Iqbalabdi/be-dapoint/assets/75016595/f09a1016-3f6c-4c8f-aa82-0f628c1f59c6)

Credits:  
https://github.com/bxcodec/go-clean-arch

## ERD
![image](https://github.com/Iqbalabdi/be-dapoint/assets/75016595/dffbc2f7-115d-4fcb-b231-5fcf09afa7ff)


# List API Dapoint
## 1. Register user
```
method: POST
endpoint: "/user/register"
request form-data: {
    name, email, password
}
response(200): {
    "message":"Success"
}
```

## 2. List All user
```
method: GET
endpoint: "/user"
response(200): {
    "message":"Success",
      data: [
              {entities.User}
          ]
}
```

## 3. List user by ID
```
method: GET
endpoint: "/user/{id}"
response(200): {
    "message":"Success",
      data: [
              {entities.User}
          ]
}
```

## 4. Login user
```
method: POST
endpoint: "/user/{id}"
request form-data: {
    email, password
}
response(200): {
    "message":"Success"
}
```

## 5. Update user
```
method: PUT
endpoint: "/user/{id}"
request form-data: {
      {entities.User}
}
response(200): {
    "message":"Success",
    data: [
        {entities.User}
    ]
}
```

## 6. Get Total User
```
method: GET
endpoint: "/user/total"
response(200): {
    "message":"Success",
    data: [
        {entities.User}
    ]
}
```

## 7. Get Transaction by UserID
```
method: GET
endpoint: "/user_transaction/{:userID}"
response(200): {
    "message":"Success",
    data: [
        {entities.Transaction}
    ]
}
```

## 8. Get All redeemed voucher by UserID
```
method: GET
endpoint: "/redeem_history/{:userid}"
response(200): {
    "message":"Success",
    data: [
        {entities.RedeemVoucher}
    ]
}
```

## 9. Create Voucher
```
method: POST
endpoint: "/vouchers/create"
header: {
    Authorization Token
}
request form-data: {
      {entities.Voucher}
}
response(200): {
    "message":"Success",
}
```

## 10. Get All Voucher
```
method: POST
endpoint: "/vouchers/getall"
header: {
    Authorization Token
}
response(200): {
    "message":"Success",
    data: [
        {entities.Voucher}
    ]
}
```

## 11. Update Voucher
```
method: PUT
endpoint: "/vouchers/update/{:id}"
header: {
    Authorization Token
}
request form-data: {
      {entities.Voucher}
}
response(200): {
    "message":"Success",
    data: [
        {entities.Voucher}
    ]
}
```

## 13. Get Voucher by ID
```
method: PUT
endpoint: "/vouchers/getbyid/{:id}"
header: {
    Authorization Token
}
response(200): {
    "message":"Success",
    data: [
        {entities.Voucher}
    ]
}
```

## 14. Get Voucher by Type
```
method: PUT
endpoint: "/vouchers/getbyid/:tipe"
header: {
    Authorization Token
}
response(200): {
    "message":"Success",
    data: [
        {entities.Voucher}
    ]
}
```

## 15. Redeem Voucher
```
method: POST
endpoint: "user/redeem_voucher/"
header: {
    Authorization Token
}
request form-data: {
      {entities.RedeemVoucher}
}
response(200): {
    "message":"Success",
    data: [
        {entities.RedeemVoucher}
    ]
}
```

## 16. Transaction Create
```
method: POST
endpoint: "/admin/user_transaction/create"
header: {
    Authorization Token
}
request form-data: {
      {entities.Transactions}
}
response(200): {
    "message":"Success",
    data: [
        {entities.Transactions}
    ]
}
```

## 17. Get All Transaction
```
method: GET
endpoint: "/admin/user_transaction/getall"
response(200): {
    "message":"Success",
    data: [
        {entities.Transactions}
    ]
}
```
