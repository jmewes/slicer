import { afterAll, beforeAll, expect, test } from 'vitest'

let dataset: number[]

beforeAll(async () => {
  // Given a shared dataset prepared before all tests run
  dataset = [10, 20, 30]
})

afterAll(async () => {
  // Then clean up shared dataset after all tests finish
  dataset = []
})

test('calculates sum of dataset elements', async () => {
  // Given the initialized dataset
  // When summing up all array elements
  const total = dataset.reduce((acc, val) => acc + val, 0)

  // Then total sum should be 60
  expect(total).toBe(60)
})
