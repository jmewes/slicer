import { describe, expect, test } from 'vitest'

describe('Arithmetic module', () => {
  describe('Addition sub-module', () => {
    test('adds two positive integers correctly', () => {
      // Given two addends 12 and 8
      const a = 12
      const b = 8

      // When adding them
      const sum = a + b

      // Then sum should equal 20
      expect(sum).toBe(20)
    })
  })
})
