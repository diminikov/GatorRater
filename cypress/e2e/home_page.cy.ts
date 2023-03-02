describe('Sprint 2 Test 1', ()=> {
    it('Clicking on the tabs', ()=> {
        cy.visit('http://localhost:4200')

        cy.contains('About Us').click()

        cy.contains('Login').click()

        cy.contains('Find a Course to Rate').click()


    })
})

describe('Sprint 2 Test 2', ()=> {
    it('clicking on the course', ()=>{
        cy.visit('http://localhost:4200')

        cy.contains('Introduction to Software Engineering').click()

    
    })
})

describe('Sprint 2 Test 3', ()=> {
    it('Search Bar', ()=>{
        cy.visit('http://localhost:4200')

        cy.get('input[placeholder=\"Course Name\"').click().type('Introduction to Computer Organization')

        cy.contains('Introduction to Computer Organization').click()

    
    })
})


describe('Sprint 2 Test 4', ()=> {
    it('Username and Password', ()=> {
        cy.visit('http://localhost:4200')

        cy.contains('Login').click()

        cy.contains('User Name').type('YonasWorku')

        cy.contains('Password').type('Password')

        cy.contains('Submit').click()
    })
})



