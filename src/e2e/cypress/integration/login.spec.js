describe('Login correctly', function () {
    it('should redirect to login when not authenticated', function () {
        cy.clearCookies()
        cy.visit('http://localhost:3000/admin')
        cy.url().should('eq', 'http://localhost:3000/login')
    })
    it('should login with correct credentials and redirect to backoffice', function() {
        cy.clearCookies()
        cy.visit('http://localhost:3000/login')
        cy.login('admin', 'admin')
        cy.url().should('eq', 'http://localhost:3000/admin/#!/talks')
    })
})
