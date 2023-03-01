describe('The Home page', ()=> {
    it('successfully loads', ()=> {
        cy.visit('http://localhost:4200')

        cy.contains('About Us').click()

        cy.contains('Login').click()

        cy.contains('Find a Course to Rate').click()


    })
})