describe('TakeQuiz Page', () => {
    beforeEach(() => {
      cy.visit('/takequiz'); // Ensure you're on the takequiz page before each test
    });
  
    it('should display the initial welcome message', () => {
      cy.get('.quiz-content h2').should('contain', 'Welcome to the Quizzes');
    });
  
    it('should show the topic selection sidebar', () => {
      cy.get('.quiz-sidebar').should('be.visible');
      cy.get('.quiz-sidebar h3').should('exist');
      cy.get('.quiz-sidebar li').should('have.length.gt', 0); // Ensure there are quiz topics
    });
  
    it('should start a quiz when clicking on a topic', () => {
      cy.get('.quiz-sidebar li').first().click(); // Click the first topic
      cy.get('.quiz-content h2').should('contain', 'Welcome to the Quiz of');
      cy.get('.start-quiz-btn').click(); // Click to start the quiz
      cy.get('.quiz-timer').should('be.visible'); // Verify timer shows up
      cy.get('.question-text').should('be.visible'); // Verify first question is displayed
    });
  
    it('should allow navigating through quiz questions', () => {
      cy.get('.quiz-sidebar li').first().click();
      cy.get('.start-quiz-btn').click();
      cy.get('.option-btn').first().click(); // Select first option for the question
      cy.get('.next-btn').click(); // Go to next question
      cy.get('.quiz-progress span').should('contain', 'Question 2/'); // Check progress
    });
  
    it('should display the completion modal after answering the last question', () => {
      cy.get('.quiz-sidebar li').first().click();
      cy.get('.start-quiz-btn').click();
      cy.get('.option-btn').first().click();
      cy.get('.next-btn').click(); // Answer first question
      cy.get('.option-btn').first().click(); // Answer second question
      cy.get('.submit-btn').click(); // Submit the quiz
      cy.get('.modal-content').should('be.visible'); // Verify modal appears upon completion
    });
  
    it('should show score in the completion modal', () => {
      cy.get('.quiz-sidebar li').first().click();
      cy.get('.start-quiz-btn').click();
      cy.get('.option-btn').first().click();
      cy.get('.next-btn').click(); // Answer first question
      cy.get('.option-btn').first().click(); // Answer second question
      cy.get('.submit-btn').click(); // Submit the quiz
      cy.get('.modal-content .score').should('exist'); // Verify score is displayed
    });
  
    it('should show an error message if quiz data is not available', () => {
      cy.intercept('GET', '**/quizdata/*', { statusCode: 404 }).as('quizDataRequest'); // Mock 404 error
      cy.get('.quiz-sidebar li').first().click();
      cy.get('.start-quiz-btn').click();
      cy.wait('@quizDataRequest');
      cy.get('.error-message').should('contain', 'Unable to load quiz data. Please try again later.');
    });
  
    it('should allow users to view question options', () => {
      cy.get('.quiz-sidebar li').first().click();
      cy.get('.start-quiz-btn').click();
      cy.get('.question-text').should('be.visible');
      cy.get('.option-btn').should('have.length.greaterThan', 1); // Verify multiple options exist for a question
    });
  
    it('should navigate to the results page after completing the quiz', () => {
      cy.get('.quiz-sidebar li').first().click();
      cy.get('.start-quiz-btn').click();
      cy.get('.option-btn').first().click(); // Answer the first question
      cy.get('.next-btn').click(); // Answer the next question
      cy.get('.submit-btn').click(); // Submit quiz
      cy.url().should('include', '/results'); // Verify the user is redirected to results page
    });

    it('should show a confirmation message if the quiz is resumed', () => {
      // Assuming a user can resume a quiz
      cy.get('.quiz-sidebar li').first().click();
      cy.get('.resume-quiz-btn').click(); // Click to resume quiz
      cy.get('.confirmation-modal').should('be.visible'); // Confirm resume modal appears
      cy.get('.confirmation-modal .message').should('contain', 'Are you sure you want to resume the quiz?');
    });
  
    it('should navigate back to quiz topics from the quiz', () => {
      cy.get('.quiz-sidebar li').first().click();
      cy.get('.start-quiz-btn').click();
      cy.get('.back-to-topics-btn').click(); // Click to go back to quiz topics
      cy.url().should('include', '/takequiz'); // Verify we're back to the quiz topics page
    });
  });
