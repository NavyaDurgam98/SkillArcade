describe('Category Page Tests', () => {
    const mockResponse = [
      { category: "Computer Science", imgPath: "computerscience.jpg" },
      { category: "Cloud Computing", imgPath: "cloudcomputing.jpg" }
    ];
  
    beforeEach(() => {
      cy.intercept('GET', 'http://localhost:8080/api/categories?searchText=comp', (req) => {
        const search = req.query["searchText"];
        if (typeof search === 'string' && search.toLowerCase().includes('comp')) {
          req.reply(mockResponse);
        } else {
          req.reply([]);
        }
      }).as('getCategories');
  
      const categoryName = 'Computer Science';
      cy.visit(`/${encodeURIComponent(categoryName)}`);
    });
  
    it('Should call API and display filtered subcategories when searchText = "comp"', () => {
      cy.wait('@getCategories');
  
      cy.get('.subcategory-card')  // Use the appropriate class or selector from your HTML
        .should('have.length', mockResponse.length);
    });
  
    it('Should display correct category titles from mock data', () => {
      cy.wait('@getCategories');
  
      cy.get('.subcategory-card .card-title')  // Adjust selector as needed
        .then(($titles) => {
          expect($titles).to.have.length(mockResponse.length);
          $titles.each((index, el) => {
            expect(el.textContent?.trim()).to.eq(mockResponse[index].category);
          });
        });
    });
  
    it('Should display correct images for each category', () => {
      cy.wait('@getCategories');
  
      cy.get('.subcategory-card img').each(($img, index) => {
        cy.wrap($img)
          .should('have.attr', 'src')
          .and('include', mockResponse[index].imgPath);
      });
    });
  
    it('Should navigate to quiz topic page on clicking a subcategory card', () => {
      cy.wait('@getCategories');
  
      cy.get('.subcategory-card').first().click();
  
      const expectedUrl = `/Computer%20Science/${encodeURIComponent(mockResponse[0].category)}`;
      cy.url().should('include', expectedUrl);
    });
  
    it('Should display "No results found" when no match is returned', () => {
      // Intercept with an empty result
      cy.intercept('GET', 'http://localhost:8080/api/categories?searchText=xyz', {
        statusCode: 200,
        body: []
      }).as('noResults');
  
      // Simulate changing the search text programmatically
      cy.window().then((win) => {
        (win as any).ng.getComponent(document.querySelector('app-category')).activeComponentService.searchText$.next('xyz');
      });
  
      cy.wait('@noResults');
      cy.get('.subcategory-card').should('have.length', 0);
      cy.contains('No results found').should('be.visible'); // Update message based on your UI
    });
  });
  