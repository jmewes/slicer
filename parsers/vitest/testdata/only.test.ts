import { expect, test } from 'vitest'

test.only('divides two numbers', () => {
  // Given dividend 20 and divisor 5
  const dividend = 20
  const divisor = 5

  // When dividing dividend by divisor
  const quotient = dividend / divisor

  // Then quotient should equal 4
  expect(quotient).toBe(4)
})
