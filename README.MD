# factoryExample
Demonstration of interface based coding style, in which the domain logic need not be aware of runtime context.

Domain Logic is specified as interactions between interfaces, which are assembled at the start of the application
using dependency injection and using factories for construction of concrete classes.

A key aspect of the construction phase at application start is to load the security policy, which determines the
appropriate domain logic that can be accessed by any specific request.  Likely the security policy is specified
via role based access controls but this is not a requirement.  

Each factory is aware of the policy, and can select the appropriate domain logic instance to emit based on the
current runtime context and the policy.

The initial construction phase emits instances of a "shim" for each interface type, which are used to wire the
application together.  Shims can test for runtime context changes and ask their factory to return new instances
of the domain logic that should be used in this context.

As a result, the domain logic need not be aware of its context, making its source code clearer and simpler to test.



