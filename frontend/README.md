# Go Next

> Short description of what this project does.


- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [Available Scripts](#available-scripts)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Contributing](#contributing)

---

## Tech Stack

**Frontend:**
- [Next.js 14](https://nextjs.org/) - React Framework
- [TypeScript](https://www.typescriptlang.org/) - Type Safety
- [Tailwind CSS](https://tailwindcss.com/) - Styling

**Backend:**
- [Go](https://golang.org/) - REST API
- [PostgreSQL](https://www.postgresql.org/) - Database
- [GORM](https://gorm.io/) - ORM

---

## Prerequisites

Make sure you have the following installed:

- [Node.js](https://nodejs.org/) >= 18.x
- [npm](https://www.npmjs.com/) or [yarn](https://yarnpkg.com/)
- [Go](https://golang.org/) >= 1.21

---


### 1. Install dependencies

```bash
cd frontend
npm install
```

### 2. Setup environment variables

```bash
cp .env.example .env.local
```

Fill in the required values in `.env.local` (see [Environment Variables](#environment-variables)).

### 3. Run the development server

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) in your browser.

---

## 4. Environment Variables

Create a `.env.local` file in the `frontend` directory:

```env
COMING SOON
```

---


### 5. Docker

```bash
docker build -t frontend .
docker run -p 3000:3000 frontend
```

---

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## License

This project is licensed under the [MIT License](LICENSE).