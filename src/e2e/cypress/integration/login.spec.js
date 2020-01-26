describe('Login correctly', function () {
    it('should redirect to login when not authenticated', function () {
        cy.clearCookies()
        cy.visit('http://localhost:3000/admin')
        cy.url().should('eq', 'http://localhost:3000/login')
    })
    it('should login with correct credentials and redirect to backoffice', function() {
        cy.clearCookies()
        cy.visit('http://localhost:3000/login')
        cy.get('input#username').type('admin')
        cy.get('input#password').type('admin')
        cy.get('input[type=submit]').click()
        cy.url().should('eq', 'http://localhost:3000/admin/#!/talks')
    })
})
