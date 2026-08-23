import { describe, expect, test } from 'vitest'

describe('Well formed part', () => {
  test('works fine', () => {
    // Given a valid value
    const value = 1

    // When nothing else happens
    // Then it should still equal 1
    expect(value).toBe(1)
  })
})

test('broken part', () => {
  // Given an unterminated call below
  expect(1).toBe(1)
