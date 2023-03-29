describe('Sprint 3 Test 1', ()=> {
    it('sign up ', ()=> {
        cy.visit('http://localhost:4200')

        cy.contains('New user? Click to signup!!').click()

        cy.contains('Email address').type('YonasWorku')

        cy.contains('Password').type('Password')

        cy.contains('Sign up').click()



    })
})


describe('Sprint 3 Test 2', ()=> {
    it('login', ()=> {
        cy.visit('http://localhost:4200')

        cy.contains('Email address').type('YonasWorku')

        cy.contains('Password').type('Password')

        cy.contains('Login!').click()



    })
})


describe('Sprint 3 Test 3', ()=> {
    it('Clicking on the tabs', ()=> {
        cy.visit('http://localhost:4200')

        cy.contains('About Us').click()

        cy.contains('Login').click()



    })
})

describe('Sprint 3 Test 4', ()=> {
    it('clicking on anything other than the three tabs', ()=>{
        cy.visit('http://localhost:4200')

        cy.contains('Find a Course to Rate').click()

        cy.contains('Manage').click()


    
    })
})

describe('Sprint 3 Test 5', ()=> {
    it('Discussion Board', ()=>{
        cy.visit('http://localhost:4200')

        cy.contains('Find a Course to Rate').click()

        cy.contains('Introduction to Software Engineering').trigger('mouseover')

        cy.contains('Introduction to Software Engineering Link').click()





    
    })
})

describe('Sprint 3 Test 6', ()=> {
    it('Search Bar', ()=>{
        cy.visit('http://localhost:4200')

        cy.contains('Find a Course to Rate').click()

        cy.get('input[placeholder=\"Course Name\"').click().type('Introduction to Computer Organization')

        cy.contains('Introduction to Computer Organization').click()

    
    })
})





