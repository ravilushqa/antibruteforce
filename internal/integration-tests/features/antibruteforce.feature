Feature: antibruteforce
  In order to use antibruteforce
  As an GRPC client
  I need to be able to send grpc requests

  Scenario Outline: should check bruteforce
    Given login is "<login>"
    And password is "<password>"
    And ip is "<ip>"
    When I call grpc method "Check"
    Then response error should be "<error>"

    Examples:
      | login  | password | ip      | error              |
      | login1 | pass1    | 1.2.3.4 | nil                |
      | login1 | pass1    | 1.2.3.4 | nil                |
      | login1 | pass1    | 1.2.3.4 | bucket is overflow |
      | login1 | pass2    | 1.2.3.4 | bucket is overflow |