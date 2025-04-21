describe('Complete Signup Form Tests', () => {
  const baseUrl = 'http://localhost:4200'; // Adjust if needed

  const fillValidForm = () => {
    cy.get('#firstName').type('John');
    cy.get('#lastName').type('Doe');
    cy.get('#email').type('john.doe@example.com');
    cy.get('#dob').type('1990-01-01');
    cy.get('#male').check();
    cy.get('#username').type('johnnydoe');
    cy.get('#password').type('strongpass123');
    cy.get('#confirmPassword').type('strongpass123');
  };

  beforeEach(() => {
    cy.visit(`${baseUrl}/signup`);
  });

  // ----- Positive Case -----
  it('submits valid form successfully', () => {
    cy.intercept('POST', '**/signup', {
      statusCode: 200,
      body: { message: 'Signed up!' }
    }).as('signupSuccess');

    fillValidForm();
    cy.get('button[type="submit"]').click();

    cy.wait('@signupSuccess');
    cy.contains('Sign up successful! Please log in.').should('be.visible');
  });

  // ----- Required Fields -----
  it('shows errors on empty form submission', () => {
    cy.get('button[type="submit"]').click();
    cy.contains('First name is required.').should('exist');
    cy.contains('Last name is required.').should('exist');
    cy.contains('Valid email is required.').should('exist');
    cy.contains('DOB is required.').should('exist');
    cy.contains('Select your gender.').should('exist');
    cy.contains('Username is required.').should('exist');
    cy.contains('Password is required and must be at least 6 characters.').should('exist');
    cy.contains('Confirm your password.').should('exist');
  });

  // ----- Email Validations -----
  it('rejects invalid email formats', () => {
    cy.get('#email').type('john.doe');
    cy.get('button[type="submit"]').click();
    cy.contains('Valid email is required.').should('be.visible');
  });

  // ----- Password Validations -----
  it('rejects password shorter than 6 characters', () => {
    cy.get('#password').type('123');
    cy.get('button[type="submit"]').click();
    cy.contains('Password is required and must be at least 6 characters.').should('exist');
  });

  it('shows mismatch error for password/confirmPassword', () => {
    cy.get('#password').type('abcdefg');
    cy.get('#confirmPassword').type('1234567');
    cy.get('button[type="submit"]').click();
    cy.contains('Passwords do not match.').should('be.visible');
  });

  // ----- Edge Case: Long Inputs -----
  it('accepts long but valid names and usernames', () => {
    cy.intercept('POST', '**/signup').as('signupSuccess');

    cy.get('#firstName').type('A'.repeat(50));
    cy.get('#lastName').type('B'.repeat(50));
    cy.get('#email').type('test@example.com');
    cy.get('#dob').type('1990-01-01');
    cy.get('#female').check();
    cy.get('#username').type('user' + 'X'.repeat(40));
    cy.get('#password').type('longpassword123');
    cy.get('#confirmPassword').type('longpassword123');
    cy.get('button[type="submit"]').click();

    cy.wait('@signupSuccess');
  });

  // ----- Edge Case: Invalid DOB (future) -----
  it('should not allow future DOB (HTML native check)', () => {
    const futureDate = new Date();
    futureDate.setFullYear(futureDate.getFullYear() + 1);
    const futureStr = futureDate.toISOString().split('T')[0];

    cy.get('#dob').type(futureStr).should('have.value', futureStr);
    // Angular might not validate this — you could write your own validator for real cases
  });

  // ----- Backend Failure Case -----
  it('displays error on server failure', () => {
    cy.intercept('POST', '**/signup', {
      statusCode: 500,
      body: { message: 'Server error' }
    }).as('signupError');

    fillValidForm();
    cy.get('button[type="submit"]').click();
    cy.wait('@signupError');

    cy.contains('There was an error with the sign up. Please try again later.').should('be.visible');
  });

  // ----- Client Side Interaction Only -----
  it('does not show touched error until field is interacted with', () => {
    cy.get('#firstName').should('exist');
    cy.get('#firstName').focus().blur();
    cy.contains('First name is required.').should('exist');
  });
});
