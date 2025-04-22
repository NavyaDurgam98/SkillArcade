# SkillArcade : Dynamic Online Quiz Platform


**Project Description**  
SkillArcade is an web application designed to help students, job seekers, and professionals improve their technical skills through engaging quizzes. Users can take quizzes with time limits that adjust according to the difficulty level. The platform includes features like score based leaderboards, a rewards system to unlock advanced quizzes, and instant feedback to make learning competitive and interactive. Quizzes are grouped by technology to ensure a convenient experience for participants.

**FrontEnd Engineers**  
Navya Durgam - 33725389  
Lasya Sree Devabhaktuni - 15186666  

**BackEnd Engineers**  
Tejasri Baddam - 36778822  
Sri Charan Pabbathi - 80984724   

## STEPS TO EXECUTE THE PROJECT

### Project Structure
```
SkillArcade/
├── Backend/      → Go backend server
├── Frontend/     → Angular frontend
```

---

### Backend Setup (Go)
1. Navigate to the `Backend` directory:
   ```bash
   cd Backend
   ```
2. Initialize Go modules (only needed once):
   ```bash
   go mod init 
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Create `.env` file. Set up the required `.env` file with environment variables (e.g., database URL, ports, secrets).  
   **Example `.env` file structure:**
```
  PORT=8080
  MONGO_USER=your_mongodb_username
  MONGO_PASSWORD=your_mongodb_password
  MONGO_CLUSTER=your_mongodb_cluster_url
  MONGO_DBNAME=your_database_name
  JWT_SECRET=your_secret_key
  SENDGRID_API_KEY=your_sendgrid_api_key
  CORS_ORIGINS=http://localhost:4200, http://localhost:8080
```
   > *Make sure the `.env` file is placed in the `Backend` directory before running the server.*

5. Start the backend server:
   ```bash
   go run main.go
   ```

---

### Frontend Setup (Angular)
1. Navigate to the `Frontend` directory:
   ```bash
   cd Frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run the Angular app:
   ```bash
   ng serve
   ```

---

###  Test Login Credentials
Use the following test credentials to log in:
```json
{
  "username": "skonduru",
  "password": "konduru@123"
}
```


