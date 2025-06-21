# UMKM Backend API - Postman Documentation

## Base URL
```
http://localhost:8080
```

## Environment Variables (Postman)
Create these variables in your Postman environment:
- `base_url`: `http://localhost:8080`
- `access_token`: (akan diisi setelah login)

---

## 1. Authentication

### 1.1 Register User
**POST** `{{base_url}}/api/v1/auth/signup`

**Headers:**
```
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

---

### 1.2 Login User
**POST** `{{base_url}}/api/v1/auth/signin`

**Headers:**
```
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
```
Content-Type: application/json
Authorization: Bearer {{access_token}}
```

**Request Body:**
```json
{
    "name": "Toko Kelontong Pak John",
    "description": "Toko kelontong lengkap di desa dengan berbagai kebutuhan sehari-hari",
    "logo": "https://example.com/logo.jpg",
    "address": "Jl. Mawar No. 123, Desa Sukamaju, Kec. Bogor Timur",
    "phone": "081234567890",
    "whatsapp": "6281234567890"
}
```

**Response (200):**
```json
{
    "message": "store successfully created",
    "data": {
        "id": 1,
        "name": "Toko Kelontong Pak John",
        "description": "Toko kelontong lengkap di desa dengan berbagai kebutuhan sehari-hari",
        "logo": "https://example.com/logo.jpg",
        "address": "Jl. Mawar No. 123, Desa Sukamaju, Kec. Bogor Timur",
        "phone": "081234567890",
        "whatsapp": "6281234567890",
        "user_id": 1,
        "is_active": true,
        "created_at": "2024-01-15T10:30:00Z",
        "updated_at": "2024-01-15T10:30:00Z"
    }
}
```

---

### 2.2 Get Store
**GET** `{{base_url}}/api/v1/store`

**Headers:**
```
Authorization: Bearer {{access_token}}
```

**Response (200):**
```json
{
    "data": {
        "id": 1,
        "name": "Toko Kelontong Pak John",
        "description": "Toko kelontong lengkap di desa dengan berbagai kebutuhan sehari-hari",
        "logo": "https://example.com/logo.jpg",
        "address": "Jl. Mawar No. 123, Desa Sukamaju, Kec. Bogor Timur",
        "phone": "081234567890",
        "whatsapp": "6281234567890",
        "user_id": 1,
        "is_active": true,
        "created_at": "2024-01-15T10:30:00Z",
        "updated_at": "2024-01-15T10:30:00Z"
    }
}
```

---

### 2.3 Update Store
**PUT** `{{base_url}}/api/v1/store`

**Headers:**
```
Content-Type: application/json
Authorization: Bearer {{access_token}}
```

**Request Body:**
```json
{
    "name": "Toko Kelontong Pak John - Updated",
    "description": "Toko kelontong terlengkap di desa dengan pelayanan 24 jam",
    "logo": "https://example.com/new-logo.jpg",
    "address": "Jl. Mawar No. 123, Desa Sukamaju, Kec. Bogor Timur",
    "phone": "081234567890",
    "whatsapp": "6281234567890",
    "is_active": true
}
```

**Response (200):**
```json
{
    "message": "store successfully updated",
    "data": {
        "id": 1,
        "name": "Toko Kelontong Pak John - Updated",
        "description": "Toko kelontong terlengkap di desa dengan pelayanan 24 jam",
        "logo": "https://example.com/new-logo.jpg",
        "address": "Jl. Mawar No. 123, Desa Sukamaju, Kec. Bogor Timur",
        "phone": "081234567890",
        "whatsapp": "6281234567890",
        "user_id": 1,
        "is_active": true,
        "created_at": "2024-01-15T10:30:00Z",
        "updated_at": "2024-01-15T11:45:00Z"
    }
}
```

---

## 3. Product Management

### 3.1 Create Product
**POST** `{{base_url}}/api/v1/products`

**Headers:**
```
Content-Type: application/json
Authorization: Bearer {{access_token}}
```

**Request Body:**
```json
{
    "name": "Beras Premium 5kg",
    "description": "Beras premium kualitas terbaik dari petani lokal, pulen dan wangi",
    "price": 65000,
    "image": "https://example.com/beras-premium.jpg",
    "category": "Sembako",
    "stock": 50
}
```

**Response (200):**
```json
{
    "message": "product successfully created",
    "data": {
        "id": 1,
        "name": "Beras Premium 5kg",
        "description": "Beras premium kualitas terbaik dari petani lokal, pulen dan wangi",
        "price": 65000,
        "image": "https://example.com/beras-premium.jpg",
        "category": "Sembako",
        "stock": 50,
        "store_id": 1,
        "is_active": true,
        "created_at": "2024-01-15T11:00:00Z",
        "updated_at": "2024-01-15T11:00:00Z"
    }
}
```

---

### 3.2 Get All Products
**GET** `{{base_url}}/api/v1/products`

**Headers:**
```
Authorization: Bearer {{access_token}}
```

**Response (200):**
```json
{
    "data": [
        {
            "id": 1,
            "name": "Beras Premium 5kg",
            "description": "Beras premium kualitas terbaik dari petani lokal, pulen dan wangi",
            "price": 65000,
            "image": "https://example.com/beras-premium.jpg",
            "category": "Sembako",
            "stock": 50,
            "store_id": 1,
            "is_active": true,
            "created_at": "2024-01-15T11:00:00Z",
            "updated_at": "2024-01-15T11:00:00Z"
        },
        {
            "id": 2,
            "name": "Minyak Goreng 1L",
            "description": "Minyak goreng berkualitas untuk kebutuhan memasak sehari-hari",
            "price": 15000,
            "image": "https://example.com/minyak-goreng.jpg",
            "category": "Sembako",
            "stock": 30,
            "store_id": 1,
            "is_active": true,
            "created_at": "2024-01-15T11:15:00Z",
            "updated_at": "2024-01-15T11:15:00Z"
        }
    ],
    "count": 2
}
```

---

### 3.3 Get Product by ID
**GET** `{{base_url}}/api/v1/products/1`

**Headers:**
```
Authorization: Bearer {{access_token}}
```

**Response (200):**
```json
{
    "data": {
        "id": 1,
        "name": "Beras Premium 5kg",
        "description": "Beras premium kualitas terbaik dari petani lokal, pulen dan wangi",
        "price": 65000,
        "image": "https://example.com/beras-premium.jpg",
        "category": "Sembako",
        "stock": 50,
        "store_id": 1,
        "is_active": true,
        "created_at": "2024-01-15T11:00:00Z",
        "updated_at": "2024-01-15T11:00:00Z"
    }
}
```

---

### 3.4 Update Product
**PUT** `{{base_url}}/api/v1/products/1`

**Headers:**
```
Content-Type: application/json
Authorization: Bearer {{access_token}}
```

**Request Body:**
```json
{
    "name": "Beras Premium 5kg - Grade A",
    "description": "Beras premium grade A kualitas terbaik dari petani lokal, pulen dan wangi",
    "price": 70000,
    "image": "https://example.com/beras-premium-a.jpg",
    "category": "Sembako",
    "stock": 45,
    "is_active": true
}
```

**Response (200):**
```json
{
    "message": "product successfully updated",
    "data": {
        "id": 1,
        "name": "Beras Premium 5kg - Grade A",
        "description": "Beras premium grade A kualitas terbaik dari petani lokal, pulen dan wangi",
        "price": 70000,
        "image": "https://example.com/beras-premium-a.jpg",
        "category": "Sembako",
        "stock": 45,
        "store_id": 1,
        "is_active": true,
        "created_at": "2024-01-15T11:00:00Z",
        "updated_at": "2024-01-15T12:30:00Z"
    }
}
```

---

### 3.5 Delete Product
**DELETE** `{{base_url}}/api/v1/products/1`

**Headers:**
```
Authorization: Bearer {{access_token}}
```

**Response (200):**
```json
{
    "message": "product successfully deleted"
}
```

---

## 4. Website Builder

### 4.1 Create Website
**POST** `{{base_url}}/api/v1/website`

**Headers:**
```
Content-Type: application/json
Authorization: Bearer {{access_token}}
```

**Request Body:**
```json
{
    "template": "modern",
    "domain": "toko-pak-john",
    "custom_css": "body { font-family: 'Arial', sans-serif; } .header { background-color: #4CAF50; }",
    "custom_html": "<div class='banner'>Selamat Datang di Toko Pak John!</div>"
}
```

**Response (200):**
```json
{
    "message": "website successfully created",
    "data": {
        "id": 1,
        "store_id": 1,
        "template": "modern",
        "custom_css": "body { font-family: 'Arial', sans-serif; } .header { background-color: #4CAF50; }",
        "custom_html": "<div class='banner'>Selamat Datang di Toko Pak John!</div>",
        "domain": "toko-pak-john",
        "is_published": false,
        "created_at": "2024-01-15T12:00:00Z",
        "updated_at": "2024-01-15T12:00:00Z"
    }
}
```

---

### 4.2 Get Website
**GET** `{{base_url}}/api/v1/website`

**Headers:**
```
Authorization: Bearer {{access_token}}
```

**Response (200):**
```json
{
    "data": {
        "id": 1,
        "store_id": 1,
        "template": "modern",
        "custom_css": "body { font-family: 'Arial', sans-serif; } .header { background-color: #4CAF50; }",
        "custom_html": "<div class='banner'>Selamat Datang di Toko Pak John!</div>",
        "domain": "toko-pak-john",
        "is_published": false,
        "created_at": "2024-01-15T12:00:00Z",
        "updated_at": "2024-01-15T12:00:00Z"
    }
}
```

---

### 4.3 Update Website
**PUT** `{{base_url}}/api/v1/website`

**Headers:**
```
Content-Type: application/json
Authorization: Bearer {{access_token}}
```

**Request Body:**
```json
{
    "template": "modern",
    "domain": "toko-pak-john-official",
    "custom_css": "body { font-family: 'Arial', sans-serif; background-color: #f5f5f5; } .header { background-color: #2E7D32; color: white; }",
    "custom_html": "<div class='banner'>🌟 Selamat Datang di Toko Pak John - Terpercaya Sejak 2020! 🌟</div>",
    "is_published": true
}
```

**Response (200):**
```json
{
    "message": "website successfully updated",
    "data": {
        "id": 1,
        "store_id": 1,
        "template": "modern",
        "custom_css": "body { font-family: 'Arial', sans-serif; background-color: #f5f5f5; } .header { background-color: #2E7D32; color: white; }",
        "custom_html": "<div class='banner'>🌟 Selamat Datang di Toko Pak John - Terpercaya Sejak 2020! 🌟</div>",
        "domain": "toko-pak-john-official",
        "is_published": true,
        "created_at": "2024-01-15T12:00:00Z",
        "updated_at": "2024-01-15T13:15:00Z"
    }
}
```

---

### 4.4 Generate QR Code
**GET** `{{base_url}}/api/v1/website/qr`

**Headers:**
```
Authorization: Bearer {{access_token}}
```

**Response (200):**
- Content-Type: image/png
- Binary PNG image data (QR Code)

**Note:** Response akan berupa file gambar PNG yang bisa disimpan langsung. QR Code berisi URL catalog: `http://localhost:8080/catalog/toko-pak-john-official`

---

## 5. Public Catalog

### 5.1 Get Public Catalog
**GET** `{{base_url}}/catalog/toko-pak-john-official`

**Headers:** (No authentication required)

**Response (200):**
```json
{
    "store": {
        "id": 1,
        "name": "Toko Kelontong Pak John",
        "description": "Toko kelontong lengkap di desa dengan berbagai kebutuhan sehari-hari",
        "logo": "https://example.com/logo.jpg",
        "address": "Jl. Mawar No. 123, Desa Sukamaju, Kec. Bogor Timur",
        "phone": "081234567890",
        "whatsapp": "6281234567890",
        "user_id": 1,
        "is_active": true,
        "created_at": "2024-01-15T10:30:00Z",
        "updated_at": "2024-01-15T11:45:00Z"
    },
    "products": [
        {
            "id": 1,
            "name": "Beras Premium 5kg - Grade A",
            "description": "Beras premium grade A kualitas terbaik dari petani lokal, pulen dan wangi",
            "price": 70000,
            "image": "https://example.com/beras-premium-a.jpg",
            "category": "Sembako",
            "stock": 45,
            "store_id": 1,
            "is_active": true,
            "created_at": "2024-01-15T11:00:00Z",
            "updated_at": "2024-01-15T12:30:00Z"
        },
        {
            "id": 2,
            "name": "Minyak Goreng 1L",
            "description": "Minyak goreng berkualitas untuk kebutuhan memasak sehari-hari",
            "price": 15000,
            "image": "https://example.com/minyak-goreng.jpg",
            "category": "Sembako",
            "stock": 30,
            "store_id": 1,
            "is_active": true,
            "created_at": "2024-01-15T11:15:00Z",
            "updated_at": "2024-01-15T11:15:00Z"
        }
    ],
    "count": 2
}
```

---

## 6. Order Management

### 6.1 Create Order (Public - Customer)
**POST** `{{base_url}}/api/v1/orders/1`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
    "customer_name": "Jane Doe",
    "customer_phone": "081987654321",
    "items": [
        {
            "product_id": 1,
            "quantity": 2,
            "price": 70000
        },
        {
            "product_id": 2,
            "quantity": 3,
            "price": 15000
        }
    ],
    "notes": "Mohon diantar sore hari sekitar jam 4. Rumah cat hijau di sebelah warung bu Siti."
}
```

**Response (200):**
```json
{
    "message": "order successfully created",
    "order": {
        "id": 1,
        "store_id": 1,
        "customer_name": "Jane Doe",
        "customer_phone": "081987654321",
        "items": "[{\"product_id\":1,\"quantity\":2,\"price\":70000},{\"product_id\":2,\"quantity\":3,\"price\":15000}]",
        "total_amount": 185000,
        "status": "pending",
        "notes": "Mohon diantar sore hari sekitar jam 4. Rumah cat hijau di sebelah warung bu Siti.",
        "created_at": "2024-01-15T14:30:00Z",
        "updated_at": "2024-01-15T14:30:00Z"
    },
    "whatsapp_url": "https://wa.me/6281234567890?text=*Pesanan%20Baru%20%23%201*%0A%0ANama:%20Jane%20Doe%0ATelepon:%20081987654321%0A%0A*Detail%20Pesanan:*%0A-%20Beras%20Premium%205kg%20-%20Grade%20A%20x2%20=%20Rp%20140000%0A-%20Minyak%20Goreng%201L%20x3%20=%20Rp%2045000%0A%0A*Total:%20Rp%20185000*%0A%0ACatatan:%20Mohon%20diantar%20sore%20hari%20sekitar%20jam%204.%20Rumah%20cat%20hijau%20di%20sebelah%20warung%20bu%20Siti.",
    "instructions": "Click the WhatsApp URL to send your order directly to the store"
}
```

---

### 6.2 Get Store Orders
**GET** `{{base_url}}/api/v1/orders`

**Headers:**
```
Authorization: Bearer {{access_token}}
```

**Response (200):**
```json
{
    "data": [
        {
            "id": 1,
            "store_id": 1,
            "customer_name": "Jane Doe",
            "customer_phone": "081987654321",
            "items": "[{\"product_id\":1,\"quantity\":2,\"price\":70000},{\"product_id\":2,\"quantity\":3,\"price\":15000}]",
            "total_amount": 185000,
            "status": "pending",
            "notes": "Mohon diantar sore hari sekitar jam 4. Rumah cat hijau di sebelah warung bu Siti.",
            "created_at": "2024-01-15T14:30:00Z",
            "updated_at": "2024-01-15T14:30:00Z"
        },
        {
            "id": 2,
            "store_id": 1,
            "customer_name": "Ahmad Rizki",
            "customer_phone": "082123456789",
            "items": "[{\"product_id\":1,\"quantity\":1,\"price\":70000}]",
            "total_amount": 70000,
            "status": "pending",
            "notes": "Bayar cash ya pak",
            "created_at": "2024-01-15T15:00:00Z",
            "updated_at": "2024-01-15T15:00:00Z"
        }
    ],
    "count": 2
}
```

---

## 7. Error Responses

### 7.1 Validation Error (400)
```json
{
    "error": "Key: 'CreateStoreRequest.Name' Error:Tag: 'required'"
}
```

### 7.2 Unauthorized (401)
```json
{
    "error": "unauthorized"
}
```

### 7.3 Not Found (404)
```json
{
    "error": "store not found"
}
```

### 7.4 Internal Server Error (500)
```json
{
    "error": "internal server error"
}
```

---

## 8. Postman Collection Setup

### Environment Variables
```
base_url: http://localhost:8080
access_token: (will be set automatically after login)
```

### Pre-request Scripts (for authenticated requests)
```javascript
// This script can be added to folder/collection level
if (!pm.environment.get("access_token")) {
    throw new Error("Access token not found. Please login first.");
}
```

### Test Scripts (for login request)
```javascript
if (pm.response.code === 200) {
    const response = pm.response.json();
    pm.environment.set("access_token", response.access_token);
    console.log("Access token saved to environment");
}
```

---

## 9. Complete User Flow Example

### Step 1: Register & Login
1. Register user dengan `POST /api/v1/auth/signup`
2. Login dengan `POST /api/v1/auth/signin`
3. Token akan otomatis tersimpan di environment

### Step 2: Setup Store
1. Create store dengan `POST /api/v1/store`
2. Add products dengan `POST /api/v1/products`

### Step 3: Create Website
1. Create website dengan `POST /api/v1/website`
2. Update dan publish dengan `PUT /api/v1/website`
3. Generate QR code dengan `GET /api/v1/website/qr`

### Step 4: Test Public Access
1. Get catalog dengan `GET /catalog/{domain}` (tanpa auth)
2. Create order dengan `POST /api/v1/orders/{storeId}`
3. Check orders dengan `GET /api/v1/orders`

### Step 5: WhatsApp Integration
- Setelah customer create order, akan mendapat WhatsApp URL
- URL tersebut akan membuka WhatsApp dengan pesan yang sudah terformat
- Store owner bisa langsung berkomunikasi dengan customer

---

## 10. Notes & Tips

### Authentication
- Semua endpoint kecuali auth dan public catalog memerlukan Bearer token
- Token expires dalam 24 jam
- Untuk testing, simpan token di Postman environment

### WhatsApp Integration
- Nomor WhatsApp harus dalam format internasional (contoh: 6281234567890)
- Pesan akan otomatis ter-encode untuk URL
- Customer akan langsung terhubung ke toko via WhatsApp

### QR Code
- QR code berisi URL ke public catalog
- Bisa dicetak untuk promosi offline
- Scan QR → Lihat katalog → Pesan → WhatsApp

### File Upload
- Untuk gambar, gunakan URL eksternal sementara
- Bisa dikembangkan dengan file upload endpoint terpisah

### Domain
- Domain website harus unique
- Digunakan untuk public catalog access
- Format: alphanumeric dan dash saja