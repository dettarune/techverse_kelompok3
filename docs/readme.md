# UMKM Backend API - Postman Documentation

## Base URL
```
http://localhost:8080
```

## Environment Variables (Postman)
- `base_url`: http://localhost:8080
- `access_token`: *(akan diisi setelah login)*

---

## 1. Authentication

### 1.1 Register User
**POST** `{{base_url}}/api/v1/auth/signup`

**Headers:**
```http
Content-Type: application/json
```

**Request Body:**
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
}
```

**Response (200):**
```json
{
    "message": "user successfully registered"
}
```

### 1.2 Login User
**POST** `{{base_url}}/api/v1/auth/signin`

**Headers:**
```http
Content-Type: application/json
```

**Request Body:**
```json
{
    "email": "john@example.com",
    "password": "password123"
}
```

**Response (200):**
```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "message": "successfully signed in"
}
```

**Postman Script (Tests tab):**
```javascript
if (pm.response.code === 200) {
    const response = pm.response.json();
    pm.environment.set("access_token", response.access_token);
}
```

---

## 2. Store Management
### 2.1 Create Store
**POST** `{{base_url}}/api/v1/store`

**Headers:**
- Content-Type: application/json
- Authorization: Bearer {{access_token}}

**Request Body:** *(see full documentation for full JSON example)*

**Response (200):** `store successfully created`

### 2.2 Get Store
**GET** `{{base_url}}/api/v1/store`

**Headers:**
- Authorization: Bearer {{access_token}}

**Response (200):** *Store detail object*

### 2.3 Update Store
**PUT** `{{base_url}}/api/v1/store`

**Headers:**
- Content-Type: application/json
- Authorization: Bearer {{access_token}}

**Response (200):** `store successfully updated`

---

## 3. Product Management
### 3.1 Create Product
**POST** `{{base_url}}/api/v1/products`

### 3.2 Get All Products
**GET** `{{base_url}}/api/v1/products`

### 3.3 Get Product by ID
**GET** `{{base_url}}/api/v1/products/:id`

### 3.4 Update Product
**PUT** `{{base_url}}/api/v1/products/:id`

### 3.5 Delete Product
**DELETE** `{{base_url}}/api/v1/products/:id`

*(Each endpoint uses Bearer token and structured JSON similar to store)*

---

## 4. Website Builder

### 4.1 Create Website
**POST** `{{base_url}}/api/v1/website`

### 4.2 Get Website
**GET** `{{base_url}}/api/v1/website`

### 4.3 Update Website
**PUT** `{{base_url}}/api/v1/website`

### 4.4 Generate QR Code
**GET** `{{base_url}}/api/v1/website/qr`

---

## 5. Public Catalog

### 5.1 Get Public Catalog
**GET** `{{base_url}}/catalog/{domain}`

No auth required.

---

## 6. Order Management

### 6.1 Create Order (Public)
**POST** `{{base_url}}/api/v1/orders/{storeId}`

### 6.2 Get Store Orders
**GET** `{{base_url}}/api/v1/orders`

---

## 7. Error Responses

### 7.1 400 Validation Error
```json
{
    "error": "Key: 'CreateStoreRequest.Name' Error:Tag: 'required'"
}
```

### 7.2 401 Unauthorized
```json
{
    "error": "unauthorized"
}
```

### 7.3 404 Not Found
```json
{
    "error": "store not found"
}
```

### 7.4 500 Internal Server Error
```json
{
    "error": "internal server error"
}
```

---

## 8. Postman Collection Setup

**Pre-request Scripts**
```javascript
if (!pm.environment.get("access_token")) {
    throw new Error("Access token not found. Please login first.");
}
```

**Login Test Script**
```javascript
if (pm.response.code === 200) {
    const response = pm.response.json();
    pm.environment.set("access_token", response.access_token);
    console.log("Access token saved to environment");
}
```

---

## 9. Complete User Flow Example

1. Register → Login → Token saved
2. Create Store → Add Products
3. Create Website → Update & Publish → Generate QR Code
4. Test Public Access (Catalog + Order)
5. WhatsApp integration

---

## 10. Notes & Tips

- Gunakan format nomor internasional untuk WhatsApp
- Token berlaku 24 jam
- QR Code bisa dicetak untuk promosi offline

---

*Kelompok 3*

