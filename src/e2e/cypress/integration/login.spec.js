describe('Authentication', function () {

    it('should redirect to login when not authenticated', function () {
        cy.clearCookies()
        cy.visit('http://localhost:3000/admin')
        cy.url().should('eq', 'http://localhost:3000/login')
    })

    it('should display an error message when credentials are wrong', function () {
        cy.clearCookies()
        cy.visit('http://localhost:3000/login')
        cy.loginAs('adminWrong')
        cy.url().should('eq', 'http://localhost:3000/login')
        cy.contains('.x-login-error', 'Wrong credentials')
    })

    it('should re-enter credentials after a failure and login successfully', function () {
        cy.clearCookies()
        cy.visit('http://localhost:3000/login')
        cy.loginAs('adminWrong')
        cy.loginAs('admin')
        cy.url().should('eq', 'http://localhost:3000/admin/#!/talks')
    })

    it('should login with correct credentials and redirect to backoffice', function() {
        cy.clearCookies()
        cy.visit('http://localhost:3000/login')
        cy.loginAs('admin')
        cy.url().should('eq', 'http://localhost:3000/admin/#!/talks')
    })

})
