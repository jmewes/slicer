import { expect, test } from 'vitest'

async function delayedMultiply(a: number, b: number): Promise<number> {
  return new Promise((resolve) => setTimeout(() => resolve(a * b), 50))
}

test('multiplies numbers within time limit', async () => {
  // Given two numbers 6 and 7
  const a = 6
  const b = 7

  // When multiplying them with an asynchronous delay
  const product = await delayedMultiply(a, b)

  // Then the product should be 42
  expect(product).toBe(42)
}, 1_000)
