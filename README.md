# Food Recipe Project

This project is a full-stack application for managing food recipes. It consists of three main components: a Go-based backend service, a Nuxt.js frontend, and a Hasura GraphQL Engine with database configuration and migration.

[Watch the demo video]([https://example.com/path/to/your/video.mp4](https://drive.google.com/file/d/1B-POPAETXUxu2Q7-3VbxgEWHre55fA3a/view?usp=sharing))

---

## Project Structure

### 1. `service`
This folder contains the Go backend service.
- **Description**: Provides the API required for core business logic and supports the GraphQL engine dependency.
- **Usage**:
  ```bash
  cd service
  go run main.go
  ```

### 2. `frontend`
This folder contains the Nuxt.js frontend application.
- **Description**: Provides the user interface for interacting with the food recipe platform.
- **Usage**:
  ```bash
  cd frontend
  npm install
  npm run dev
  ```

### 3. `FoodRecipe`
This folder contains migration files and Hasura metadata.
- **Description**: Includes database migration scripts and configurations for Hasura GraphQL Engine.
- **Usage**:
  ```bash
  cd FoodRecipe
  sudo docker-compose up -d
  hasura migrate apply
  hasura metadata apply
  ```

---

## Running the Project

To run the project, follow these steps in order:

### 1. Start the Go Backend Service
The Go service must be running before starting the GraphQL engine due to dependency.
```bash
cd service
go run main.go
```

### 2. Start the Hasura GraphQL Engine and Database
Use `docker-compose` to start the Hasura GraphQL Engine along with the database configuration.
```bash
cd FoodRecipe
sudo docker-compose up -d
```

### 3. Start the Frontend Application
After starting the backend and database, launch the Nuxt.js frontend to provide the user interface.
```bash
cd frontend
npm install
npm run dev
```

---

## Notes
- Ensure the `service` is running **before** starting the GraphQL engine to avoid dependency issues.
- Make sure Docker and Docker Compose are installed and properly configured on your machine.
- You can access the Nuxt.js frontend at [http://localhost:3000](http://localhost:3000) after starting the application.
- The Hasura GraphQL Console will be available at [http://localhost:8080](http://localhost:8080) once the engine starts.

---

## Requirements
- **Go**: To run the backend service.
- **Node.js and npm**: To run the frontend application.
- **Docker & Docker Compose**: To manage the Hasura GraphQL Engine and database.

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## Contributions
Contributions are welcome! Feel free to open a pull request or file an issue in the repository.

