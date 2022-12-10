# API Documentation

## **User**

---

### **Register User**

Description<br>
This is an endpoint for user who doesn't have an account in our application.

Endpoint<br>
**POST** `api/v1/users/register`

Request Parameter<br>

```json
{
  "full_name": "full name",
  "email": "email",
  "password": "password"
}
```

Response Example<br>

```json
{
  "status": "Register Account Successfully",
  "error": null,
  "data": {
    "id": 1,
    "full_name": "full name",
    "email": "email",
    "balance": 0,
    "created_at": "time"
  }
}
```

Error Response Example<br>

```json
{
  "code": 400,
  "message": "email is already taken"
}
```

### **Login**

Description<br>
This is an endpoint for user to login in our application

Endpoint<br>
**POST** `api/v1/users/login`

Request Parameter<br>

```json
{
  "email": "email",
  "password": "password"
}
```

Response Example<br>

```json
{
  "status": "Login Successfully",
  "error": null,
  "data": {
    "token": "jwt access token"
  }
}
```

Error Response Example<br>

```json
{
  "code": 400,
  "message": "wrong password"
}
```

### **Topup Balance**

Description<br>
This is an endpoint where user can topup their's balance on the account to pay the products.

Endpoint<br>
**PATCH** `api/v1/users/topup`

Request Parameter<br>

```json
{
  "balance": 1000000
}
```

Response Example<br>

```json
{
  "status": "Top Up Balance Successfully",
  "error": null,
  "data": {
    "message": "Your balance has been successfully updated to 10000000"
  }
}
```

<br></br>

## Category

---

### **Create Category**

Description<br>
This is an endpoint for admin to create a new user's category.

Endpoint<br>
**POST** api/v1/categories

Request Parameter<br>

```bash
Headers: Bearer Token
```

```json
{
  "type": "string"
}
```

Response Example<br>

```json
{
  "id": "integer",
  "type": "string",
  "sold_product_amount": "integer",
  "created_at": "date"
}
```

### **View Category**

Description<br>
This is an endpoint to see all the user's categories.

Endpoint<br>
**GET** api/v1/categories

Request Parameter<br>

```bash
Headers: Bearer Token
```

Response Example<br>

```json
{
  "id": "integer",
  "type": "string",
  "sold_product_amount": "integer",
  "created_at": "date",
  "updated_at": "date",
  "Products": [
    {
      "id": "integer",
      "title": "string",
      "price": "stock",
      "created_at": "date",
      "updated_at": "date"
    },
    {
      "id": "integer",
      "title": "string",
      "price": "stock",
      "created_at": "date",
      "updated_at": "date"
    }
  ]
}
```

### **Update Category**

Description<br>
This is an endpoint for admin to edit all the user`s categories by categoryId.

Endpoint<br>
**PATCH** api/v1/categories/category:Id

Request Parameter<br>

```bash
Headers: Bearer Token
```

````json
{
    "type": "string"
}

Response Parameter<br>

```json
{
    "id": "integer",
    "type": "string",
    "sold_product_amount": "integer",
    "updated_at": "date"
}
````

### **Delete Category**

Description<br>
This is an endpoint for admin to delete all the user`s categories by categoryId.

Endpoint<br>
**DELETE** api/v1/categories/category:Id

Request Parameter<br>

```bash
Header: Bearer Token
```

Response Parameter<br>

```json
{
  "message": "Category has been successfully deleted"
}
```

## Product

---

### **Create Product**

Description<br>
This is an endpoint where user can create the product.

Endpoint<br>
**POST** `api/v1/products`

Request Parameter<br>

```json
{
  "title": "title",
  "price": 1,
  "stock": 1,
  "category_id": 1
}
```

Response Example<br>

```json
{
  "status": "SUCCESS",
  "error": null,
  "data": {
    "id": 1,
    "title": "title",
    "price": 10000,
    "stock": 1,
    "category_id": 1,
    "created_at": "2022-12-10T13:51:29.311693261Z"
  }
}
```

Error Response Example<br>

```json
{
  "code": 400,
  "message": "data not found"
}
```

### **View Products**

Description<br>
This is an endpoint where user can view the product.

Endpoint<br>
**GET** `api/v1/products`

Request Parameter<br>

```bash
Header: Bearer Token
```

Response Example<br>

```json
{
    "status": "SUCCESS",
    "error": null,
    "data": {
        "id": 1,
        "title": "title",
        "price": 10000,
        "stock": 1,
        "category_id": 1,
        "created_at": "2022-12-10T13:51:29.311693261Z"
    },
    {
        "id": 2,
        "title": "title",
        "price": 10000,
        "stock": 1,
        "category_id": 1,
        "created_at": "2022-12-10T13:55:29.311693261Z"
    }
}
```

Error Response Example<br>

```json
{
  "code": 404,
  "message": "data not found"
}
```

### **Update Product**

Description<br>
This is an endpoint where user can update the product.

Endpoint<br>
**PUT** `api/v1/products/:id`

Request Parameter<br>

```json
{
  "title": "title",
  "price": 1,
  "stock": 1,
  "category_id": 1
}
```

Response Example<br>

```json
{
  "status": "SUCCESS",
  "error": null,
  "data": {
    "id": 1,
    "title": "title",
    "price": 10000,
    "stock": 1,
    "category_id": 1,
    "created_at": "2022-12-10T13:51:29.311693261Z"
  }
}
```

Error Response Example<br>

```json
{
  "code": 400,
  "message": "data not found"
}
```

### **Delete Product**

Description<br>
This is an endpoint where user can delete the product.

Endpoint<br>
**DELETE** `api/v1/products/:id`

Request Parameter<br>

```json
{
  "title": "title",
  "price": 1,
  "stock": 1,
  "category_id": 1
}
```

Response Example<br>

```json
{
  "status": "SUCCESS",
  "error": null,
  "data": {
    "message": "Your product has been succesfully deleted"
  }
}
```

Error Response Example<br>

```json
{
  "code": 400,
  "message": "data not found"
}
```

## **Transaction History**

---

### **Transaction**

Description<br>
This is an endpoint where user doing a transaction in our application with their balance and our products.

Endpoint<br>
**POST** `api/v1/transactions`

Request Parameter<br>

```json
{
  "product_id": 1,
  "quantity": 1
}
```

Response Example<br>

```json
{
  "status": "Transaction Successful",
  "error": null,
  "data": {
    "message": "You have successfully purchased the product",
    "transaction_bill": {
      "total_price": 1000000,
      "quantity": 1,
      "product_title": "Logitech Mousepad"
    }
  }
}
```

Error Response Example<br>

```json
{
  "code": 400,
  "message": "not enough balance on account"
}
```

### **View My Transactions**

Description<br>
This is endpoint to see login users all transaction with specific products.

Endpoint<br>
**GET** `api/v1/transactions/my-transactions`

Request Parameter<br>

```bash
Header: Bearer Token
```

Response Example<br>

```json
{
  "status": "View My Transaction Success",
  "error": null,
  "data": [
    {
      "id": 7,
      "product_id": 1,
      "user_id": 5,
      "quantity": 1,
      "total_price": 1000000,
      "product": {
        "id": 1,
        "title": "Logitech Mousepad",
        "price": 1000000,
        "stock": 9,
        "category_Id": 1,
        "created_at": "2022-12-09T01:41:25Z",
        "updated_at": "2022-12-10T13:09:22Z"
      }
    }
  ]
}
```

Error Response Example<br>

```json
{
  "code": 404,
  "message": "data not found"
}
```

### **View Users Transactions**

Description<br>
This is an endpoint for admin to see all the user's transaction history.

Endpoint<br>
**GET** `api/v1/transactions/users-transactions`

Request Parameter<br>

```bash
Header: Bearer Token
```

Response Example<br>

```json
{
  "status": "View Users Transaction Success",
  "error": null,
  "data": [
    {
      "id": 7,
      "product_id": 1,
      "user_id": 5,
      "quantity": 1,
      "total_price": 1000000,
      "product": {
        "id": 1,
        "title": "Logitech Mousepad",
        "price": 1000000,
        "stock": 9,
        "category_Id": 1,
        "created_at": "2022-12-09T01:41:25Z",
        "updated_at": "2022-12-10T13:09:22Z"
      },
      "user": {
        "id": 5,
        "full_name": "full_name",
        "email": "email",
        "balance": 1000000,
        "created_at": "time",
        "updated_at": "time"
      }
    }
  ]
}
```

Error Response Example<br>

```json
{
  "code": 404,
  "message": "data not found"
}
```
