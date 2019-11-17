Feature: antibruteforce
  In order to use antibruteforce
  As an GRPC client
  I need to be able to send grpc requests

  Scenario Outline: should check for bruteforce
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

  Scenario Outline: should reset bucket
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method | login  | password | ip      | error              |
      | Check  | login1 | pass2    | 1.2.3.4 | bucket is overflow |
      | Reset  | login1 |          | 1.2.3.4 | nil                |
      | Check  | login1 | pass2    | 1.2.3.4 | nil                |

  Scenario Outline: should add subnet to blacklist
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"
    And subnet is "192.168.0.0/25"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method       | login  | password | ip           | error            |
      | Check        | login3 | pass3    | 192.168.0.30 | nil              |
      | BlacklistAdd |        |          |              | nil              |
      | Check        | login3 | pass3    | 192.168.0.30 | ip in black list |

  Scenario Outline: should remove subnet from blacklist
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"
    And subnet is "192.168.0.0/25"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method          | login  | password | ip           | error            |
      | Check           | login4 | pass4    | 192.168.0.30 | ip in black list |
      | BlacklistRemove |        |          |              | nil              |
      | Check           | login4 | pass4    | 192.168.0.30 | nil              |

  Scenario Outline: should add subnet to whitelist
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"
    And subnet is "192.168.0.0/25"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method       | login  | password | ip           | error              |
      | Check        | login5 | pass5    | 192.168.0.30 | nil                |
      | Check        | login5 | pass5    | 192.168.0.30 | nil                |
      | Check        | login5 | pass5    | 192.168.0.30 | bucket is overflow |
      | WhitelistAdd |        |          |              | nil                |
      | Check        | login5 | pass5    | 192.168.0.30 | nil                |

  Scenario Outline: should remove subnet from whitelist
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"
    And subnet is "192.168.0.0/25"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method          | login  | password | ip           | error              |
      | Check           | login6 | pass6    | 192.168.0.30 | nil                |
      | Check           | login6 | pass6    | 192.168.0.30 | nil                |
      | Check           | login6 | pass6    | 192.168.0.30 | nil                |
      | WhitelistRemove |        |          |              | nil                |
      | Check           | login6 | pass6    | 192.168.0.30 | nil                |
      | Check           | login6 | pass6    | 192.168.0.30 | nil                |
      | Check           | login6 | pass6    | 192.168.0.30 | bucket is overflow |