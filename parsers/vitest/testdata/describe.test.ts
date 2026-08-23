import { describe, expect, test } from 'vitest'

describe('Arithmetic operations', () => {
  test('multiplies two positive numbers', () => {
    // Given two factors 4 and 3
    const factorA = 4
    const factorB = 3

    // When multiplying the factors
    const product = factorA * factorB

    // Then the product should be 12
    expect(product).toBe(12)
  })
})
