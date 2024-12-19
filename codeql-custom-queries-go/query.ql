import go

/**
 * @name Insecure Random Number Generator
 * @description Detects usage of insecure random number generators like math/rand.
 * @kind problem
 * @problem.severity warning
 */
from CallExpr call
where call.getTarget().hasQualifiedName("math/rand", "Intn")
select call, "Usage of math/rand.Intn detected. Consider using crypto/rand for secure randomness."
