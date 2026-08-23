import { expect, it } from 'vitest'

it('subtracts numbers correctly', () => {
  // Given minuend 10 and subtrahend 4
  const minuend = 10
  const subtrahend = 4

  // When subtracting subtrahend from minuend
  const difference = minuend - subtrahend

  // Then the difference should equal 6
  expect(difference).toBe(6)
})
