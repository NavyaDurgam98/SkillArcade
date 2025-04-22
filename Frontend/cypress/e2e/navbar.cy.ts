describe('Navbar Component', () => {
    beforeEach(() => {
      // Visit the root or a route where the navbar is visible
      cy.visit('/');
    });
  
    it('should render the logo, search input, About button, and profile dropdown', () => {
      cy.get('.navbar-brand img').should('exist');
      cy.get('input[placeholder="Search..."]').should('exist');
      cy.contains('button', 'About').should('exist');
      cy.get('#profileDropdown').should('exist');
    });
  
    it('should allow typing into the search box and trigger the input event', () => {
      const searchText = 'test search';
  
      cy.get('input[placeholder="Search..."]')
        .type(searchText)
        .should('have.value', searchText);
    });
  
    it('should open the About modal when About button is clicked', () => {
      // Stub the modal or check if modal gets triggered (assuming modal exists)
      cy.get('button').contains('About').click();
      
      // Assuming modal has id #aboutModal and becomes visible
      cy.get('#aboutModal').should('be.visible');
    });
  
    it('should navigate to User Profile when dropdown item is clicked', () => {
      cy.get('#profileDropdown').click();
      cy.contains('a.dropdown-item', 'User Profile').click();
      cy.url().should('include', '/profile');
    });
  
    it('should navigate to Leader Board when dropdown item is clicked', () => {
      cy.get('#profileDropdown').click();
      cy.contains('a.dropdown-item', 'Leader Board').click();
      cy.url().should('include', '/leaderboard');
    });
  
    it('should remove tokens and redirect to login on logout', () => {
      localStorage.setItem('authToken', 'test-token');
      localStorage.setItem('userId', '12345');
  
      cy.get('#profileDropdown').click();
      cy.contains('a.dropdown-item', 'Logout').click();
  
      cy.url().should('include', '/login');
      cy.window().then((win) => {
        expect(win.localStorage.getItem('authToken')).to.be.null;
        expect(win.localStorage.getItem('userId')).to.be.null;
      });
    });
  });
  