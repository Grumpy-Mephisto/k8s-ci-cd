<div align="center">
    <h1><code>💻</code> Operating Systems Project</h1>
    <h3>📚 A project for the Operating Systems course at the King Mongkut's University of Technology Ladkrabang.</h3>
</div>

## `⭐` Introduction

โปรเจคนี้เป็นส่วนหนึ่งของวิชา **Operating System** ที่ให้สร้าง Container ของ Docker ขึ้นมา และอธิบายการทำงานของ Docker และ Container ต่างๆ ที่เราสร้างขึ้นมา

## `📝` Requirements

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Kubernetes](https://kubernetes.io/)

## `📚` How to use

1. โคลน Repository นี้ลงมา

```bash
git clone git@github.com:PunGrumpy/k8s-ci-cd.git
```

2. สั่งรัน Application

   2.1 สำหรับ Docker

   ```bash
   cd deployment
   docker compose up -d
   ```

   2.2 สำหรับ Kubernetes

   ```bash
   kubectl apply -k k8s/
   ```

3. เข้าไปที่ [http://localhost:3000](http://localhost:3000) เพื่อเข้าสู่เว็บไซต์

### `🔮` Optional

- สามารถเข้าไปที่ [http://localhost:9021](http://localhost:9021) เพื่อเข้าสู่ Control Center ของ Kafka ได้
- สามารถเข้าไปที่ [http://localhost:8083](http://localhost:8083) เพื่อเข้าสู่ Debezium ได้

## `📝` Description

### `🐳` Docker

ใช้ Docker ในการสร้าง Image ของ Web Server โดย Web Server เราใช้ **Go** ในการทำ Web Server และใช้ **Fiber** เป็น Web Framework ในการทำ Web Server

### `🏭` Container

#### `🐳` Docker Compose

ใช้ Docker Compose ในการสร้าง Container ของ Web Server และ Database ขึ้นมา โดยใช้ Docker Compose ในการสร้าง Container ขึ้นมา 5 ตัว ได้แก่ Web Server, Kafka, Zookeeper, Control Center และ Debezium (connector) โดยที่ Web Server จะเป็น Container ที่เราสร้างขึ้นมาเอง ส่วน Kafka, Zookeeper, Control Center และ Debezium จะเป็น Container ที่เราดึงมาจาก Docker Hub

#### `⚓` Kubernetes

ใช้ Kubernetes ในการสร้าง Container ของ Web Server และ Database ขึ้นมา โดยใช้ Kubernetes ในการสร้าง Container ขึ้นมา 5 ตัว ได้แก่ Web Server, Kafka, Zookeeper, Control Center และ Debezium (connector) โดยที่ Web Server จะเป็น Container ที่เราสร้างขึ้นมาเอง ส่วน Kafka, Zookeeper, Control Center และ Debezium จะเป็น Container ที่เราใช้ Image จาก Docker Hub

### `📇` GitOps

#### `🔍` GitHub Actions

ใช้ GitHub Actions ในการทำ CI โดยที่เมื่อเราทำการ Test และ Build แล้ว จะทำการ Push Image ขึ้นไปบน Docker Hub และทำการอัปเดต Image บน Kubernetes โดยอัตโนมัติ และทำการ CD โดยที่เมื่อเราทำการ Deploy ไปที่ Remote Server (ลองแล้วทั้ง Kubernetes และ Docker Compose ทั้งสองตัว)

> **Note**: แต่ที่ Disable `cd.yml` เพราะ Remote Server ของเราขนาดไม่เพียงพอ จึงแนะนำให้ run ใน Local แทน

#### `🔍` ArgoCD

ใช้ ArgoCD ในการทำ CD โดยที่เมื่อเราทำการ Push Image ขึ้นไปบน Docker Hub แล้ว จะทำการอัปเดต Image บน Kubernetes โดยอัตโนมัติ

> **Note**: เราไม่ได้ใช้เพราะเนื่องจาก Remote Server ของเราขนาดไม่เพียงพอ จึงแนะนำให้ run ใน Local แทน
