# Sprint 2 Report


## Visual Demo Links
- [DEMO](https://tinyurl.com/mrykh238)


## User Stories:
- Created 4 user stories spanning the features/user requirements of the Skill Arcade application. 
- Additional user stories will be introduced in the Later sprints to address emerging requirements, refine existing features, and ensure better alignment with project goals and user needs.
- These user stories can be found under the issues tab with the label [userstories](https://github.com/NavyaDurgam98/SkillArcade/issues) in this repository.


## Detailed Devlopment
### Backend Development  
- **Database**:
  - Quizzes Collection - This collection includes categories of technologies, their sub-categories and Quiz topics under each sub category. Each quiz topic has quiz_topic_name and a unique quiz_topic_id.  
  - QuizQuestions Collection - This collection includes all the questions for a particular quiz topic.
  - QuizQuestions collection's _id corresponds to quiz_topic_id in Quizzes Collection.
  
Quizzes Collection  
```json
{
  "_id": {
    "$oid": "67c5fd94b35fea672a80a3ea"
  },
  "category": "Computer Science",
  "sub_categories": [
    {
      "sub_category": "Programming Languages",
      "quiz_topics": [
        {
          "quiz_topic_id": {
            "$oid": "67c5fd91b35fea672a80a3e1"
          },
          "quiz_topic_name": "C++"
        },
        {
          "quiz_topic_id": {
            "$oid": "67c5fd91b35fea672a80a3e2"
          },
          "quiz_topic_name": "Java"
        },
        {
          "quiz_topic_id": {
            "$oid": "67c5fd92b35fea672a80a3e3"
          },
          "quiz_topic_name": "Python"
        },
        {
          "quiz_topic_id": {
            "$oid": "67c5fd92b35fea672a80a3e4"
          },
          "quiz_topic_name": "C#"
        }
      ]
    },
    {
      "sub_category": "Data Structures and Algorithms",
      "quiz_topics": [
        {
          "quiz_topic_id": {
            "$oid": "67c5fd92b35fea672a80a3e5"
          },
          "quiz_topic_name": "Arrays & Strings"
        },
        {
          "quiz_topic_id": {
            "$oid": "67c5fd92b35fea672a80a3e6"
          },
          "quiz_topic_name": "Graphs"
        }
      ]
    }
  ]
}
```
QuizQuestions Collection
```json
{
  "_id": {
    "$oid": "67c5fd91b35fea672a80a3e1"
  },
  "quiz_topic": "C++",
  "questions": [
    {
      "question": "What is the output of 'std::cout << 5/2;' in C++?",
      "options": [
        { "option_id": { "$numberInt": "1" }, "value": "2" },
        { "option_id": { "$numberInt": "2" }, "value": "2.5" },
        { "option_id": { "$numberInt": "3" }, "value": "2" },
        { "option_id": { "$numberInt": "4" }, "value": "Compilation Error" }
      ],
      "correct_option": { "$numberInt": "1" }
    },
    {
      "question": "Which keyword is used to prevent a function from being overridden?",
      "options": [
        { "option_id": { "$numberInt": "1" }, "value": "static" },
        { "option_id": { "$numberInt": "2" }, "value": "const" },
        { "option_id": { "$numberInt": "3" }, "value": "final" },
        { "option_id": { "$numberInt": "4" }, "value": "volatile" }
      ],
      "correct_option": { "$numberInt": "3" }
    },
    {
      "question": "Which of the following is not a valid C++ access specifier?",
      "options": [
        { "option_id": { "$numberInt": "1" }, "value": "public" },
        { "option_id": { "$numberInt": "2" }, "value": "protected" },
        { "option_id": { "$numberInt": "3" }, "value": "private" },
        { "option_id": { "$numberInt": "4" }, "value": "internal" }
      ],
      "correct_option": { "$numberInt": "4" }
    },
    {
      "question": "What does the 'new' keyword do in C++?",
      "options": [
        { "option_id": { "$numberInt": "1" }, "value": "Allocates memory" },
        { "option_id": { "$numberInt": "2" }, "value": "Deallocates memory" },
        { "option_id": { "$numberInt": "3" }, "value": "Initializes a variable" },
        { "option_id": { "$numberInt": "4" }, "value": "None of the above" }
      ],
      "correct_option": { "$numberInt": "1" }
    },
    {
      "question": "Which of the following is a C++ STL container?",
      "options": [
        { "option_id": { "$numberInt": "1" }, "value": "Stack" },
        { "option_id": { "$numberInt": "2" }, "value": "Queue" },
        { "option_id": { "$numberInt": "3" }, "value": "Vector" },
        { "option_id": { "$numberInt": "4" }, "value": "All of the above" }
      ],
      "correct_option": { "$numberInt": "4" }
    }
  ]
}
```
- **Dataset**:
  -Quiz Dataset: This dataset used in this project was custom-built by us to ensure it aligns with the project's needs.
- **API**'s:
  Implemented API's to fetch quiz data like names of categories, sub-categories of each category, the quiz topics and the corresponding questions.

### Backend Testing
- **Testify & mtest**:
  - Used **testify** suite for structured testing and **mtest** for mocking mongoDB interactions.
  - The tests structure allows to thoroughly test database interactions , covering success and various failure scenarios, all without needing a real database.
  - Code includes a complete test suite covering all services to validate their functionality and performance. These tests help maintain code quality and prevent regressions.

## TestFetchQuizQuestions
-  **Success**: Fetches quiz questions successfully.
-  **not_found**: Handles case where quiz questions are not found.
-  **database_error**: Handles database failure scenario.

## TestFetchSubCategories
-  **Success**: Successfully retrieves subcategories.
-  **not_found**: Handles case where subcategories are not found.
-  **database_error**: Handles database failure scenario.

## TestFetchQuizTopics
-  **Success**: Successfully fetches quiz topics.
-  **not_found**: Handles case where quiz topics are not found.
-  **database_error**: Handles database failure scenario.

## TestFetchCategories
-  **Success**: Retrieves categories successfully.
-  **not_found**: Handles case where categories are not found.
-  **database_error**: Handles database failure scenario.


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







