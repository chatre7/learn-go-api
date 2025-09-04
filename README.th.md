# Learn API
[![Test and Coverage](https://github.com/chatre7/learn-go-api/actions/workflows/test-coverage.yml/badge.svg)](https://github.com/chatre7/learn-go-api/actions/workflows/test-coverage.yml)
[![codecov](https://codecov.io/gh/chatre7/learn-go-api/branch/main/graph/badge.svg)](https://codecov.io/gh/chatre7/learn-go-api)

ภาษา: [English](README.md) | ไทย

RESTful API ที่พัฒนาด้วย Go, PostgreSQL และ Docker

## สารบัญ

- [คุณสมบัติ](#คุณสมบัติ)
- [โครงสร้างโปรเจกต์](#โครงสร้างโปรเจกต์)
- [เว็บเฟรมเวิร์ก](#เว็บเฟรมเวิร์ก)
- [API Endpoints](#api-endpoints)
- [เริ่มต้นใช้งาน](#เริ่มต้นใช้งาน)
- [เอกสาร API](#เอกสาร-api)
- [การทดสอบ](#การทดสอบ)
- [โครงสร้างฐานข้อมูล](#โครงสร้างฐานข้อมูล)

## คุณสมบัติ

- การทำงานแบบ CRUD สำหรับเอนทิตี
- เชื่อมต่อฐานข้อมูล PostgreSQL
- มี Unit Tests ครอบคลุม
- ทำงานผ่าน Docker ได้
- โครงสร้างแบบแยกความรับผิดชอบ (Clean Architecture)
- เว็บเฟรมเวิร์ก: [Fiber](https://gofiber.io/) – เฟรมเวิร์กสไตล์ Express สำหรับ Go

## โครงสร้างโปรเจกต์

```
.
├── cmd/
│   └── api/
│       └── main.go          # จุดเริ่มต้นของแอปพลิเคชัน
├── internal/
│   ├── handlers/            # ฮैंडเลอร์สำหรับคำขอ HTTP
│   ├── services/            # ตรรกะทางธุรกิจ (Business Logic)
│   ├── models/              # โครงสร้างข้อมูล (Data Structures)
│   ├── repository/          # เลเยอร์เข้าถึงข้อมูล (Data Access)
│   ├── database/            # ยูทิลิตีสำหรับเชื่อมต่อฐานข้อมูล
│   └── app/                 # ตัวช่วยประกอบแอป (NewFiberApp)
├── pkg/
│   ├── errors/              # ยูทิลิตีสำหรับจัดการข้อผิดพลาด
│   └── validation/          # ยูทิลิตีสำหรับตรวจสอบความถูกต้องของข้อมูล
├── tests/
│   ├── e2e/                 # การทดสอบแบบ End-to-End
│   ├── handlers/            # การทดสอบเลเยอร์ HTTP
│   ├── services/            # การทดสอบตรรกะทางธุรกิจ
│   ├── repository/          # การทดสอบเลเยอร์เข้าถึงข้อมูล
│   └── app/                 # การทดสอบการประกอบแอป (health/routes)
├── docs/                    # เอกสาร Swagger
├── Dockerfile               # การตั้งค่า Container
├── docker-compose.yml       # การตั้งค่าแบบหลายคอนเทนเนอร์
├── init.sql                 # สคริปต์เริ่มต้นฐานข้อมูล
├── go.mod                   # การอ้างอิงโมดูล Go
└── README.md                # README ภาษาอังกฤษ
```

## เว็บเฟรมเวิร์ก

โปรเจกต์นี้ใช้ [Fiber](https://gofiber.io/) ซึ่งเป็นเว็บเฟรมเวิร์กที่ได้รับแรงบันดาลใจจาก Express สร้างบน Fasthttp ซึ่งเป็นเอนจิน HTTP ที่เร็วมากสำหรับ Go ออกแบบมาเพื่อการพัฒนาอย่างรวดเร็ว ใช้หน่วยความจำต่ำ และเน้นประสิทธิภาพ

### ทำไมต้อง Fiber?

- เร็ว: สร้างบน Fasthttp ซึ่งมีประสิทธิภาพสูง
- คล้าย Express: คุ้นเคยสำหรับผู้ที่มาจากโลก Node.js
- เบา: โอเวอร์เฮดน้อยและใช้หน่วยความจำต่ำ
- มิดเดิลแวร์ครบ: รองรับความสามารถ HTTP ที่พบบ่อย
- ทดสอบง่าย: มีเครื่องมือช่วยทดสอบฮ্যান্ডเลอร์ HTTP

## API Endpoints

| เมธอด | เอ็นด์พอยต์              | คำอธิบาย                |
|-------|---------------------------|--------------------------|
| GET   | /api/v1/entities          | ดึงเอนทิตีทั้งหมด       |
| GET   | /api/v1/entities/{id}     | ดึงเอนทิตีตาม ID        |
| POST  | /api/v1/entities          | สร้างเอนทิตีใหม่        |
| PUT   | /api/v1/entities/{id}     | อัปเดตเอนทิตีตาม ID     |
| DELETE| /api/v1/entities/{id}     | ลบเอนทิตีตาม ID         |
| GET   | /swagger/*                 | Swagger UI               |
| GET   | /health                   | ตรวจสอบสถานะระบบ        |

## เริ่มต้นใช้งาน

### ข้อกำหนดเบื้องต้น

- Go เวอร์ชัน 1.19 ขึ้นไป
- Docker และ Docker Compose
- PostgreSQL (หากรันโดยไม่ใช้ Docker)

### รันด้วย Docker

1. สร้างและเริ่มบริการ:
   ```bash
   docker-compose up --build
   ```

2. API จะพร้อมใช้งานที่ `http://localhost:8080`

### รันบนเครื่อง (Local)

1. ติดตั้ง dependencies:
   ```bash
   go mod download
   ```

2. ตั้งค่าตัวแปรสภาพแวดล้อม:
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=postgres
   export DB_NAME=learnapi
   ```

3. รันแอปพลิเคชัน:
   ```bash
   go run cmd/api/main.go
   ```

## เอกสาร API

โปรเจกต์นี้จัดทำเอกสารด้วย Swagger หลังจากเริ่มแอปพลิเคชันแล้ว สามารถเปิด Swagger UI ได้ที่:
- `http://localhost:8080/swagger/index.html`

## การทดสอบ

รันการทดสอบทั้งหมดพร้อมเก็บ Coverage ครอบคลุมแพ็กเกจแอป:
```bash
go test -short -v \
  -covermode=atomic \
  -coverpkg=./internal/...,./pkg/...,./cmd/... \
  ./... \
  -coverprofile=coverage.txt
```

สรุป Coverage:
```bash
go tool cover -func=coverage.txt
```

ดูรายงานแบบ HTML:
```bash
go tool cover -html=coverage.txt
```

การทดสอบ Repository ต้องใช้ PostgreSQL วิธีที่แนะนำ:
- ใช้ Docker Compose: `docker-compose up -d db`
- หรือกำหนดตัวแปรแวดล้อมสำหรับฐานข้อมูลโลคอล:
  - `TEST_DB_HOST=127.0.0.1`
  - `TEST_DB_PORT=5432`
  - `TEST_DB_USER=postgres`
  - `TEST_DB_PASSWORD=postgres`
  - `TEST_DB_NAME=learnapi_test`

การทดสอบ E2E ต้องให้ API รันที่ `http://localhost:8080` หากไม่ต้องการรัน E2E ให้ตั้งค่า `SKIP_E2E_TESTS=true` ก่อนรันทดสอบ

### ความครอบคลุมของโค้ด (Code Coverage)

CI จะอัปโหลด Coverage ไปยัง Codecov โดยใช้ไฟล์ `codecov.yml` เพื่อไม่รวมไฟล์เอกสาร ม็อก และโค้ดทดสอบออกจากตัวหาร Coverage (badge แสดงไว้ด้านบน)

ดูรายละเอียดเพิ่มเติมเกี่ยวกับ E2E ได้ที่ `tests/e2e/README.md`

## โครงสร้างฐานข้อมูล

แอปพลิเคชันนี้ใช้ตาราง `entities` แบบง่าย ๆ:

```sql
CREATE TABLE entities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
