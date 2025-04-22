# Sprint 4 Report


# Visual Demo

[Watch the demo video](https://tinyurl.com/mw56549p)



# Detailed Development

## Backend Development
### 1. **AUTHENTICATION & ROUTE PROTECTION**  
- This application uses JWT (JSON Web Token) authentication to ensure that only authenticated users can access secure backend APIs and restricted pages in the Angular frontend.

### **Backend Protection**:
- JWTs are generated upon successful login and must be included in the Authorization header (as a Bearer token) for all /api/* routes. The Go backend verifies the token and rejects unauthorized access.

### **Frontend Route Guards**:
- Angular’s AuthGuard is used with canActivate to protect application routes like /dashboard, /profile, /categories, etc. Unauthenticated users are automatically redirected to the login page.

### **Token Handling**:
- Tokens are stored securely in localStorage on the client.
- An Angular HTTP interceptor automatically attaches the token to all outgoing API requests.


 ### **Logout Flow**:
- Logging out removes the token from local storage and blocks further access until the user logs in again.


## Frontend Development


## Features  

### 1. **Signup Component**  
- This Signup Component which has various fileds like FirstName, LastName, Email, Date Of Birth, Gender, Password and Confirm Password.
- We have done an authentication for Password and Confirm Password making sure that they match for sucessful signup.

### 2. **TakeQuiz Component**  
- We have added a left navigation panel to the takequiz component where it shows the current category, subcategories and the list of quiz topics.
- Whenever a user selects a quiz topics the remainig topics are disabled until the current quiz ends.
- We have implemented a timer based quiz which is for 100 seconds. A pop up gets displayed at the last 60 seconds and as soon as the quiz gets submitted the timer is stopped.
- When a user clicks on retry or starts a new quiz the timer starts again from 100 seconds.
 
### 3.**Search Bar**  
- UI Distortion is resolved. Previously if a user is searching for a technology UI was getting distorted whenever a search API was hit.
- Updated CSS for the search API endpoint for all 3 components (categories, subcategories, quiztopics).
 
### 4. **Password Hashing and Encoding:**
- User Password are securely encrypted using Bcrypt package.
- Encrypted password is stored in the mongodb document.
- Whenever a user tries to login with username and password , the input password is matched with the decryoted password in the mongodb document.

# Frontend Tests

### 1. **Unit Tests**

#### **LeaderboardComponent**
- should display leaderboard data in the table  
- should sort leaderboard data correctly  
- should toggle sorting order on column click  
- should create the component  
- should load leaderboard data on init  
- should handle API error gracefully  

#### **DashboardComponent**
- should call goToCategory with correct parameter  
- should display categories correctly  

#### **NavbarComponent**
- should call logout() and navigate to /login  
- should create the NavbarComponent  
- should call openAbout() and show the modal  

#### **CategoryComponent**
- should call goToQuizTopic() with correct subcategory when button is clicked  
- should create the category component  
- should display the correct number of subcategory cards  
- should display correct subcategory names  
- should display no subcategories if the API returns an empty array  

#### **LoginComponent**
- should initialize the form with empty fields  
- should show "Login" title initially  
- should validate required password field  
- should create the component  
- should toggle to Forgot Password view  
- should show error message if form submission fails  
- should enable submit button if form is valid  
- should validate required username field  
- should call onSubmit() when form is submitted  
- should navigate to signup page when clicking "Sign up here"  

#### **UserProfileComponent**
- should initialize with default profile  
- should load user profile successfully  
- should handle profile load error  
- should create  
- should use default values when stats are missing  
- should handle stats load error  
- should handle missing user ID  

#### **QuizComponent**
- should display correct quiz topic names  
- should display no quiz topics if the API returns an empty array  
- should create the quiz component  
- should call takeQuiz() when the button is clicked  
- should display the correct number of quiz topic cards  

#### **TakequizService**
- should submit quiz results  
- should handle errors when fetching quiz data  
- should be created  
- should fetch quiz data for a given topic  
- should fetch quiz topics for a category and subcategory  

#### **TakeQuizComponent**
- should select an answer  
- should increment question index on next question  
- should check and increase score if correct answer is selected  
- should create the component  
- should retrieve route parameters on initialization  
- should load quiz data from service  
- should reset the quiz on retry  
- should reset quiz data when closing modal  
- should navigate to a new quiz topic when selecting a topic  
- should return correct progress width  
- should detect last question correctly  

#### **SignupComponent**
- should initialize the form with empty fields  
- should display error message on signup failure  
- should validate email field  
- should call onSubmit() when form is submitted  
- should create the component  
- should validate required fields  
- should display success message on successful signup  
- should validate password minimum length  
- should enable submit button when form is valid  


### **Cypress End-to-End (E2E) Tests Documentation**

## 1. SignUp Page
- Submits valid form successfully
- Shows errors on empty form submission
- Rejects invalid email formats
- Rejects password shorter than 6 characters
- Shows mismatch error for password/confirmPassword
- Accepts long but valid names and usernames
- Should not allow future DOB (HTML native check)
- Displays error on server failure
- Does not show touched error until field is interacted with

## 2. Login Page
- Should display login form correctly
- Should show validation errors for empty fields
- Should log in successfully with valid credentials
- Should show error message on invalid login
- Should switch to forgot password mode
- Should validate email field in forgot password
- Should send forgot password request successfully
- Should show error if forgot password request fails
- Should navigate to signup page when clicking on "New user? Sign up here"

## 3. Dashboard Page
- Should display the correct number of category cards
- Should display correct category titles on each card
- Should display an image for each category card
- Should have a Sub Topics button on every card
- Should navigate to correct category page when Sub Topics is clicked
- Should not show About modal by default
- Should open About modal when openAbout is triggered manually
- Should show cards in a responsive layout (basic check)
- Should show correct image paths in `img src` attributes

## 4. Category Page
- Should call API and display filtered subcategories when searchText = "comp"
- Should display correct category titles from mock data
- Should display correct images for each category
- Should navigate to quiz topic page on clicking a subcategory card
- Should display "No results found" when no match is returned

## 5. Navbar Component
- Should render the logo, search input, About button, and profile dropdown
- Should allow typing into the search box and trigger the input event
- Should open the About modal when About button is clicked
- Should navigate to User Profile when dropdown item is clicked
- Should navigate to Leader Board when dropdown item is clicked
- Should remove tokens and redirect to login on logout

## 6. TakeQuiz Page
- Should display the initial welcome message
- Should show the topic selection sidebar
- Should start a quiz when clicking on a topic
- Should allow navigating through quiz questions
- Should display the completion modal after answering the last questions
- Should show score in the completion modal
- Should show an error message if quiz data is not available
- Should allow users to view question options
- Should navigate to the results page after completing the quiz
- Should show a confirmation message if the quiz is resumed
- Should navigate back to quiz topics from the quiz















# Backend Unit Tests

### 1. **LeaderboardServiceTestSuite**
- TestGetLeaderboardService
  - success_top10
  - success_user_rank
  - user_not_found
  - invalid_user_id_format
  - aggregate_error

### 2. **FetchCategories**
- success_with_results
- success_no_results
- case_insensitive_search
- database_error
- decoding_error

### 3. **UserLoginService**
- user_not_found
- invalid_password
- invalid_user_id_format

### 4. **QuizQuestionServiceTestSuite**
- TestFetchQuizQuestions
  - success
  - not_found
  - database_error

### 5. **FetchQuizTopics**
- success_with_results
- success_no_results
- search_with_filtered_results
- subcategory_not_found
- database_error

### 6. **UserRegistrationService**
- success
- email_already_exists
- failed_to_hash_password
- insert_error

### 7. **FetchSubCategories**
- success_with_results
- success_no_results
- search_with_filtered_results
- category_not_found
- database_error

### 8. **SubmitQuizServiceTestSuite**
- TestSubmitQuizService
  - success_new_user_score
  - success_update_existing_quiz
  - invalid_user_id_format
  - invalid_quiz_topic_id_format
  - insert_failure
  - update_failure

### 9. **UserProfileServiceTestSuite**
- TestGetUserProfileService
  - success
  - user_not_found
  - invalid_user_id_format

### 10. **UserHistoryServiceTestSuite**
- TestGetUserHistoryService
  - success_user_history_found
  - invalid_user_id_format
  - user_not_found_or_no_history




  
# Backend ApI Documentation

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
  "username": "SKonduru",
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
```plaintext
POST /submitquiz
```
```json
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
**Endpoint:** `GET /Component?searchText=<text>`
- Searches for categories, subcategories, or quiz topics based on a search string of 3 or more characters.

**Example Request:**
```plaintext
GET /categories?searchText=compu
```

**Example Response:**
```json
[
    {
        "category": "Computer Science",
        "imgPath": "imgCS1.png"
    },
    {
        "category": "Cloud Computing",
        "imgPath": "img1.png"
    }
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
  "username": "SKonduru",
  "email": "konduru.s@example.com",
  "first_name": "Suppi",
  "last_name": "Konduru"
}
```















