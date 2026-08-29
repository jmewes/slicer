function setup(initialValue: number = 0): Calculator {
  // Given a calculator initialized with initial value
  const calc = new Calculator(initialValue);
  return calc;
}

describe('Calculator', () => {
  it('should add numbers', () => {
    // Given a calculator with value 10
    const calc = setup(10);
    // When adding 5
    calc.add(5);
    // Then the result should be 15
    expect(calc.getValue()).toBe(15);
  });
});
