# Sprint 3 Report


## Visual Demo Links
- [DEMO](https://shorturl.at/3hXCB)



## User Stories:
- Created 4 user stories spanning the features/user requirements of the Skill Arcade application. 
- Additional user stories will be introduced in the Later sprints to address emerging requirements, refine existing features, and ensure better alignment with project goals and user needs.
- These user stories can be found under the issues tab with the label [userstories](https://github.com/NavyaDurgam98/SkillArcade/issues) in this repository.



## Detailed Devlopment

# Backend Development

## UserScore Collection
A new collection named `UserScore` has been created to store user quiz data. Each document in this collection contains the following fields:

- **`user_id`**: Unique identifier for the user.
- **`quizzes`**: An array of quiz objects, where each object contains:
  - `quiz_topic_id`: Unique identifier for the quiz topic.
  - `quiz_topic_name`: Name of the quiz topic.
  - `score`: The score obtained by the user in the quiz.
  - `attempts`: The number of attempts made by the user for that quiz.
  - `submitted_at`: The timestamp when the quiz was submitted.
- **`total_score`**: The cumulative score obtained by the user across all quizzes.

### Sample Document
```json
{
  "_id": "67e6f5a5f2301ff2ed918420",
  "user_id": "679d5a260264697ca72d7c4a",
  "quizzes": [
    {
      "quiz_topic_id": "67c5fd91b35fea672a80a3e2",
      "quiz_topic_name": "Java",
      "score": 6,
      "attempts": 2,
      "submitted_at": "2025-03-28T19:17:29.080+00:00"
    },
    {
      "quiz_topic_id": "67c5fd91b35fea672a80a3e1",
      "quiz_topic_name": "C++",
      "score": 5,
      "attempts": 2,
      "submitted_at": "2025-03-28T19:24:53.009+00:00"
    }
  ],
  "total_score": 11
}
```

---
## Leaderboard API

### 1. Get User Rank and Details
**Endpoint:** `GET /leaderboard?user_id=<user_id>`
- Fetches the rank, quizzes taken, score, and attempts of a specific user.

**Example Request:**
```plaintext
GET /leaderboard?user_id=679d5a260264697ca72d7c4a
```

**Example Response:**
```json
{
  "rank": 5,
  "username": "JohnDoe",
  "total_score": 11,
  "quizzes_taken": 2
}
```

### 2. Get Top 10 Leaderboard
**Endpoint:** `GET /leaderboard`
- Retrieves the top 10 users with the highest scores, including their name, rank, score, and number of quizzes attempted.

**Example Request:**
```plaintext
GET /leaderboard
```

**Example Response:**
```json
[
  {"rank": 1, "username": "Alice", "total_score": 45, "quizzes_taken": 10},
  {"rank": 2, "username": "Bob", "total_score": 42, "quizzes_taken": 9}
]
```

---
## Submit Quiz API
**Endpoint:** `POST /submitquiz`
- Submits quiz results to the `UserScore` collection.
- Creates a new document if the user does not exist.
- Updates the existing document if the user exists.
- If the quiz already exists, increases the attempt count and updates the score only if the new score is higher.

**Example Request:**
```json
POST /submitquiz
{
  "user_id": "679d5a260264697ca72d7c4a",
  "quiz_topic_id": "67c5fd91b35fea672a80a3e2",
  "score": 7
}
```

**Example Response:**
```json
{
  "message": "Quiz submitted successfully",
  "updated": true
}
```

---
## Search API
**Endpoint:** `GET /search?q=<query>`
- Searches for categories, subcategories, or quiz topics based on a search string of 3 or more characters.

**Example Request:**
```plaintext
GET /search?q=Java
```

**Example Response:**
```json
[
  {"type": "quiz_topic", "name": "Java Basics"},
  {"type": "subcategory", "name": "Java OOP Concepts"}
]
```

---
## User History API
**Endpoint:** `GET /userhistory?user_id=<user_id>`
- Fetches all completed quizzes for a given user, including quiz ID, name, score, attempts, and submission time.

**Example Request:**
```plaintext
GET /userhistory?user_id=679d5a260264697ca72d7c4a
```

**Example Response:**
```json
[
  {
    "quiz_id": "67c5fd91b35fea672a80a3e2",
    "name": "Java",
    "score": 6,
    "attempts": 2,
    "submitted_at": "2025-03-28T19:17:29.080+00:00"
  }
]
```

---
## User Profile API
**Endpoint:** `GET /userprofile?user_id=<user_id>`
- Retrieves basic user information from the `UserDetails` table, including name, email, username, first name, and last name.

**Example Request:**
```plaintext
GET /userprofile?user_id=679d5a260264697ca72d7c4a
```

**Example Response:**
```json
{
  "username": "JohnDoe",
  "email": "john.doe@example.com",
  "first_name": "John",
  "last_name": "Doe"
}
```

### Backend Testing([[List of BackEnd Unit Tests](https://github.com/NavyaDurgam98/SkillArcade/issues/86)](https://github.com/NavyaDurgam98/SkillArcade/issues/86))


### Frontend Development


## Features  

### 1. **User Profile Component**  
The **User Profile Component** includes the following features:  
 
-  **Read-Only User Details**: First Name, Last Name, and Email fields are displayed as **non-editable** to ensure data integrity.  

### 2. **Navbar Component**  
- Implemented a **navigation bar** to provide seamless access to different sections of the application.
#### **About Popup**  
-  Clicking on the **"About"** section opens a popup with relevant information about **SkillArcade**.  
  - The popup can be closed by:  
  - Clicking the **"X"** button.  
  - Clicking **outside** the popup area.

####  **Profile Icon & Dropdown**  
- A **profile icon** is displayed in the dashboard header.  
- Clicking on the **profile icon** opens a dropdown with two options:  
  - **User Profile** – Navigates to the user's profile page.  
  - **Logout/Signout** – Logs the user out of the application.  

### 3. **Leaderboard Component**  
The **Leaderboard Component** helps users track their ranking and performance:  

- **User Rank Display:**  
  - **"You attempted:"** Displays the number of quiz attempts made by the user.  
  - **"Your Rank:"** Shows the user’s current ranking.  

- **Leaderboard Grid:**  
  - Displays rankings of all users in a **sortable table** with the following columns:  
    1. **Serial Number** (Auto-incremented)  
    2. **User Name**  
    3. **Rank**  

- **Sorting Options:**  
  - The **User Name** and **Rank** columns are **sortable** in ascending order to allow users to analyze ranking trends easily.

### 4. **Search**
The application includes a search feature that allows users to find relevant content efficiently. Users can search by entering a minimum of *three (3) characters*.

#### Searchable Entities
Users can search for the following entities:
- *Categories*  
- *Subcategories*  
- *Quizzes*  

#### Usage
1. Navigate to the search bar in the application.  
2. Enter at least *three (3) characters* to trigger the search.  
3. The results will dynamically display matching *categories, subcategories, and quizzes* based on the entered text and current page.






### Frontend Testing:([[List of Frontend Unit Tests](https://github.com/NavyaDurgam98/SkillArcade/issues/84)](https://github.com/NavyaDurgam98/SkillArcade/issues/84))
- **Jasmine & Karma Unit Tests**:
  - Developed test cases for **Navigation Bar Component, Leader Board Component, User Profile Component**.
  - Ensured each component met expected functionality and behavior.















