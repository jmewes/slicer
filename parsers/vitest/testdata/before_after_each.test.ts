import { afterEach, beforeEach, expect, test } from 'vitest'

let accumulator: number

beforeEach(() => {
  // Given a freshly reset accumulator set to 0
  accumulator = 0
})

afterEach(() => {
  // Then ensure accumulator is reset to 0 after test run
  accumulator = 0
})

test('accumulates addition operations', () => {
  // Given accumulator initialized at 0
  // When adding 15 to the accumulator
  accumulator += 15

  // Then accumulator should equal 15
  expect(accumulator).toBe(15)
})
