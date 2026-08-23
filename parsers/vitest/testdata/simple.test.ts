import { expect, test } from 'vitest'

test('adding two numbers', () => {
  // Given two numbers 5 and 3
  const a = 5
  const b = 3

  // When they are added together
  const sum = a + b

  // Then the sum should equal 8
  expect(sum).toBe(8)
})
