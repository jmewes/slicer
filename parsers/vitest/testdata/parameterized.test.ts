import { expect, test } from 'vitest'

test.for([
  [1, 2, 3],
  [5, 5, 10],
  [100, 200, 300],
])('adds %i and %i to get %i', ([a, b, expected]) => {
  // Given inputs a and b
  // When adding a and b together
  const sum = a + b

  // Then sum should equal expected
  expect(sum).toBe(expected)
})
