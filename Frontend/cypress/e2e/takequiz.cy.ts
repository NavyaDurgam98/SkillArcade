// describe('Takequiz Component', () => {
//     beforeEach(() => {
//       // Visit the takequiz page for a given category, subcategory, and topic
//       cy.visit('/takequiz/Science/Physics/Gravity');
      
//       // Ensure quiz topics are loaded before moving on
//       cy.get('.quiz-sidebar h3', { timeout: 10000 }).should('exist'); // Wait for the sidebar h3 (category > subcategory)
//       cy.get('.quiz-sidebar ul li', { timeout: 10000 }).should('have.length.greaterThan', 0); // Wait for the quiz topics to be listed
//     });
  
//     it('should load the quiz topics list', () => {
//       cy.get('.quiz-sidebar h3').should('contain', 'Science > Physics'); // Ensure the correct category and subcategory are displayed
//       cy.get('.quiz-sidebar ul li').should('have.length.greaterThan', 0); // Ensure quiz topics are listed
//     });
  
//     it('should start the quiz when clicking the start button', () => {
//       // Wait for the start button to be visible
//       cy.get('.start-quiz-btn', { timeout: 10000 }).should('be.visible'); 
//       cy.get('.start-quiz-btn').click(); // Click the start button
  
//       // Verify that the quiz has started
//       cy.get('.quiz-timer').should('exist'); // Timer should be visible
//       cy.get('.quiz-progress').should('exist'); // Progress bar should be visible
//       cy.get('.question-text').should('exist'); // Question text should be visible
//     });
  
//     it('should allow answering a question and move to the next one', () => {
//       cy.get('.start-quiz-btn').click(); // Start quiz
  
//       // Wait for the question to load
//       cy.get('.question-text').should('exist'); 
  
//       // Select an answer (first option)
//       cy.get('.option-btn').first().click();
//       cy.get('.next-btn').should('not.be.disabled'); // Next button should be enabled after selecting an answer
  
//       // Click next to move to the next question
//       cy.get('.next-btn').click();
  
//       // Verify the next question is displayed
//       cy.get('.question-text').should('exist');
//     });
  
//     it('should allow submitting the quiz', () => {
//       cy.get('.start-quiz-btn').click(); // Start quiz
  
//       // Wait for the question to load
//       cy.get('.question-text').should('exist');
  
//       // Select an answer
//       cy.get('.option-btn').first().click();
  
//       // Submit the quiz on the last question
//       cy.get('.submit-btn').click();
  
//       // Verify the modal with the score is shown
//       cy.get('.modal-overlay').should('exist');
//       cy.get('.modal-content h2').should('contain', 'Quiz Completed!');
//     });
  
//     it('should show time warning when only 1 minute is left', () => {
//       cy.get('.start-quiz-btn').click(); // Start quiz
  
//       // Set up the timer to run and trigger time warning
//       cy.window().then((win) => {
//         cy.stub(win, 'setInterval').callsFake((callback) => {
//           callback(); // Trigger callback immediately for testing
//         });
//       });
  
//       // Wait for the warning to show
//       cy.get('.time-warning-modal').should('exist');
//       cy.get('.time-warning-modal p').should('contain', '⚠️ Only 1 minute left!');
//     });
  
//     it('should retry the quiz after completion', () => {
//       cy.get('.start-quiz-btn').click(); // Start quiz
  
//       // Select an answer and submit
//       cy.get('.option-btn').first().click();
//       cy.get('.submit-btn').click();
  
//       // Verify the quiz modal appears and retry button is visible
//       cy.get('.modal-overlay').should('exist');
//       cy.get('.retry-btn').click(); // Click retry
  
//       // Verify that the quiz resets (ensure first question is displayed again)
//       cy.get('.question-text').should('exist');
//       cy.get('.option-btn').first().should('not.have.class', 'selected'); // Ensure no option is pre-selected
//     });
//   });
  