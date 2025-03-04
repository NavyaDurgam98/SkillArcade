# Sprint 2 Report

## Visual Demo Links
## Visual Demo Links

## Visual Demo Links
- [DEMO](https://tinyurl.com/mrykh238)


## User Stories:
- Created 4 user stories spanning the features/user requirements of the Skill Arcade application. 
- Additional user stories will be introduced in the Later sprints to address emerging requirements, refine existing features, and ensure better alignment with project goals and user needs.
- These user stories can be found under the issues tab with the label [userstories](https://github.com/NavyaDurgam98/SkillArcade/issues) in this repository.


## Detailed Devlopment
 
### Frontend Development  
- Implemented **Technology Panels** with the following features:
  - **Dashboard Component** - Displays all the **domains** in Computer Science Engineering.
  - **Technology Category Panel** - Fetches categories dynamically and renders them.
  - **Sub-Technology Category Panel** - Displays relevant **quiz topics** based on the selected category.
- Developed **Quiz Component** with the following functionalities:
  - Clicking on a **quiz topic** opens a page to **take the quiz**.
  - Implemented **quiz rendering** when clicking the **Start button**.
  - Built browser-side **validation** to check quiz answers.
  - **Redirected users** to quiz topics after quiz completion.
- **Data Integration**:
  - Integrated the backend API to fetch **categories and quiz topics**.
  - Ensured smooth **state management** for navigation.
 
### Frontend Testing:([[List of Frontend Unit Tests & cypress Tests](https://github.com/NavyaDurgam98/SkillArcade/issues/41)](https://github.com/NavyaDurgam98/SkillArcade/issues/41))
- **Jasmine & Karma Unit Tests**:
  - Developed test cases for **Login, Signup, Dashboard, Category, Quiz and TakeQuiz Components**.
  - Ensured each component met expected functionality and behavior.
- **Cypress E2E Testing**:
  - Conducted **end-to-end tests** for **Login**.
  - Validated **form submission, routing, error handling, and UI consistency**.







