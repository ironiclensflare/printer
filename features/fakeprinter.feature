Feature: Fake printer
    In order to test the printer locally
    As a developer
    I need to be able to simulate an actual print output

    Scenario: Printing some text
        Given I have an instance of the fake printer
        When I print "Hello World!"
        Then I should receive a simulated printout containing "Hello World!"

    Scenario: Printing a file
        Given I have an instance of the fake printer
        When I print the file "blah.png"
        Then I should receive a simulated printout containing "Just pretend I printed blah.png"
