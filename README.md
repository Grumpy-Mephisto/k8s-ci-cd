<div align="center">
    <h1><code>💻</code> Operating Systems Project</h1>
    <h3>📚 A project for the Operating Systems course at the King Mongkut's University of Technology Ladkrabang.</h3>
</div>

## `⭐` Introduction

โปรเจคนี้เป็นส่วนหนึ่งของวิชา **Operating System** ที่ให้สร้าง Container ของ Docker ขึ้นมา และอธิบายการทำงานของ Docker และ Container ต่างๆ ที่เราสร้างขึ้นมา

## `📝` Requirements

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## `📚` How to use

1. โคลน Repository นี้ลงมา

```bash
git clone
```

2. สั่งรัน Docker

   2.1 ใช้คำสั่ง `docker compose` ในการรัน

   ```bash
   cd web-server
   docker-compose up -d # docker compose up -d
   ```

   2.2 ใช้คำสั่ง `docker` ในการรัน

   ```bash
   cd web-server
   docker build -t web-server .
   docker run -d -p 3000:3000 web-server
   ```

3. Open your browser and go to `http://localhost:3000` (default port)
