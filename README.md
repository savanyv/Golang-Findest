# Golang Echo GORM API

Backend API menggunakan **Golang**, **Echo Framework**, dan **GORM** untuk mengelola transaksi pengguna.

## 🚀 Fitur API
- **Autentikasi dengan JWT**
- **Manajemen Transaksi**
  - 🔹 **Buat transaksi baru** → `POST /transactions`
  - 🔹 **Ambil daftar transaksi** (dengan filter opsional `userID` & `status`) → `GET /transactions`
  - 🔹 **Ambil transaksi berdasarkan ID** → `GET /transactions/:id`
  - 🔹 **Update status transaksi** → `PUT /transactions/:id`
  - 🔹 **Hapus transaksi** → `DELETE /transactions/:id`
- **Dashboard API**
  - 📊 **Total transaksi sukses hari ini** → `GET /dashboard/summary`
  - 📊 **Rata-rata jumlah transaksi per user**
  - 📊 **Daftar 10 transaksi terbaru**

## 📌 Teknologi yang Digunakan
- **Golang** (Echo Framework)
- **GORM** (ORM untuk PostgreSQL/MySQL)
- **JWT** (JSON Web Token untuk autentikasi)
- **Testify** (Unit testing)

## 🛠 Instalasi & Menjalankan Server
1. **Clone repository**
   ```sh
   git clone https://github.com/savanyv/Golang-Findest.git
   cd repo-name
     ```
   
2. **Buat file .env**
     ```env
     PGHOST = localhost
     PGPORT = 5432
     PGUSER = your-username
     PGPASSWORD = your-password
     PGDATABASE = your-database
     SECRETKEY = your-secret-key
     ```
3. **Install depedencies**
     ```sh
     go mod tidy
     ```
4. **Jalankan Server**
     ```sh
     go run cmd/main.go
     ```
     Server Berjalan di localhost:7000
## 📖 API Documentation
🔹Autentikasi

- 🔐 Login

     endpoint: ```POST /api/login```

     Request:
     ```json
     {
          "email": "user@exapmle.com",
          "password": "user"
     }
     ```
     Response:
     ```json
     {
          "data": {
               "id": 1,
               "name": "user",
               "email": "user@example.com",
               "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsImVtYWlsIjoic2F2YW55dkBleGFtcGxlLmNvbSIsImlzcyI6ImZpbmRlc3QiLCJleHAiOjE3Mzk3NzM2MDgsImlhdCI6MTczOTY4NzIwOH0.85KF2FfcVXvf_gbXUKZIm6r51TCaaMBGA3X3KDgsu6g"
          },
          "message": "successfully logged in"
     }
     ```
- 🔐 Register

     endpoint: ```POST /api/register```

     Request:
     ```json
     {
          "name": "user",
          "email": "user@example.com",
          "password": "user"
     }
     ```
     Response:
     ```json
     {
          "data": {
               "id": 1,
               "name": "user",
               "email": "user@example.com",
               "token": ""
          },
          "message": "successfully registered"
     }
     ```

🔹 Transaksi
- 📌 Buat transaksi baru

     endpoint: ```POST /api/transactions```

     Headers: ```Authorization: Bearer <token>```

     Request:
     ```json
     {
          "amount": 10000000
     }
     ```
     Response:
     ```json
     {
          "data": {
               "id": 1,
               "user_id": 1,
               "amount": 10000000,
               "status": "pending",
               "created_at": "2025-02-16T13:27:07.605913894+07:00"
          },
          "message": "successfully created transaction"
     }
     ```

- 📌 Ambil Daftar Transaksi

     endpoint: ```GET /api/transactions?user_id=1&status=success```
     ```json
     {
          "data": [
               {
                    "id": 1,
                    "user_id": 1,
                    "amount": 1000000,
                    "status": "pending",
                    "created_at": "2025-02-16T13:25:16.88007+07:00"
               },
          ],
          "message": "successfully get transactions"
     }
     ```

- 📌 Ambil Transaksi Berdasarkan ID

     endpoint: ```GET /api/transactions/1```

     headers: ```Authorization: Bearer <token>```

     Response:
     ```json
     {
          "data": {
               "id": 1,
               "user_id": 1,
               "amount": 1000000,
               "status": "pending",
               "created_at": "2025-02-16T13:25:16.88007+07:00"
          },
          "message": "successfully get transaction"
     }
     ```
- 📌 Update Status Transaksi

     endpoint: ```PUT /api/transactions/1```

     request:
     ```json
     {
          "status": "success"
     }
     ```
     response:
     ```json
     {
          "data": {
               "id": 1,
               "user_id": 2,
               "amount": 10000000,
               "status": "success",
               "created_at": "2025-02-16T13:27:07.605913+07:00"
          },
          "message": "successfully update status transaction"
     }
     ```

- 📌 Delete Transaction

     endpoint: ```DELETE /api/transactions/1```

     headers: ```Authorization: Bearer <token>```

     response:
     ```json
     {
          "message": "successfully delete transaction"
     }
     ```

🔹Dashboard
- 📊 Get Dashboard Summary

     endpoint: ```GET /api/dashboard/summary```

     response:
     ```json
     {
          "data": [
               {
                    "total_success_transactions": 13000000,
                    "average_transaction_per_user": 3750000,
                    "latest_transactions": [
                         {
                              "id": 1,
                              "user_id": 1,
                              "amount": 10000000,
                              "status": "success",
                              "created_at": "2025-02-16T13:27:07.605913+07:00"
                         },
                    ]
               }
          ],
          "message": "successfully get dashboard summary"
     }
     ```
