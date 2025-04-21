
describe('Dashboard Page Tests', () => {
    beforeEach(() => {
      localStorage.setItem('authToken', 'mock-token');
  
      // Load fixture and intercept API call
      cy.fixture('categories.json').then((mockCategories) => {
        cy.intercept('GET', 'http://localhost:8080/api/categories', {
          statusCode: 200,
          body: mockCategories,
        }).as('getCategories');
      });
  
      cy.visit('/dashboard');
      cy.wait('@getCategories');
    });
  
   
  
    it('Should display the correct number of category cards', () => {
      cy.get('.card').should('have.length', 10);
    });
  
    it('Should display correct category titles on each card', () => {
      const titles = [
        'Computer Science',
        'Web Development',
        'Cybersecurity',
        'Artificial Intelligence & Machine Learning',
        'Cloud Computing',
        'Databases',
        'Operating Systems',
        'Software Engineering',
        'Networking',
        'DevOps & Automation'
      ];
  
      cy.get('.card-title').each(($el, index) => {
        cy.wrap($el).should('contain.text', titles[index]);
      });
    });
  
    it('Should display an image for each category card', () => {
      cy.get('.card img').each(($img) => {
        cy.wrap($img)
          .should('have.attr', 'src')
          .and('match', /assets\/.*\.(png|jpg)$/);
      });
    });
  
    it('Should have a Sub Topics button on every card', () => {
      cy.get('.card').each(($card) => {
        cy.wrap($card).find('button').contains('Sub Topics');
      });
    });
  
    it('Should navigate to correct category page when Sub Topics is clicked', () => {
      cy.get('.card').first().within(() => {
        cy.contains('Sub Topics').click();
      });
      cy.url().should('include', '/Computer%20Science');
    });
  
    it('Should not show About modal by default', () => {
      cy.get('#aboutModal').should('not.be.visible');
    });
  
    it('Should open About modal when openAbout is triggered manually', () => {
      cy.window().then((win) => {
        win.document.getElementById('aboutModal')?.classList.remove('fade');
        (win as any).bootstrap.Modal.getOrCreateInstance('#aboutModal').show();
      });
  
      cy.get('#aboutModal').should('be.visible');
      cy.contains('About SkillArcade').should('exist');
    });
  
    it('Should show correct image paths in img src attributes', () => {
      const expectedImages = [
        'imgCS1.png',
        'img1.png',
        'img1.png',
        'img1.png',
        'img1.png',
        'img1.png',
        'img1.png',
        'img1.png',
        'img1.png',
        'img1.png'
      ];
  
      cy.get('.card img').each(($img, index) => {
        cy.wrap($img).should('have.attr', 'src').and('include', expectedImages[index]);
      });
    });
  
    it('Should show cards in a responsive layout (basic check)', () => {
      cy.viewport(375, 667); // iPhone 6/7/8 size
      cy.get('.card').should('be.visible');
    });
  });
  