describe('My First Test', () => {
  it('Visits the local browser', () => {
    cy.visit('http://localhost:4200')

    //cy.contains('type').click()

    //cy.url().should('include', '/commands/actions')

    //cy.get('.action-email').type('fake@email.com')

    //cy.get('.action-email').should('have.value', 'fake@email.com')
  })
})