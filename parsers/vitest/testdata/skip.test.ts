import { expect, test } from 'vitest'

test.skip('division by zero produces infinity', () => {
  // Given dividend 10 and divisor 0
  const dividend = 10
  const divisor = 0

  // When dividing dividend by divisor
  const quotient = dividend / divisor

  // Then result should be Infinity
  expect(quotient).toBe(Infinity)
})
