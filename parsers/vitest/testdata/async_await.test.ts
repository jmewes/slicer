import { expect, test } from 'vitest'

async function asyncAdd(a: number, b: number): Promise<number> {
  return Promise.resolve(a + b)
}

test('adds numbers asynchronously', async () => {
  // Given two numbers 7 and 8
  const a = 7
  const b = 8

  // When adding them via an async function
  const sum = await asyncAdd(a, b)

  // Then the calculated sum should be 15
  expect(sum).toBe(15)
})
