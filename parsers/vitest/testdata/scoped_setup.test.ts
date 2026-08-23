import { beforeEach, describe, expect, test } from 'vitest'

describe('Scoped calculation context', () => {
  let baseValue: number

  beforeEach(() => {
    // Given a scoped base value set to 10
    baseValue = 10
  })

  test('adds value within scoped setup', () => {
    // Given base value 10
    // When adding 5 to base value
    const result = baseValue + 5

    // Then result should be 15
    expect(result).toBe(15)
  })
})
