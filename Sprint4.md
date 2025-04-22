# Sprint 4 Report


## Visual Demo Links
- 


## Detailed Devlopment

# Backend Development
### 1. **AUTHENTICATION & ROUTE PROTECTION**  
-This application uses JWT (JSON Web Token) authentication to ensure that only authenticated users can access secure backend APIs and restricted pages in the Angular frontend.

### **Backend Protection**:
 -JWTs are generated upon successful login and must be included in the Authorization header (as a Bearer token) for all /api/* routes. The Go backend verifies the token and rejects unauthorized access.

### **Frontend Route Guards**:
 -Angular’s AuthGuard is used with canActivate to protect application routes like /dashboard, /profile, /categories, etc. Unauthenticated users are automatically redirected to the login page.

### **Token Handling**:
 -Tokens are stored securely in localStorage on the client.
 -An Angular HTTP interceptor automatically attaches the token to all outgoing API requests.


 ### **Logout Flow**:
 -Logging out removes the token from local storage and blocks further access until the user logs in again.


















### Backend Testing ([List of BackEnd Unit Tests](https://github.com/NavyaDurgam98/SkillArcade/issues/86))




### Frontend Development


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
  

### Frontend Testing:([[List of Frontend Unit Tests](https://github.com/NavyaDurgam98/SkillArcade/issues/84)](https://github.com/NavyaDurgam98/SkillArcade/issues/84))















