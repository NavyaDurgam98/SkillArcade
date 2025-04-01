# Sprint 3 Report


## Visual Demo Links
- [DEMO](https://tinyurl.com/mrykh238)



## User Stories:
- Created 4 user stories spanning the features/user requirements of the Skill Arcade application. 
- Additional user stories will be introduced in the Later sprints to address emerging requirements, refine existing features, and ensure better alignment with project goals and user needs.
- These user stories can be found under the issues tab with the label [userstories](https://github.com/NavyaDurgam98/SkillArcade/issues) in this repository.



## Detailed Devlopment
### Backend Development  


### Backend Testing




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















